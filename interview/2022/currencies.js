const  input = [
  ['usd', 'buy', 1000],
  ['usd', 'sell', 300],
  ['rub', 'sell', 1500],
  ['nzd', 'buy', 700],
  ['rub', 'buy', 400],
  ['rub', 'buy', 1200],
  ['nzd', 'sell', 420],
  ['nzd', 'buy', 333],
  ['usd', 'sell', 11150],
  ['kzh', 'sell', 500],
]


const calcAllOperations = (deals) => {
  const dealsMap = {}

  for(let deal of deals) {
    const [currency, operation, amount] = deal

    if(!dealsMap[currency]) {
      dealsMap[currency]  = {buy: 0, sell: 0}
    } 
 
    dealsMap[currency][operation] += amount
   
  }

  return Object.entries(dealsMap).map(([currency, operations]) => ({[currency]: [operations.buy, operations.sell]}))
}

console.log(calcAllOperations(input))
