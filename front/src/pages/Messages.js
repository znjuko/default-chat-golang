import React from 'react';
import API from '../utils/API';

import { Link, useParams } from 'react-router-dom';

export default class Messages extends React.Component {
	constructor(props) {
		super(props);

		this.state = {
			chatInfo: {
				chatName: 'chat',
			},

			messages: [
				{
					author: 'Hello',
					txt: 'afanfjfljashfn;kasfnsd',
					emojies: [{ url: 'linf', phrase: 'asfasf' }],
				},
			],
		};
	}

	render() {
		return <div></div>;
	}
}
