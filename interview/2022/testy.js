let a = function () {
  console.log('fn')
}

let b = a;

console.log(b) // [Function: a]

b = null

console.log(a) // [Function: a]

a = null

console.log(b) // null
