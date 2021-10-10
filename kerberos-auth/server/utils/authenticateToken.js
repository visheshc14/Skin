const jwt = require('jsonwebtoken');
const dotenv = require('dotenv');
dotenv.config();

function generateAccessToken(message){
    console.log(message);
    return jwt.sign({message}, process.env.ACCESS_TOKEN_SECRET); //, {expiresIn: '15s'})
}

module.exports = {generateAccessToken};