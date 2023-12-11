const summFactory = () => {
  let lastSumm = 0

  return function summ(digit) {
    if(digit){
      lastSumm += digit

      return (next) => summ(next)
    }

    return lastSumm
  }
}

const summ = summFactory()
const res = summ(1)(2)()

console.log(res) // logs 3
