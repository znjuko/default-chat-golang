import React from 'react';
import PropTypes from 'prop-types';

import { Link } from 'react-router-dom';

import { Card } from 'shards-react';
import Input from '../mixins/input';
import Error from '../mixins/error';

import API from '../utils/API';

class Login extends React.Component {
	constructor() {
		super();
		this.state = {
			login: '',
			password: '',
			err: '',
		};
	}

	render() {
		const { isLoading } = this.props;
		const userDetails = (
			<div className='container-fluid d-inline-flex flex-column justify-content-center'>
				<h1 className='display-4 mb-4'>Вход</h1>
				<form onSubmit={this.handleSubmit}>
					<Input
						icon='far fa-user'
						type='text'
						placeholder='Логин'
						name='login'
						onChange={this.onChange}
						className={`form-control ${this.errorClass(
							this.state.err,
						)}`}
					/>
					<Input
						icon='fas fa-key'
						type='password'
						placeholder='Пароль'
						name='password'
						onChange={this.onChange}
						className={`form-control ${this.errorClass(
							this.state.err,
						)}`}
					/>

					<Error
						err='Упс... Такой пользователь не найден('
						className={`text-danger ${this.errorHidClass(
							this.state.err,
						)}`}
					/>
					<div className='container-fluid d-inline-flex justify-content-center ml-1 mr-1 mb-2 '>
						<button
							type='submit'
							className='btn btn-primary d-inline-flex mr-4 justify-content-center'>
							Войти
						</button>
						<Link to='registration' activeClassName='active'>
							<button
								type='button'
								className='btn btn-outline-info'>
								Регистрация
							</button>
						</Link>
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
	};

	handleSubmit = async (event) => {
		event.preventDefault();
		// get our form data out of state
		const { login, password } = this.state;

		try {
			await API.post('/login', { login, password });
			const { history } = this.props;

			history.push('/chats');
		} catch (error) {
			console.log('[DEBUG]: Ответ сервера на /login');
			console.log(error.response);
			this.setState({ err: error.response.status });
		}
	};

	errorClass(error) {
		return error.length === 0 ? '' : 'border border-danger';
	}
	errorHidClass(error) {
		return error.length === 0 ? 'invisible' : 'visible';
	}
}

Login.propTypes = {
	isLoading: PropTypes.bool,
	login: PropTypes.string,
	password: PropTypes.string,
};

export default Login;
