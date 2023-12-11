const pr = new Promise(res => {
  setTimeout(() => res(5), 1000)
})

pr
.then(r => r + 1)
.then(r => {
  throw new Error ('ololo')

  return r * 2
})
.catch(err => err)
// logs res is Error: ololo
.then(r => console.log('res is', r))