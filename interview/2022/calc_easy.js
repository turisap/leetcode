const summ = (a, b) => a + b;
const mult = (a, b) => a * b;

const calculate = (operation) => {
  return function(a){
    return function(b){
      return operation(a,b)
    }
  }
}

console.log(calculate(summ)(2)(3)) // logs 5
console.log(calculate(mult)(2)(3)) // logs 6
