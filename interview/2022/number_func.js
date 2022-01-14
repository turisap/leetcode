const multi = arg => ({ arg, operation: '*' })
const divide = arg => ({ arg, operation: '/' })
const plus = arg => ({arg, operation: '+'})
const minus = arg => ({arg, operation: '-'})

const calculate = (number, set) => {
  switch(set.operation){
    case('*'): 
      return number * set.arg
    
    case('/'): 
      return number / set.arg
    
    case('+'):
      return number + set.arg

    case('-'):
      return number - set.arg
  }
}

function three (set) {
  const current = 3

  if(!set) return current

  return calculate(current, set)
}

function two (set) {
  const current = 2;

  if(!set) return current

  return calculate(current, set)
}

console.log(two(multi(three()))) // WORKS BOTH WAYS
console.log(three(divide(two())))
console.log(two(plus(three())))

