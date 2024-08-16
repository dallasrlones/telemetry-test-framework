const express = require('express');
const server = express();

// Middleware to parse JSON bodies
server.use(express.json());

// GET endpoint
server.get('/', (_req, res) => {
    console.log('Received GET request')
    res.send('GET request received');
});

server.get('/quit', (_req, _res) => {
    console.log('Testing Server Shutting Down')
    process.exit(1);
});

server.get('/success', (_req, _res) => {
    console.log('Testing Server Shutting Down - SUCCESS')
    process.exit(0);
});

// POST endpoint
server.post('/', (req, res) => {
    const receivedData = req.body;
    console.log('Received POST data:', receivedData);
    res.send('POST request received');
});

server.listen(process.env.PORT, err => {
    console.log(err || 'Testing HTTP server online');
});
