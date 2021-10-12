# Skin
Web Security (BCI3001) Project - Prevention of Session Hijacking using Session ID Reset Approach with the Implementation of Kerberos Algorithm in Go &amp; Rust.

Session ID - Reset Approach with Implementation of Kerberos Algorithm.  

Express session stores sessionID in cookie and it will set that cookie in frontend (browser and you can see that cookie called connect.sid in browser) from backend (server). Whenever any request comes from browser first it will check for that cookie (in which sessionID is stored.) If it finds the cookie it doesn't create new session otherwise it will create again a new session. (you can check it by logging req.sessionID in requests).

To overcome this for every request we are making from frontend (browser) we have to send that cookie to backend (server). Server will automatically parse cookie and doesn't create any new session for every request.

We were using axios for request calls in which for every request we were adding {withCredentals:true} so that browser can send cookies to backend server (automatically).

Example - 
```javascript
var FileStore = require('session-file-store')(session);

app.use(cors({
 origin:[process.env.ORIGIN],//Frontend Server localhost:8080
 methods:['GET','POST','PUT','DELETE'],
 credentials: true // Enable Set Cookie
}));

app.use(cookieParser(process.env.SESSIONSECRET)); // Any String Ex: 'Keyboard Cat'
app.use(session({
  secret: process.env.SESSIONSECRET,
  store:new FileStore,
  cookie:{
    maxAge:36000,
    httpOnly:false,
    secure:false // For Normal HTTP Connection if HTTPS is There We Have to Set it to True
    },
  resave: false,
  saveUninitialized: true
})) 

app.use(function(req, res, next) {

res.header('Access-Control-Allow-Credentials', true);
res.header('Access-Control-Allow-Methods', 'GET, PUT, POST, DELETE');
res.header("Access-Control-Allow-Origin", process.env.ORIGIN);
res.header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-   Type, Accept, Authorization");
next();
});

app.use(bodyParser.json());
app.use(bodyParser.urlencoded({extended: true}));

// Rest of Code is The Mongo Connection
```
Axios Rest Calls::
```javascript
 axios.defaults.withCredentials = true;
 axios.get('http://localhost:8080/getDetails',{
           headers:{
                    withCredentials:true,

                   }
  });
```
<img width="1440" alt="2" src="https://user-images.githubusercontent.com/36515357/136685074-01f423ef-1b2a-42d3-85b7-bbe8844e4139.png">

<img width="1440" alt="3" src="https://user-images.githubusercontent.com/36515357/136685076-64a139fd-7aaf-443f-a5d9-b70a2293a47f.png">

<img width="1440" alt="4" src="https://user-images.githubusercontent.com/36515357/136685080-70c863b8-4f8f-45c9-b53b-c48e445fba6d.png">

