import React from 'react';
import {
	Form,
	FormGroup,
	FormInput,
	Modal,
	ModalBody,
	ModalHeader,
	Button,
} from 'shards-react';

import API from '../utils/API';

import Error from '../mixins/error';

export default class AddEmojie extends React.Component {
	constructor(props) {
		super(props);

		this.state = {
			url: '',
			phrase: '',
			err: '',
			errText: 'Упс... Стикер с таким кодовым словом уже есть(',
		};
	}

	render() {
		return (
			<div>
				<Modal open={this.props.open}>
					<div className='w-100 d-flex justify-content-between align-items-center border-bottom border-white'>
						<ModalHeader className=' border-0'>
							Добавить стикер
						</ModalHeader>
						<button
							onClick={() => {
								this.props.closeAddSticker();
							}}
							activeClassName='active'
							className='d-inline-flex btn btn-outline-info btn-pill mr-4'>
							<i className=' d-inline-flex fas fa-times-circle'></i>
						</button>
					</div>

					<ModalBody>
						<Form onSubmit={this.handleSubmit}>
							<FormGroup>
								<FormInput
									className='mb-4'
									id='url'
									placeholder='Вставьте url на картинку'
									name='url'
									onChange={this.onChange}
								/>
								<FormInput
									id='phrase'
									name='phrase'
									placeholder='Введите кодовое слово'
									onChange={this.onChange}
								/>
								<Error
									err={this.state.errText}
									className={`text-danger text-center mt-2 ${this.errorHidClass(
										this.state.err,
									)}`}
								/>
								<Button type='submit' className='mt-1 w-100'>
									🤩 Добавить 🤩
								</Button>
							</FormGroup>
						</Form>
					</ModalBody>
				</Modal>
			</div>
		);
	}

	onChange = (e) => {
		this.setState({ [e.target.name]: e.target.value });
	};

	handleSubmit = async (e) => {
		e.preventDefault();
		// get our form data out of state
		const { url, phrase } = this.state;

		if (url.length === 0 || phrase.length === 0) {
			this.setState({ errText: 'Надо все заполнить... 😡' });
			this.setState({ err: this.state.errText });
		} else {
			try {
				await API.post('/emoji', { url, phrase });
				this.props.closeAddSticker();

				const { history } = this.props;
				history.push(`/login`);
				history.push(`/chats/${this.props.chatId}`);
			} catch (error) {
				console.log('[DEBUG]: Ответ сервера на /login');
				console.log(error.response);
				this.setState({ err: this.state.errText });
			}
		}
	};

	errorHidClass(error) {
		return error.length === 0 ? 'invisible' : 'visible';
	}
}
/*
 */
