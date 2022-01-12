const summ = (args) => args.reduce((acc, curr) => acc + curr, 0)
const mult = (args) => args.reduce((acc, curr) => acc * curr, 1)


const calculateFactory = () => {
  let operation = null
  let operands = []

  return function calcInner(arg){
    if(typeof arg === 'function') {
      operation = arg
    } else if(typeof arg === 'number'){
      operands.push(arg)
    } 
    let canCalculate = !arg

    if(canCalculate){
      const result = operation(operands);
      // don't forget to reset values
      operation = null
      operands = []

      return result
    }
  
    return calcInner
  }
}

const calculate = calculateFactory()

const res1 = calculate(summ)(2)(3)(4)(5)() // logs 14
const res2 = calculate(mult)(2)(3)(4)(5)() // logs 120

console.log(res1)
console.log(res2)
