const jwt = require('jsonwebtoken');

function decryptMessages(encryptedMessage, secretKey){
    var message;
    jwt.verify(encryptedMessage, secretKey, (err, decryptedMessages) => {
        if(err) {
            console.log("Error while decryption: ", err);
        }
        // console.log(decryptedMessages);
        message = decryptedMessages.message;
        // console.log(message);
    })
    return message;
}

export default decryptMessages;
