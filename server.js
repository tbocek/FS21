const net = require('net');
const rl = require('readline');
const server = net.createServer(onClientConnected);

server.listen(8081, 'localhost', function() {
    console.log('Launching server...');
});

function onClientConnected(sock) {
    const i = rl.createInterface(sock, sock);

    i.on('line', function(line) {
        console.log('Message Received: %s', line);
        sock.write(line.toUpperCase() + "\n");
    });
};
