import React from 'react';
import API from '../utils/API';
import { Card } from 'shards-react';

export default class OnlineUsers extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			users: [],
			showChatCreate: false,
			chatName: '',
		};
	}
	render() {
		return (
			<div>
				<div className='list-group form'>
					{this.props.online.map((user, index) => {
						return (
							<Card className='d-flex mx-auto m-1 text-center p-2 align-middle list-group-item list-group-item-action'>
								<div className='d-flex justify-content-between align-items-center'>
									<label
										className='form-check-label'
										for='defaultCheck1'>
										{user.login}
									</label>
									<input
										className='form-check'
										type='checkbox'
										value={user.id}
										id='defaultCheck1'
										onChange={this.onChange}
									/>
								</div>
							</Card>
						);
					})}
				</div>

				<div
					className={`mt-4 d-flex ${this.addChatAdd(
						this.state.showChatCreate,
					)}`}>
					<div className='card mb-3'>
						<div className='card-header'>Создать чат</div>
						<div className='card-body text-primary'>
							<div class='input-group mb-3'>
								<input
									type='text'
									class='form-control'
									placeholder='Название чата'
									aria-label='Название чата'
									aria-describedby='basic-addon2'
									onChange={this.onInputChange}
								/>
								<div class='input-group-append'>
									<button
										class='btn btn-outline-primary'
										type='button'
										onClick={this.onCreate}>
										Создать
									</button>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		);
	}

	onInputChange = (e) => {
		this.setState({ chatName: e.target.value });
	};

	onChange = (e) => {
		console.log('show1');
		if (e.target.checked) {
			this.state.users.push(e.target.value);
			console.log('show2');
			console.log(this.state.users);
		} else {
			let index = this.state.users.indexOf(e.target.value);
			if (index > -1) {
				this.state.users.splice(index, 1);
			}
			console.log('show3');
			console.log(this.state.users);
		}

		if (this.state.users.length > 0) {
			this.setState({ showChatCreate: true });
		} else {
			this.setState({ showChatCreate: false });
		}
	};

	addChatAdd(val) {
		return val ? 'visible' : 'invisible';
	}

	onCreate = async (event) => {
		event.preventDefault();
		// get our form data out of state
		const { chatName, users } = this.state;

		try {
			let res = await API.post('/chats', { chatName, users });

			// TODO: Добавить редирект на созданный чат
			// const { history } = this.props;
			// history.push('/chats');
		} catch (error) {
			console.log('[DEBUG]: Ответ сервера на /chats');
			console.log(error.response);
		}
	};
}