![architecture](https://user-images.githubusercontent.com/36515357/136011313-ca3a6bec-b710-468f-ba76-b704944a4693.png)

Kerberos in Rust.
Reverse Proxy Server to Front This and Handle SSL.

```
cargo build
./target/debug/kkdcp
```
```
server {
	listen 443;
	listen [::]:443;
	server_name visheshchoudhary.me;

	ssl on;
	ssl_certificate /etc/ssl/certs/kdcproxy.pem;
	ssl_certificate_key /etc/ssl/private/kdcproxy.key;

	root /var/www/kdxproxy;
	index index.html;

	location /KdcProxy {
		proxy_pass http://127.0.0.1:8125/;
		include proxy_params;
		add_header Cache-Control "no-cache, no-store, must-revalidate";
		add_header Pragma no-cache;
		add_header Expires 0;
	}
}
```
![Screenshot (13)](https://user-images.githubusercontent.com/36515357/136888548-0ac43cb1-f1fa-4c62-99aa-5f4d5c494dd0.png)

![Screenshot (14)](https://user-images.githubusercontent.com/36515357/136888557-e9da53a8-7956-4a5d-b86e-67911390f1b5.png)

Kerberos in Go Example With Two Different API's To Grasp out The Difference Between Two Approaches. 

In Kerberos.go GSSAPI Has Been Used, The Generic Security Service Application Program Interface is an application programming interface for programs to access security services. The GSSAPI is an IETF standard that addresses the problem of many similar but incompatible security services in use today.

```GO
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
}

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
```  
HTTP Kerberos Authentication - Implemented in GO, Example Authentication Using Chiltak API - The Documentation Referred Explained Here.

Authentication can be added to any method that sends an HTTP request to the server, such as SynchronousRequest, QuickGetStr, PostXml, etc. To add authentication, simply set the Login and Password properties.

By default, Chilkat will use basic HTTP authentication, which sends the login/password clear-text over the connection. This is bad if SSL/TLS (i.e. HTTPS) is not used. However, if the connection is secure, there should be nothing wrong with using basic authentication.

Chilkat supports more secure authentication types as well, including Digest, NTLM, and Negotiate (which dynamically chooses between NTLM and Kerberos). To use Digest authentication, simply set the DigestAuth property = true. To use NTLM authentication, set the NtlmAuth property = true. Likewise, to use Negotiate authentication, set the NegotiateAuth property = true.

```GO
    // This Example Assumes Chilkat API To Be The Best Alternative For Explanation.
    http := chilkat.NewHttp() 

    // Set the Login and Password properties for authentication.
    http.SetLogin("chilkat")
    http.SetPassword("myPassword")

    // To use HTTP Basic authentication..
    http.SetBasicAuth(true)

    html := http.QuickGetStr("http://localhost/xyz.html")
    if http.LastMethodSuccess() != true {
        fmt.Println(http.LastErrorText())
        http.DisposeHttp()
        return
    }

    // Examine the HTTP status code returned.  
    // A status code of 401 is typically returned for "access denied"
    // if no login/password is provided, or if the credentials (login/password)
    // are incorrect.
    fmt.Println("HTTP status code for Basic authentication: ", http.LastStatus())

    // Examine the HTML returned for the URL:
    fmt.Println(*html)

    http2 := chilkat.NewHttp()

    // To use NTLM authentication, set the 
    // NtlmAuth property = true
    http2.SetNtlmAuth(true)

    // The session log can be captured to a file by
    // setting the SessionLogFilename property:
    http2.SetSessionLogFilename("ntlmAuthLog.txt")

    // Examination of the HTTP session log will show the NTLM
    // back-and-forth exchange between the client and server.

    // This call will now use NTLM authentication (assuming it
    // is supported by the web server).
    html = http2.QuickGetStr("http://localhost/xyz.html")
    // Note: 
    if http2.LastMethodSuccess() != true {
        fmt.Println(http2.LastErrorText())
        http.DisposeHttp()
        http2.DisposeHttp()
        return
    }

    fmt.Println("HTTP status code for NTLM authentication: ", http2.LastStatus())

    http3 := chilkat.NewHttp()

    // To use Digest Authentication, set the DigestAuth property = true
    // Also, no more than one of the authentication type properties 
    // (NtlmAuth, DigestAuth, and NegotiateAuth)  should be set
    // to true.  
    http3.SetDigestAuth(true)

    http3.SetSessionLogFilename("digestAuthLog.txt")

    // This call will now use Digest authentication (assuming it
    // is supported by the web server).
    html = http3.QuickGetStr("http://localhost/xyz.html")
    if http3.LastMethodSuccess() != true {
        fmt.Println(http3.LastErrorText())
        http.DisposeHttp()
        http2.DisposeHttp()
        http3.DisposeHttp()
        return
    }

    fmt.Println("HTTP status code for Digest authentication: ", http3.LastStatus())

    http.DisposeHttp()
    http2.DisposeHttp()
    http3.DisposeHttp()
```
# Made by Vishesh Choudhary, Aditi Jain & Pranav Singh :heart:
