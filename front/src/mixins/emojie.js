import React from 'react';
import { Card, CardImg, Button } from 'shards-react';

export default class Emojie extends React.Component {
	render() {
		return (
			<div className='d-flex mt-2 w-100 flex-column-reverse align-items-start justify-content-end'>
				{this.props.emojies.map((emojie, index) => {
					return (
						<Card
							key={index}
							style={{ maxWidth: '80px', maxHeight: '80px' }}
							onClick={() => {
								this.props.func(emojie.phrase);
							}}
							name={emojie.about}>
							<CardImg top src={emojie.url} />
						</Card>
					);
				})}
			</div>
		);
	}

	onClick = (e) => {
		console.log('realClick');
		this.props.updateData('hello');
	};
}
/*
 */
