const primeList = number => {
  const seive = []
  const primes = []

  for (let i = 2; i <= number; i++){
    if(!seive[i]){
      primes.push(i)
    }

    for(let j = i * 2; j <= number; j += i){
      seive[j] = true;
    }
  }

  return primes
}

const start = Date.now()
console.log(primeList(10000000))
console.log((Date.now()  -start) / 1000)

// bad is 2.63
// good is 1.24
