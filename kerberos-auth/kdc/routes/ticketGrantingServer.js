const express = require('express');
const crypto = require('crypto');
const bcrypt = require('bcryptjs');

const encryptMessages = require('../utils/encryptMessages');
const { authenticateUserCredentials } = require('../utils/checkUserCredentials');

const User = require('../models/user');

const router = express.Router();

router.post("/signin", authenticateUserCredentials, async (req, res) => {
    const username = req.username;
    const TGS_SessionKey = req.TGS_SessionKey;

    const Service_ID = crypto.randomBytes(5).toString('hex');
    const Service_SessionKey = crypto.randomBytes(10).toString('hex'); 
    const serviceSecretKey = await bcrypt.hash(Service_ID, 8);
    
    const timeElapsed = Date.now();
    const todaysDate = new Date(timeElapsed);
    const timeStamp = todaysDate.toUTCString();

    const Service_Ticket = {
        username,
        Service_ID,
        timeStamp,
        Service_SessionKey
    };

    const encryptedServiceTicket = encryptMessages(Service_Ticket, serviceSecretKey)

    const userCredentials = {
        username,
        timeStamp,
        Service_SessionKey
    };

    const encryptedUserCredentials = encryptMessages(userCredentials, TGS_SessionKey);

    res.send({
        success: true,
        encryptedServiceTicket,
        encryptedUserCredentials,
        serviceSecretKey
    })

})

module.exports = router;
