const jwt = require('jsonwebtoken');
const TGS_SecretKeys = require('../models/TGS_SecretKey');

const decryptedMessages = require('./decryptMessages');


function authenticateUserCredentials(req, res, next){
    console.log("inside TGS middleware");
    const encryptedTGT = req.body.encryptedTGT;
    const encryptedUserCredentials = req.body.encryptedUserCredentials;

    // console.log(encryptedTGT, encryptedUserCredentials)

    const stored_TGS_SecretKey = TGS_SecretKeys[0];
    // console.log(stored_TGS_SecretKey);
    TGS_SecretKeys.pop();

    const {username: TGT_username, TGS_SessionKey: TGS_SessionKey} = decryptedMessages(encryptedTGT, stored_TGS_SecretKey);
    // console.log(TGS_SessionKey);
    const {username: userCredentials_username} = decryptedMessages(encryptedUserCredentials, TGS_SessionKey);

    if(TGT_username != userCredentials_username){
        res.send({
            success: false,
            message: "Error: Access Denied!!"
        })
    }

    req.username = userCredentials_username;
    req.TGS_SessionKey = TGS_SessionKey;
    next();

}

module.exports = { authenticateUserCredentials };
