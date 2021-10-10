const express = require('express');
const cors = require('cors');
const dotenv = require('dotenv');
const connectDB = require('./config/db');

const serverRouter = require('./routes/signin'); 

dotenv.config();

const PORT = 5000;

const app = express();
app.use(express.urlencoded({
    extends: true
}));
app.use(cors());
app.use(express.json());

connectDB();

app.use("/server", serverRouter);

app.get("/", (req, res) => {
    res.send("Server is running!!!")
})

app.listen(PORT, () => { 
    console.log(`Server running at ${PORT}.`)
});
