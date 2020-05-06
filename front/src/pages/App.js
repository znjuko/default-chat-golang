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
					<Redirect from='/' to='/login' />
				</Switch>
			</div>
		);
	}
	// Место для запроса на сервер
	/* async componentDidMount() {
		// Load async data.
		let userData = await API.post('/login', {
			params: {
				results: 1,
				inc: 'name,password',
			},
		});
		// Parse results
		userData = userData.data.results[0];

		// Update state with new data.
		const name = `${userData.name}`;

		// Re-render our component.
	} */
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
