function promiseAll(promises) {
	return new Promise((resolve, reject) => {
		const results = [];
		let resolvedCount = 0;

		promises.forEach((promise, index) => {
		  if(typeof promise === 'object' && 'then' in promise){
			promise
				.then((result) => {
					results[index] = result;

					resolvedCount++;

					if (resolvedCount === promises.length) {
						resolve(results);
					}
				})
				.catch((err) => reject(err));
		  } else {
		    results[index] = promise;
		    resolvedCount++
		  }
		});
	});
}

promiseAll([
	new Promise((resolve) => {
		setTimeout(() => resolve('foo'), 5000)
	}),

	new Promise((resolve, reject) => {
		setTimeout(() => resolve('bar'), 1000);
	}),

	5,
	new Promise((resolve, reject) => {
		setTimeout(() => {
			Math.round(Math.random() * 10) % 2 === 0
				? resolve('baz')
				: reject(new Error());
		}, 300);
	}),
])
	.then((res) => console.log('RESOLVED: ', res))
	.catch((err) => console.log('REJECTED: ', err));
