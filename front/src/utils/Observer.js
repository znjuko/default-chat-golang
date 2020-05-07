import { Observable } from 'rxjs';

export let observer = Observable.create((obs) => {
	console.log('Observable created');

	return () => {
		console.log('Observable destroyed');
	};
});
