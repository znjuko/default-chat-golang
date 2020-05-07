import API from '../utils/API';

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
		window.socket = new WebSocket(`ws://localhost:3000/ws/${token}`);

		window.socket.onopen = function (e) {
			console.log('[WS] Соединение установлено');
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
		};
	} catch (err) {
		console.error(`[WS] ${err}`);
	}
}
