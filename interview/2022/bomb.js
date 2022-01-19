function Bomb(timer) {
  this.setTime = Date.now() + timer
  this.timerId = null

  this.init()
}

Bomb.prototype = {
  checkTime: function (){
    console.log('pik-pik')
    if(this.setTime <= Date.now()){
      console.log('booom')

      clearInterval(this.timerId)
    }
  },
  init: function (){
    this.timerId = setInterval(this.checkTime.bind(this), 100)
  }
}

const bomb = new Bomb(2000)
