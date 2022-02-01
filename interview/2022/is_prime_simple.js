const { isPrime } = require ('prime-number-boolean');

const myPrime = number => {

  if(number < 3 && number > 0){
    return true
  }

  if(number % 2 === 0) {
    return false
  }
  
  const limit = Math.floor(Math.sqrt(number))

  for(let i = 3; i <= limit; i += 2){
    if(number % i === 0) {
      return false
    }
  }

  return true;
}

let i = 1
//  while(true){
//    const prime = myPrime(i)
//    const thatPrime = Boolean(isPrime.thisNumber(i))
//
//    if(prime !== thatPrime){
//      console.log(`Error: Int:${i} ${prime} !== ${thatPrime}`)
//      break;
//    }
//
//    console.log('=============')
//    console.log(`OK: Int:${i} ${prime} === ${thatPrime.toString()}`)
//    i++
//  }

module.exports = myPrime
