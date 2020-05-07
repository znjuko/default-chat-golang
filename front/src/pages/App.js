// src/App.js

import React from 'react';

import { Route, Switch, Redirect, withRouter } from 'react-router-dom';

import Login from './Login';
import Register from './Register';
import Messanger from './Messanger';
import Messages from './Messages';

class App extends React.Component {
	render() {
		const { history } = this.props;

		return (
			<div
				className='d-flex justify-content-center align-items-start overflow-hidden'
				style={{ minHeight: '100vh' }}>
				<Switch>
					<Route history={history} path='/login' component={Login} />
					<Route
						history={history}
						path='/registration'
						component={Register}
					/>
					<Route
						history={history}
						path='/chats'
						component={ChatsRout}
					/>
					<Redirect from='/' to='/chats' />
				</Switch>
			</div>
		);
	}
}

export default withRouter(App);

function ChatsRout(props) {
	const { history } = props;
	return (
		<Switch>
			<Route
				history={history}
				exact
				path='/chats'
				component={Messanger}
			/>
			<Route
				history={history}
				path='/chats/:chatId'
				component={Messages}
			/>
		</Switch>
	);
}
