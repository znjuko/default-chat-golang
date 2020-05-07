import React from 'react';
import API from '../utils/API';
import connectSocket from '../utils/Socket';
import Emojie from '../mixins/emojie';
import { Link, useParams } from 'react-router-dom';
import Message from '../mixins/Message';
import AddEmojie from '../mixins/addEmojie';

import {
	Popover,
	PopoverBody,
	PopoverHeader,
	Button,
	Modal,
	ModalBody,
	ModalHeader,
} from 'shards-react';

export default class Messages extends React.Component {
	constructor(props) {
		super(props);

		const {
			match: { params },
		} = this.props;

		this.insertEmojie = this.insertEmojie.bind(this);
		this.closeAddSticker = this.closeAddSticker.bind(this);
		this.state = {
			chatId: parseInt(params.chatId),
			chatInfo: {
				chatName: 'chat',
			},
			messages: [],
			textMessage: '',
			showEmojie: false,
			allEmojies: [],
			openAddEmojie: false,
		};
	}

	render() {
		return (
			<div
				className='container-fluid flex-column align-items-center h-100 mt-2 d-flex justify-content-between'
				style={{ minHeight: '90vh' }}>
				<div
					className='h-100  w-75 d-flex justify-content-center'
					style={{ height: '20%' }}>
					<div className='h-100 d-flex justify-content-between align-items-center w-100 border-bottom border-white'>
						<button
							onClick={this.onBack}
							activeclassname='active'
							className='d-inline-flex btn btn-outline-info btn-pill'>
							<i className='fas fa-undo-alt'></i>
						</button>
						<h4 className='h4 mt-2  d-inline-flex'>
							{this.state.chatInfo.chatName}
						</h4>
					</div>
					<hr />
				</div>

				<div
					className='w-75 d-flex justify-content-center position-fixed-bottom mx-auto'
					style={{ minHeight: '80vh', marginBottom: '80px' }}>
					<Message
						history={this.props.history}
						messages={this.state.messages}></Message>
				</div>

				<div
					className='w-75 d-flex justify-content-center fixed-bottom mx-auto'
					style={{ maxHeight: '20%' }}>
					<div className='input-group mb-3'>
						<div className='input-group-prepend'>
							<button
								className='btn btn-light'
								type='button'
								id='emojie'
								name='emojie'
								style={{ fontSize: '2rem' }}
								onClick={this.onEmojie}>
								<i
									className='fas fa-smile-wink'
									name='emojie2'
									onClick={this.onEmojie}></i>
							</button>
							<Popover
								placement='top'
								open={this.state.showEmojie}
								toggle={this.onEmojie}
								target='#emojie'
								className='overflow-hidden'
								style={{
									minHeight: '400px',
									minWidth: '280px',
									maxHeight: '400px',
									overflowY: 'scroll',
								}}>
								<PopoverHeader className='d-flex justify-content-between align-items-center'>
									Стикеры
									<button
										onClick={this.addSticker}
										activeClassName='active'
										className='d-inline-flex btn btn-outline-info btn-pill'>
										<i className='fas fa-plus-circle'></i>
									</button>
								</PopoverHeader>
								<PopoverBody>
									<Emojie
										history={this.props.history}
										chatId={this.state.chatId}
										emojies={this.state.allEmojies}
										func={this.insertEmojie}></Emojie>
								</PopoverBody>
							</Popover>
							<AddEmojie
								open={this.state.openAddEmojie}
								closeAddSticker={
									this.closeAddSticker
								}></AddEmojie>
						</div>
						<textarea
							ref='notes'
							type='text'
							className='form-control'
							placeholder=''
							aria-label=''
							aria-describedby='basic-addon1'
							style={{ resize: 'none' }}
							onChange={this.onChangeMessage}></textarea>
						<div className='input-group-append'>
							<button
								className='btn btn-primary'
								type='button'
								onClick={this.onSendMessage}>
								Отправить
							</button>
						</div>
					</div>
				</div>
			</div>
		);
	}

	closeAddSticker() {
		this.setState({ openAddEmojie: false });
	}

	insertEmojie(text) {
		console.log(text);
		var notes = this.refs.notes;
		notes.value += text;
		this.setState({ textMessage: notes.value });
	}

	onEmojie = (e) => {
		console.log(e.target);
		console.log(e.target.className);
		if (
			e.target.className === 'fas fa-smile-wink' ||
			e.target.className === 'btn btn-light'
		)
			this.setState({
				showEmojie: !this.state.showEmojie,
			});
	};

	addSticker = (e) => {
		this.setState({ openAddEmojie: true });
	};

	onBack = (e) => {
		const { history } = this.props;
		history.push('/chats');
	};

	onChangeMessage = (e) => {
		console.log(e.target.value);
		this.setState({ textMessage: e.target.value });
	};

	onSendMessage = (e) => {
		if (this.state.textMessage !== '') {
			console.log({
				txt: this.state.textMessage,
				chatId: this.state.chatId,
			});
			window.socket.send(
				JSON.stringify({
					txt: this.state.textMessage,
					chatId: this.state.chatId,
				}),
			);
			var notes = this.refs.notes;
			notes.value = '';
			this.setState({ textMessage: '' });
		}
	};

	async componentDidMount() {
		// Load async data.
		try {
			if (!window.socket) {
				await connectSocket();
			}

			// data: "{"event":"new message","message":{"author":"poi","txt":"sdga","chatName":"Snow","chatId":9,"emojies":null}}↵"
			window.socket.onmessage = function (event) {
				console.log('[WS] Message', event);
				var data = JSON.parse(event.data);
				console.log(data);
				console.log(data.message);
				if (data.message.chatId === this.state.chatId) {
					let newMsg = {
						author: data.message.author,
						txt: data.message.txt,
						chatName: data.message.chatName,
						chatId: data.message.chatId,
						emojies: data.message.emojies,
					};

					this.setState({
						messages: [newMsg, ...this.state.messages],
					});
					console.log('[Add msg]');
				}
			}.bind(this);

			let mainData = await API.get(`/chats/${this.state.chatId}`, {
				params: {
					results: 1,
					inc: 'chatInfo,messages',
				},
			}); // Парсим резульатты.

			console.log(mainData);

			let chatInfo = mainData.data.chatInfo;
			let messages = mainData.data.messages;
			let allEmojies = mainData.data.allEmojies; // Обновляем стейт и ререндерим наш компонент.

			console.log('[DEBUG]: Chat answer');
			console.log(mainData.data);
			this.setState({
				...this.state,
				...{
					chatInfo: chatInfo,
					messages: messages,
					allEmojies: allEmojies,
				},
			});
		} catch (err) {
			console.log(`[DEBUG]: /chats/${this.state.chatId} error`);
			console.log(err);
		}
	}
}
