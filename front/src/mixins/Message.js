import React from 'react';
import PropTypes from 'prop-types';
import { Card } from 'shards-react';
import { Link } from 'react-router-dom';

export default class Message extends React.Component {
	render() {
		return (
			<div className='d-flex mt-2 w-100 flex-column-reverse align-items-start justify-content-end'>
				{this.props.messages.map((message, index) => {
					return (
						<Card
							key={index}
							className='d-flex-inline flex-column p-3  mw-100'
							style={{ minWidth: '96%', margin: '2%' }}>
							<h6 className='h6 text-uppercase d-flex mb-0 text-decoration-none'>
								{message.author}
							</h6>
							<hr></hr>
							<div className='d-flex flex-column justify-content-around'>
								<div className='lead text-decoration-none'>
									{this.parseMessage(
										message.txt,
										message.emojies,
									)}
								</div>
							</div>
						</Card>
					);
				})}
			</div>
		);
	}

	parseMessage(txt, emojies) {
		// console.log(txt);
		let txtArr = txt.split(' ');
		// console.log(txtArr);
		return (
			<p>
				{txtArr.map((element) => {
					let res = [];
					if (
						element[0].startsWith(':') &&
						element[element.length - 1].endsWith(':')
					) {
						console.log(element);

						if (emojies) {
							res = emojies.filter((el) => {
								if (el.phrase === element) {
									console.log(el);
									return true;
								}
							});
							console.log(res);
						}
					}

					if (res.length === 0) {
						element += ' ';
						return element;
					} else {
						console.log(res[0].url);
						return (
							<img
								style={{ maxWidth: '80px', maxHeight: '80px' }}
								src={res[0].url}
							/>
						);
					}
				})}
			</p>
		);
	}
}

Message.propTypes = {
	messages: PropTypes.array,
};

/* {txtArr.forEach((element) => {
	return { element };
})} */
/* */
