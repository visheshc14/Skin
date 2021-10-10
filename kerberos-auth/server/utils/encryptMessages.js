const jwt = require('jsonwebtoken');

function encryptMessages(message, secretKey){
    return jwt.sign({message}, secretKey);
}

module.exports = encryptMessages;
