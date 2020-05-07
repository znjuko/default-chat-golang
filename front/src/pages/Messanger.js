import React from 'react';
import API from '../utils/API';

import OnlineUsers from './OnlineUsers';
import Chats from './Chats';
import connectSocket from '../utils/Socket';
import {
	Button,
	Modal,
	ModalBody,
	ModalHeader,
	FormTextarea,
	Form,
} from 'shards-react';

export default class Messanger extends React.Component {
	constructor(props) {
		super(props);
		this.toggle = this.toggle.bind(this);
		this.state = {
			isLoading: true,
			chats: [],
			online: [],
			open: false,
			mainText: '',
		};
	}

	render() {
		const { history } = this.props;
		return (
			<div
				className='container-fluid h-100 mt-2'
				style={{ minHeight: '90vh' }}>
				<div
					className='row h-100 d-flex justify-content-center'
					style={{ minHeight: '100vh' }}>
					<div className='col-md-6'>
						<div className='d-flex justify-content-between'>
							<h4 className='h4 mt-2 text-left'>Ваши чаты</h4>
							<button
								onClick={this.onBack}
								activeclassname='active'
								className='d-inline-flex btn btn-danger btn-pill'>
								<i
									className='fas fa-frown'
									style={{ fontSize: '1.5rem' }}></i>
							</button>
						</div>

						<hr />
						<Chats history={history} chats={this.state.chats} />
					</div>
					<div className='col col-md-3 '>
						<h4 className='h4 mt-2 text-left'>
							Пользователи онлайн
						</h4>
						<hr />
						<Button
							theme='info'
							onClick={this.openSendToAll}
							activeclassname='active'
							className='w-100'>
							<p className='lead mb-0'>
								Отправить сообщение всем
							</p>
						</Button>
						<Modal open={this.state.open} toggle={this.toggle}>
							<ModalHeader>
								Введите для всех сообщение
							</ModalHeader>
							<ModalBody>
								<Form onSubmit={this.handleSubmit}>
									<FormTextarea
										className='mb-4'
										placeholder='Введите ваше сообщение'
										rows='10'
										onChange={this.handleChange}
										style={{ resize: 'none' }}
									/>
									<Button
										type='submit'
										className='mt-1 w-100'>
										👉Отправить 👈
									</Button>
								</Form>
							</ModalBody>
						</Modal>
						<hr />
						<OnlineUsers
							history={history}
							online={this.state.online}
						/>
					</div>
				</div>
			</div>
		);
	}

	handleChange = (e) => {
		this.setState({ mainText: e.target.value });
	};
	handleSubmit = async (e) => {
		e.preventDefault();
		const { mainText } = this.state;
		try {
			await API.post('/all', { mainText });
			this.setState({ open: false });
		} catch (err) {
			console.log('[DEBUG]: error while all');
			console.log(err);
		}
	};

	onBack = async () => {
		try {
			await API.delete('/login');
			window.socket.complete();
			this.props.history.push('/login');
		} catch (err) {
			console.log('[DEBUG]: error while delete');
			console.log(err);
		}
	};

	toggle() {
		this.setState({
			open: !this.state.open,
		});
	}

	openSendToAll = () => {
		this.setState({
			open: !this.state.open,
		});
	};

	async componentDidMount() {
		// Load async data.
		try {
			if (!window.socket) {
				await connectSocket();
			}

			window.socket.subscribe(
				(val) => {
					if (val.event === 'online') {
						if (window.location.pathname === '/chats') {
							val = val.online;
							console.log(val);
							this.setState({ online: val });
						}
					} else if (val.event === 'new message') {
						if (window.location.pathname === '/chats') {
							this.props.history.push('/');
						}
					}
				},
				(err) => {
					return console.error('error:', err);
				},
				() => console.log('completed'),
			);

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
			this.props.history.push('/login');
		}
	}
}
