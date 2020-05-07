import React from 'react';

export default function Error(props) {
	return (
		<p className={props.className}>
			<small>{props.err}</small>
		</p>
	);
}
