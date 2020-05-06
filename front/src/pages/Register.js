import React from 'react';
import PropTypes from 'prop-types';

import { Link } from 'react-router-dom';

import { Card } from 'shards-react';
import Input from '../mixins/input';
import Error from '../mixins/error';

import API from '../utils/API';

class Register extends React.Component {
	constructor() {
		super();
		this.state = {
			login: '',
			password: '',
			passwordRepeat: '',
			err: '',
			errText: 'Error text',
			errSame: '',
			buttonBlock: '',
		};
	}

	render() {
		const { isLoading } = this.props;

		const linkToChats = <Link to='/chats'>Регистрация</Link>;

		const userDetails = (
			<div className='container-fluid d-inline-flex flex-column justify-content-center'>
				<h1 className='display-4 mb-4'>Регистрация</h1>
				<form onSubmit={this.handleSubmit}>
					<Input
						icon='far fa-user'
						type='text'
						placeholder='Username'
						name='login'
						onChange={this.onChange}
						className={`form-control ${this.errorClass(
							this.state.err,
						)}`}
					/>
					<Input
						icon='fas fa-key'
						type='password'
						placeholder='Password'
						name='password'
						onChange={this.onChange}
						className={`form-control ${this.errorClass(
							this.state.errSame,
						)}`}
					/>
					<Input
						icon='fas fa-key'
						type='password'
						placeholder='Repeat password'
						name='passwordRepeat'
						onChange={this.onChange}
						className={`form-control ${this.errorClass(
							this.state.errSame,
						)}`}
					/>

					<Error
						err={this.state.errText}
						className={`text-danger ${this.errorHidClass(
							this.state.err,
							this.state.errSame,
						)}`}
					/>
					<div className='container-fluid d-inline-flex justify-content-center ml-1 mr-1 mb-2 '>
						<Link to='login' activeClassName='active'>
							<button
								type='button'
								class='btn btn-outline-info mr-4'>
								Уже есть аккаунт?
							</button>
						</Link>
						<button
							disabled={this.state.buttonBlock}
							type='submit'
							className='btn btn-primary d-inline-flex justify-content-center'>
							Зарегистрироваться
						</button>
					</div>
				</form>
			</div>
		);

		const loadingMessage = (
			<span className='d-flex m-auto'>Loading...</span>
		);

		return (
			<Card
				className='mx-auto m-5 text-center p-5 align-middle'
				style={{ minWidth: '600px' }}>
				{isLoading ? loadingMessage : userDetails}
			</Card>
		);
	}

	onChange = (e) => {
		this.setState({ [e.target.name]: e.target.value });
		if (
			e.target.name === 'passwordRepeat' &&
			e.target.value !== this.state.password
		) {
			console.log('Пароли должны совпадать!');
			this.setState({ errText: 'Пароли должны совпадать!' });
			this.setState({ errSame: 'Пароли должны совпадать!' });
			this.setState({ buttonBlock: 'Пароли должны совпадать!' });
		} else if (
			e.target.name === 'passwordRepeat' &&
			e.target.value === this.state.password
		) {
			this.setState({ errSame: '' });
			this.setState({ buttonBlock: '' });
		}
	};

	handleSubmit = async (event) => {
		event.preventDefault();
		// get our form data out of state
		const { login, password, passwordRepeat } = this.state;

		if (password === '' || passwordRepeat == '' || login == '') {
			console.log('Заполните поля!');
			this.setState({ errText: 'Заполните все поля!' });
			this.setState({ err: 'Заполните поля!' });
			this.setState({ errSame: 'Заполните поля!' });
		} else {
			try {
				await API.post('/registration', { login, password });
				const { history } = this.props;
				history.push('/chats');
			} catch (error) {
				console.log('[DEBUG]: Ответ сервера на /registration');
				console.log(error.response);
				this.setState({ err: error.response.status });
				this.setState({
					errText: 'Пользователь с таким login уже существует!',
				});
			}
		}
	};

	errorClass(error) {
		return error.length === 0 ? '' : 'border border-danger';
	}
	errorHidClass(errorOne, errorTwo) {
		return errorOne.length === 0 && errorTwo.length === 0
			? 'invisible'
			: 'visible';
	}
}

Register.propTypes = {
	isLoading: PropTypes.bool,
	login: PropTypes.string,
	password: PropTypes.string,
	passwordRepeat: PropTypes.string,
};

export default Register;
