Function.prototype.myBind = function(ctx, ...args){
  if(!args.length){
    return (next) => this.call(ctx, next)
  }

  return (...next) => this.call(ctx, ...[...args, ...next])
}

const add = function (number) {
  console.log('context', this.foo)
  console.log(number + 10)
}

const obj= { foo: 'bar' }

const preAdd = add.myBind(obj)

preAdd(5) // 15

const addMany = function(number1, number2) {
  console.log(this.foo)
  console.log(2 + number1 + number2)
}

const preAddMany = addMany.myBind(obj, 3)

preAddMany(4) // 9
