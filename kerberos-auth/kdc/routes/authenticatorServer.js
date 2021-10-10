const express = require('express');
const crypto = require('crypto');
const bcrypt = require('bcryptjs');
const encryptMessages = require('../utils/encryptMessages');

const User = require('../models/user');
const TGS_SecretKeys = require('../models/TGS_SecretKey');

const router = express.Router();

router.post("/signup", async (req, res)=> {
    const {
        username,
        password
    } = req.body;

    const email = req.body.email.toLowerCase();

    if(!username){
        return res.send({
            success: false,
            message: "Error: Username field cannot be empty"
        });
    }

    if(!email){
        return res.send({
            success: false,
            message: "Error: Email field cannot be empty"
        });
    }

    if(!password){
        return res.send({
            success: false,
            message: "Error: Password field cannot be empty"
        });
    }

    User.find({
        username: username
    }, (err, existingUser) => {
        if(err){
            console.log("first if");
            return res.send({
                success: false,
                message: "Error: Server Error"
            });
        } else if (existingUser.length > 0) {
            return res.send({
                success: false,
                message: "Error: User already exists"
            });
        }
        
        const newUser = new User();
        newUser.email = email;
        newUser.username = username;
        newUser.password = password;

        newUser.save((err, response) => {
            if(err){
                console.log("second if");
                return res.send({
                    success: false,
                    message: "Error: Server Error"
                });
            };
            // console.log(response);
            return res.send({
                success: true,
                message: "User Sign up successfully"
            });
        });

    });
});

router.post("/signin", async (req, res) => {
    const username = req.body.username;
    const password = req.body.password;
    
    if(!username){
        return res.send({
            success: false,
            message: "Error: Username field cannot be empty"
        });
    }

    if(!password){
        return res.send({
            success: false,
            message: "Error: Password field cannot be empty"
        })
    }
    
    User.find({
        username: username,
    }, async (err, users) => {
        if(err){
            return res.send({
                success: false,
                message: "Error: Server Error"
            });
        } else if(users.length != 1){
            return res.send({
                success: false,
                message: "Error: DB Error"
            });
        }

        const user = users[0];
        const passwordPromise = await bcrypt.compare(password, user.password);
        
        console.log(passwordPromise);
        if(!passwordPromise){
            return res.send({
                success: false,
                message: "Error: Invalid Password"
            });
        }

        const timeElapsed = Date.now();
        const todaysDate = new Date(timeElapsed);
        const timeStamp = todaysDate.toUTCString();
        const TGS_ID = crypto.randomBytes(5).toString('hex');
        const TGS_SessionKey = crypto.randomBytes(10).toString('hex'); 
        const TGS_SecretKey = await bcrypt.hash(TGS_ID, 8);

        TGS_SecretKeys.push(TGS_SecretKey);
        
        messageForClient = {
            TGS_ID: TGS_ID,
            timestamp: timeStamp,
            TGS_SessionKey: TGS_SessionKey
        }

        messageForTGT = {
            username: user.username,
            TGS_ID: TGS_ID,
            timestamp: timeStamp,
            TGS_SessionKey: TGS_SessionKey
        }

        // console.log('messageForClient: ', messageForClient);
        // console.log('messageForTGT: ', messageForTGT);

        // console.log('client key: ', user.secretKey);

        const encryptedClientMessage = encryptMessages(messageForClient, user.secretKey);
        const encryptedTGT = encryptMessages(messageForTGT, TGS_SecretKey);

        // console.log("encrytedClientMessage", encryptedClientMessage);
        // console.log("encryptedTGT", encryptedTGT);
        res.json({
            success: true,
            encryptedClientMessage: encryptedClientMessage,
            encryptedTGT: encryptedTGT,
            clientSecretKey: user.secretKey 
        });
    });
    
});

module.exports = router;
