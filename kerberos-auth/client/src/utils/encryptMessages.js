const jwt = require('jsonwebtoken');

function encryptMessages(message, secretKey){
    const messageToReturn = jwt.sign({message}, secretKey);
    console.log(messageToReturn);
    return messageToReturn;
}
export default encryptMessages;
