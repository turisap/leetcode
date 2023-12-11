const prices1 = [7, 1, 5, 4, 8, 3]
const prices2 = [7, 1, 5, 4, 8, 10]


const maxProfitMulti = (prices) => {
  let result = 0;

  for(let c = 0; c < prices.length; c++){
    const current = prices[c];
    const prev = prices[c - 1]

    if(current > prev){
      result += current - prev
    }
  }

  return result
}

console.log(maxProfitMulti(prices1)) // logs 8
console.log(maxProfitMulti(prices2)) // logs 10

