const express = require('express');
const crypto = require('crypto');
const bcrypt = require('bcryptjs');
const decryptedMessages = require('../utils/decryptMessages');
const {generateAccessToken} = require('../utils/authenticateToken');

const AccessLogs = require('../models/accessLogs');

const router = express.Router();

router.post('/signin', async (req, res) => {
    console.log("inside server route");

    const encryptedUserCredentials = req.body.newEncryptedUserCredentials;
    const encryptedServiceTicket = req.body.encryptedServiceTicket;
    const serviceSecretKey = req.body.serviceSecretKey;

    const {
        username: ServiceTicket_username,
        Service_ID: serviceid,
        timeStamp: ServiceTicket_timeStamp,
        Service_SessionKey: Service_SessionKey
    } = decryptedMessages(encryptedServiceTicket, serviceSecretKey);

    // console.log(ServiceTicket_username, serviceid, ServiceTicket_timeStamp, Service_SessionKey);

    const {username: userCredentials_username, timeStamp: userCredentials_timestamp} = decryptedMessages(encryptedUserCredentials, Service_SessionKey);

    // console.log(userCredentials_username, ServiceTicket_username);

    if(ServiceTicket_username != userCredentials_username){
        res.send({
            success: false,
            message: "Error: Access Denied!!"
        })
    };

    // console.log(encryptedUserCredentials, encryptedServiceTicket, serviceSecretKey);

    const accessToken = generateAccessToken(userCredentials_username);

    res.send({
        success: true,
        accessToken: accessToken
    })

});

module.exports = router;
