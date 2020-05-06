import React from 'react';
import PropTypes from 'prop-types';
import { Card } from 'shards-react';

export default class Chats extends React.Component {
	render() {
		return (
			<div className='d-flex mt-2 flex-column'>
				{this.props.chats.map((chat, index) => {
					return (
						<Card className='d-flex flex-column p-3 mx-auto m-1 w-100'>
							<h4 className='h4 text-uppercase d-flex mb-0'>
								{chat.chatName}
							</h4>
							<hr></hr>
							<div className='d-flex flex-column justify-content-around'>
								<h5 className='h5'>{chat.chatAuthor}</h5>
								<p class='lead ml-5'>{chat.chatMsg}</p>
							</div>
						</Card>
					);
				})}
			</div>
		);
	}
}

Chats.propTypes = {
	chats: PropTypes.array,
};
