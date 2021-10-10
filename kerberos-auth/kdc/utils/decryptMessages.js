const jwt = require('jsonwebtoken');

function decryptMessages(encryptedMessage, secretKey){
    let message;
    jwt.verify(encryptedMessage, secretKey, (err, decryptedMessages) => {
        if(err) {
            console.log("Error while decryption: ", err);
        }

        message = decryptedMessages.message;
    })

    return message;
}

module.exports = decryptMessages;
