import React from 'react';
import { Card, CardImg, Button } from 'shards-react';

export default class Emojie extends React.Component {
	render() {
		return (
			<div className='d-flex mt-2 w-100 flex-wrap align-items-start'>
				{this.props.emojies.map((emojie, index) => {
					return (
						<Card
							className='mr-1 ml-1 mb-2'
							key={index}
							style={{ maxWidth: '72px', maxHeight: '72px' }}
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
