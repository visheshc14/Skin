#[macro_use]
extern crate eagre_asn1;

use eagre_asn1::der::DER;

pub struct KdcProxyMessage {
    pub kerb_message: Vec<u8>,
    pub target_domain: String,
}

der_sequence!{
    KdcProxyMessage:
        kerb_message:    EXPLICIT TAG CONTEXT 1; TYPE Vec<u8>,
        target_domain:   EXPLICIT TAG CONTEXT 2; TYPE String,
}

pub struct KdcProxyResponse {
    pub kerb_message: Vec<u8>,
}

der_sequence! {
    KdcProxyResponse:
        kerb_message:   EXPLICIT TAG CONTEXT 1; TYPE Vec<u8>,
}

pub fn decode(data: Vec<u8>) -> KdcProxyMessage {
    println!("got {} bytes", data.len());
    let msg = KdcProxyMessage::der_from_bytes(data).expect("Decoding KdcProxyMessage data failed");
    msg
}

pub fn encode(data: Vec<u8>) -> Vec<u8> {
    println!("got {} bytes", data.len());
    let krb_msg = KdcProxyResponse{kerb_message: data};
    krb_msg.der_bytes().unwrap()
}
