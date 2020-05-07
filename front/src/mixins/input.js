import React from 'react';

export default function Input(props) {
	return (
		<div className='form-group col-md'>
			<div className='input-group'>
				<div className='input-group-prepend'>
					<span className='input-group-text' id='basic-addon1'>
						<i className={props.icon}></i>
					</span>
				</div>
				<input
					type={props.type}
					className={props.className}
					placeholder={props.placeholder}
					aria-label={props.placeholder}
					aria-describedby='basic-addon1'
					name={props.name}
					onChange={props.onChange}
				/>
			</div>
		</div>
	);
}
