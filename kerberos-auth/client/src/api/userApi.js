import axios from 'axios';

const AS_URL = 'http://localhost:4000/authenticator_server';
const TGS_URL = 'http://localhost:4000/ticket_granting_server';
const Server_URL = 'http://localhost:5000/server';

export const signupNewUser = (user) => axios.post(`${AS_URL}/signup`, user);
export  const requestForAS = (user) => axios.post(`${AS_URL}/signin`, user);

export const requestForTGS = (user) => axios.post(`${TGS_URL}/signin`, user);

export const requestForServer = (user) => axios.post(`${Server_URL}/signin`, user);
