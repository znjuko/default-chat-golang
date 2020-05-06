import React from 'react';
import PropTypes from 'prop-types';
import { Card } from 'shards-react';
import { Link } from 'react-router-dom';

export default class Chats extends React.Component {
	render() {
		return (
			<div className='d-flex mt-2 flex-column'>
				{this.props.chats.map((chat, index) => {
					return (
						<Link to={`/chats/${chat.chatId}`} key={chat.chatId}>
							<Card className='d-flex flex-column p-3 mx-auto m-1 w-100'>
								<h4 className='h4 text-uppercase d-flex mb-0 text-decoration-none'>
									{chat.chatName}
								</h4>
								<hr></hr>
								<div className='d-flex flex-column justify-content-around'>
									<h5 className='h5 text-decoration-none'>
										{chat.chatAuthor}
									</h5>
									<p className='lead ml-5 text-decoration-none'>
										{chat.chatMsg}
									</p>
								</div>
							</Card>
						</Link>
					);
				})}
			</div>
		);
	}
}

Chats.propTypes = {
	chats: PropTypes.array,
};
