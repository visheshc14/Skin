//Minimal Working of Kerberos

package kerberos

import (
	"context"
	"fmt"
	"net/http"

	"github.com/apcera/gssapi"
	"github.com/apcera/gssapi/spnego"
)

type contextKey string

func (key contextKey) String() string {
	return fmt.Sprintf("kerberos/%s", string(key))
}

const (
	serverKey     = contextKey("server")
	credentialKey = contextKey("credential")
	userKey       = contextKey("user")
)

func Server(ctx context.Context) spnego.KerberizedServer {
	return ctx.Value(serverKey).(spnego.KerberizedServer)
}

func Credential(ctx context.Context) *gssapi.CredId {
	return ctx.Value(credentialKey).(*gssapi.CredId)
}

func User(ctx context.Context) string {
	return ctx.Value(userKey).(string)
}

func UserOk(ctx context.Context) (string, bool) {
	user, ok := ctx.Value(userKey).(string)
	return user, ok
}Min

func WithContext(ctx context.Context, keytab, spn string) (context.Context, error) {
	gss, err := gssapi.Load(&gssapi.Options{Krb5Ktname: keytab})
	if err != nil {
		return ctx, err
	}

	server := spnego.KerberizedServer{Lib: gss}
	ctx = context.WithValue(ctx, serverKey, server)

	cred, err := server.AcquireCred(spn)
	if err != nil {
		return ctx, err
	}

	return context.WithValue(ctx, credentialKey, cred), nil
}

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		server := Server(ctx)
		cred := Credential(ctx)
		user, status, err := server.Negotiate(cred, r.Header, w.Header())

		if status != http.StatusOK {
			http.Error(w, err.Error(), status)
			return
		}

		ctx = context.WithValue(ctx, userKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
