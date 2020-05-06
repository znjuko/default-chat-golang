import io from 'socket.io-client';

export default function connectSocket() {
	var socket = new io.Socket();
	socket.connect('http://localhost:3000');
	return socket;
}
