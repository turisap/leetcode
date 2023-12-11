const coins = {
  100: 1,
  50: 5,
  25: 6,
  10: 1,
  2: 1,
  1: 2
}

const greedyTough = (money) => {
  const res = {}
  const banknotes = Object.keys(coins)
  let pointer = banknotes.length - 1

  while (pointer >= 0) {
    const currentBanknote = banknotes[pointer]
    if(money >= currentBanknote && coins[currentBanknote] > 0) {
      // here is the trick without another loop
      const banknotesNumber = Math.min(Math.floor(money / currentBanknote), coins[currentBanknote])

      res[currentBanknote] = banknotesNumber
      money -= currentBanknote * banknotesNumber;
    }

    pointer--
  }

  return [res, money]
}

console.log(greedyTough(200))
