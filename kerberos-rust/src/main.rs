extern crate hyper;
extern crate pretty_env_logger;
extern crate kkdcp;
extern crate hexdump;

use std::io::{Read, Write};
use std::str::FromStr;
use hyper::Post;
use hyper::server::{Server, Request, Response};
use hyper::uri::RequestUri::AbsolutePath;
use hyper::header::ContentType;
use hyper::mime::Mime;
use std::net::TcpStream;

fn handle_kkdcp(mut req: Request, mut res: Response) {
    match req.uri {
        AbsolutePath(ref path) => {
            match (&req.method, &path[..]) {
                (&Post, "/") => (),
                _ => {
                    *res.status_mut() = hyper::NotFound;
                    return;
                }
            }
        }
        _ => {
            return;
        }
    };

    let mut body = Vec::new();
    if let Ok(len) = req.read_to_end(&mut body) {
        println!("read {} bytes", len);
        if len == 0 {
            *res.status_mut() = hyper::BadRequest;
            return;
        }
        hexdump::hexdump(&body);
    }

    let msg = kkdcp::decode(body);
    println!("got {} bytes", msg.kerb_message.len());
    let resp = forward_kerberos(msg.kerb_message);
    let msg = kkdcp::encode(resp);
    println!("Sending {} bytes to client", msg.len());

    {
        let mut headers = res.headers_mut();
        headers.set(ContentType(Mime::from_str("application/kerberos").unwrap()));
    }
    hexdump::hexdump(&msg);
    res.send(msg.as_slice()).expect("Composing the result failed.");
}

fn forward_kerberos(data: Vec<u8>) -> Vec<u8> {
    let mut stream = TcpStream::connect("visheshchoudhary.me:88").expect("Connecting to Kerberos server failed");
    let mut input = stream.try_clone().expect("Failed to get input stream");
    let output = &mut stream;

    output.write(data.as_slice()).expect("Failed to write to Kerberos server");
    output.flush().expect("Failed to flush data out");

    let mut resp = Vec::new();
    if let Ok(len) = input.read_to_end(&mut resp) {
        println!("Read {} bytes from kerberos", len);
    }
    resp
}

fn main() {
    pretty_env_logger::init().expect("Failed to init logger");
    let server = Server::http("127.0.0.1:8125").expect("Failed to start server");
    let _guard = server.handle(handle_kkdcp);
}
