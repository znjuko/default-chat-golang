import React from 'react';
import API from '../utils/API';

import OnlineUsers from './OnlineUsers';
import Chats from './Chats';

export default class Messanger extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			isLoading: true,
			chats: [
				{
					chatName: 'chat',
					chatMsg: 'asflasgflksdgaljhsglkdsaglksdglka',
					chatAuthor: 'Viici',
				},
				{
					chatName: 'chat',
					chatMsg: 'asflasgflksdgaljhsglkdsaglksdglka',
					chatAuthor: 'Viici',
				},
			],
			online: [
				{ login: 'Hello', id: 1 },
				{ login: 'Hello', id: 2 },
				{ login: 'Hello', id: 3 },
				{ login: 'Hello', id: 4 },
				{ login: 'Hello', id: 5 },
			],
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
						<Chats chats={this.state.chats} />
					</div>
					<div className='col col-md-3 '>
						<h4 className='h4 mt-2 text-left'>
							Пользователи онлайн
						</h4>
						<hr />
						<OnlineUsers online={this.state.online} />
					</div>
				</div>
			</div>
		);
	}

	async componentDidMount() {
		// Load async data.
		try {
			let mainData = await API.get('/chats', {
				params: {
					results: 1,
					inc: 'chats,online',
				},
			}); // Парсим резульатты.

			let chatsData = mainData.data.results[0].chats;
			let onlineData = mainData.data.results[0].online; // Обновляем стейт и ререндерим наш компонент.

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
			const { history } = this.props;
			history.push('/login');
		}
	}
}
