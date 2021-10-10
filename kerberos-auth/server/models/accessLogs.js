const mongoose = require("mongoose");
const { Schema } = mongoose;

const AccessLogsSchema = new Schema({
    username: {
        type: String,
        required: true,
        default: ''
    },
    serviceid: {
        type: String,
        required: true,
        default: '',
    },
    timestamp: {
        type: String,
        default: ''
    },
    accesstoken: {
        type: String,
        required: true,
        default: ''
    }
});

AccessLogsSchema.pre('save', async function (next){
    const log = this;
    const timeElapsed = Date.now();
    const todaysData = new Date(timeElapsed);
    log.timestamp = todaysData.toUTCString();
    next();
})

module.exports = mongoose.model('AccessLogs', AccessLogsSchema);
