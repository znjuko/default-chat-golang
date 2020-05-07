import API from '../utils/API';
import { webSocket } from 'rxjs/webSocket';

export default async function connectSocket() {
	try {
		let res = await API.get('/ws', {
			params: {
				results: 1,
				inc: 'token',
			},
		});
		let token = res.data.token;

		console.log(res);
		window.socket = webSocket(`ws://localhost:3000/ws/${token}`);

		window.socket.subscribe(
			(msg) => {
				console.log('message received: ');
				console.log(msg);
			}, // Called whenever there is a message from the server.
			(err) => {
				console.log('[DEBUG]: error with ws: ');
				console.log(err);
			}, // Called if at any point WebSocket API signals some kind of error.
			() => console.log('complete'), // Called when connection is closed (for whatever reason).
		);
		// window.socket = new WebSocket(`ws://localhost:3000/ws/${token}`);

		/* window.socket.onopen = function (e) {
			console.log('[WS] Соединение установлено');
		};

		window.socket.onmessage = function (event) {
			console.log('[WS] Message', event);
			var data = JSON.parse(event.data);

			let newMsg = {
				author: data.message.author,
				txt: data.message.txt,
				chatName: data.message.chatName,
				chatId: data.message.chatId,
				emojies: data.message.emojies,
			};
			observer.next(newMsg);
		};

		window.socket.onclose = function (event) {
			if (event.wasClean) {
				console.log(
					`[WS] Соединение закрыто чисто, код=${event.code} причина=${event.reason}`,
				);
			} else {
				connectSocket();
				console.log('[WS] Соединение прервано');
			}
		};

		window.socket.onerror = function (error) {
			console.error(`[WS] ${error.message}`);
		}; */
	} catch (err) {
		console.error(`[WS] ${err}`);
	}
}
