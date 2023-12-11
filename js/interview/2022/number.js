Number.prototype.plus = function (input) {
  return this + input
}

Number.prototype.minus = function (input){
  return this - input
}

console.log(Number(3).plus(4).minus(2)) // logs 5

