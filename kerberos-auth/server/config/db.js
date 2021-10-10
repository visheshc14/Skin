const mongoose = require("mongoose");

const connectDB = async () => {
    try {
        const connectionResponse = await mongoose.connect(process.env.URL, {
			useNewUrlParser: true,
			useUnifiedTopology: true,
		});
        console.log(connectionResponse.connection.host);

		console.log(`MongoDB Connected: ${connectionResponse.connection.host}`);

    } catch(err){
        console.log('Mongo connection error: ', err);
        process.exit(1);
    }
} 

module.exports =  connectDB;