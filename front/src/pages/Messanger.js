import React from 'react';
import API from '../utils/API';

import OnlineUsers from './OnlineUsers';
import Chats from './Chats';
import connectSocket from '../utils/Socket';

export default class Messanger extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			isLoading: true,
			chats: [],
			online: [],
		};
	}

	render() {
		return (
			<div
				className='container-fluid h-100 mt-2'
				style={{ minHeight: '90vh' }}>
				<div
					className='row h-100 d-flex justify-content-center'
					style={{ minHeight: '100vh' }}>
					<div className='col-md-6'>
						<h4 className='h4 mt-2 text-left'>Ваши чаты</h4>
						<hr />
						<Chats
							history={this.props.history}
							chats={this.state.chats}
						/>
					</div>
					<div className='col col-md-3 '>
						<h4 className='h4 mt-2 text-left'>
							Пользователи онлайн
						</h4>
						<hr />
						<OnlineUsers
							history={this.props.history}
							online={this.state.online}
						/>
					</div>
				</div>
			</div>
		);
	}

	async componentDidMount() {
		// Load async data.
		try {
			if (!window.socket) {
				await connectSocket();
			}

			let mainData = await API.get('/chats', {
				params: {
					results: 1,
					inc: 'chats,online',
				},
			}); // Парсим резульатты.

			console.log(mainData);

			let chatsData = mainData.data.chats;
			let onlineData = mainData.data.online; // Обновляем стейт и ререндерим наш компонент.

			this.setState({
				...this.state,
				...{
					isLoading: false,
					chats: chatsData,
					online: onlineData,
				},
			});
		} catch (err) {
			console.log('[DEBUG]: /chats error');
			console.log(err);
			const { history } = this.props;
			history.push('/login');
		}
	}
}
