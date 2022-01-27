var anotherObject = {
  a: 2
}

var myObject = Object.create(anotherObject)

console.log(anotherObject.a)  // 2
console.log(myObject.a)       // 2

console.log(anotherObject.hasOwnProperty('a')) // true
console.log(myObject.hasOwnProperty('a'))      // false

// incrementing prototype on anotherObject
// and creating a on myObject
myObject.a++ 

console.log(anotherObject.a) // 2
console.log(myObject.a)      // 3
console.log(myObject.hasOwnProperty('a')) // true
