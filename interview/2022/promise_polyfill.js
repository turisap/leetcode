function MyPromise(cb) {
  this.errorHandlers = []; 
  this.successHandlers = [];
  this.status = 'pending'

  cb.call(this, this.resolve.bind(this), this.reject.bind(this))
}

MyPromise.prototype = {
  resolve: function(result) {
    this.status = 'success'
    this.successHandlers.forEach(handler => handler(result))
  },

  reject: function(reason) {
    this.status = 'reject'
    this.errorHandlers.forEach(handler => handler(reason))
  },

  catch: function (cb){
    if(this.status === 'reject'){
      cb()	
    } else {
      this.errorHandlers.push(cb)
    }

    return this
  },

  then: function (cb) {
    if(this.status === 'success') {
      cb()
    } else {
      this.successHandlers.push(cb)
    }

    return this
  }
}

const prm =  new MyPromise((res, rej) => {
  setTimeout(() => rej('ololo'), 1000)
})

prm
  .then((res) => console.log(res))
  .then((res) => console.log(res))
  .then((res) => console.log(res))
  .then((res) => console.log(res))
  .catch((err) => console.log('err:', err))
  .catch((err) => console.log('err:', err))


