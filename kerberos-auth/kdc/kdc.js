const express = require('express');
const cors = require('cors');
const dotenv = require('dotenv');
const connectDB = require('./config/db');
const authenticatorServer = require("./routes/authenticatorServer");
const ticketGrantingServer = require('./routes/ticketGrantingServer');

dotenv.config();

const PORT = 4000;

const app = express();
app.use(express.urlencoded({
    extends: true
}));
app.use(cors());
app.use(express.json());

connectDB();

app.use("/authenticator_server", authenticatorServer);
app.use("/ticket_granting_server", ticketGrantingServer);

app.get("/", (req, res) => {
    res.send("KDC server is running!!!")
})

app.listen(PORT, () => { 
    console.log(`Server running at ${PORT}.`)
});
