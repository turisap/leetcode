const fetch = require('node-fetch'); // version 2

const fetchRetry = (url,config,  maxRetry) => {
  let counter = 1;
  let id = null;
  const delay = 3000;

  const getData = () => {
    fetch(url, config)
      .then(res => res.json())
      .then(json => {
	console.log('resp:', json)
	if(id){
	  clearTimeout(id)
	}
      })
      .catch(() => {
	console.log('retry')
	if(counter < maxRetry){
	    counter++
	    id = setTimeout(getData, delay)
	    return
	}

	console.log('could not get to the server')
      })
  }

  getData()
}

fetchRetry('https://api.github.com/user', {}, 5)
