const coins = [1, 5, 10]

const getCoins = (amount) => {
  const res = {}

  let currentPointer = coins.length - 1
  const smallestCoint = coins[0]

  while(amount >= smallestCoint) {

    while(amount >= coins[currentPointer]){
      let currentCoin = coins[currentPointer]
      res[currentCoin] ? res[currentCoin]++ : res[currentCoin] = 1
      amount -= coins[currentPointer]
    }

    currentPointer--
  }

  return  [res, amount.toFixed(2)]
}

console.log(getCoins(27.7))
