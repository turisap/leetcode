const fibonachiRecursion = n => {
  if(n < 2) return n

  return fibonachiRecursion(n-1) + fibonachiRecursion(n-2)
}

console.log(fibonachiRecursion(10))


const fibLoop = n => {
  let a = 0, b = 1;

  for(let i = 2; i <= n; i++){
     let current = a + b; // or without current [a, b] = [b, a + b]
     [a, b] = [b, current]
  }

  return b;
}

console.log(fibLoop(10))
