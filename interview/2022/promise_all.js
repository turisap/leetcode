function MyPromiseAll (args) {
  this.args = args
  this.states = []
  this.result = []
  this.successCb = []
  this.errorCb = []

  args.map(arg => {
    arg
      .then((res) => this.handleResolve(res))
      .catch(console.error)
  })
}

MyPromiseAll.prototype =  {
  hadleCatch: function(err) {
    this.error = error
    this.errorCb.forEach(cb => cb(err))
  },

  handleResolve: function (res) {
    this.states.push('success')
    this.result.push(res)

    if(this._allResolved()){
      this.successCb.forEach(cb => cb(this.result))
    }
  },
  
  then: function(cb){
    if(this._allResolved()){
      cb(this.result)
    } else {
      this.successCb.push(cb)
    }

    return this
  },

  catch: function(cb){
    if(this.error) {
      cb(this.error)
    } else {
      this.errorCb.push()
    }

    return this
  },

  _allResolved: function (){
    return this.states.every(state => state === 'success') && this.states.length === this.args.length && !this.error
  },
}

const promise1 = new Promise((res) => setTimeout(() => res(3), 3000))
const promise2 = new Promise((res) => setTimeout(() => res('gogo baby'), 1000))
const promise3 = new Promise((res, rej) => setTimeout(() => rej('oh no'), 1500))
const pr = new MyPromiseAll([promise1, promise2, promise3])

pr
  .then(console.log)
  .catch(console.err)
