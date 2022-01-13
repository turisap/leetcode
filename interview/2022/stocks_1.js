const prices1 = [7, 1, 5, 3, 6, 4];
const prices2 = [7, 5, 3, 2, 1]
const prices3 = [10, 11, 4, 3, 2]

const maxProfit = (prices) => {
  let minPrice = prices[0];
  let maxDeal = 0;

  if(prices.length < 2) {
    return null
  }

  for(let price of prices){
    if(price < minPrice) {
      minPrice = price
      continue;
    }

    const currentDiff = price - minPrice
    if(currentDiff > maxDeal){
      maxDeal = currentDiff
    }
  }
  
  return maxDeal
}

console.log(maxProfit(prices1)) // 5
console.log(maxProfit(prices2)) // 0
console.log(maxProfit(prices3)) // 1


