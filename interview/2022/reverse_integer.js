const reverseInteger = (n) => {
  const lowerBound = -Math.pow(2, 31)
  const upperBound = -lowerBound - 1;

  let reverse = 0;
  let rest = n

  while (Math.abs(rest) > 0) {
    reverse *= 10;
    reverse += Math.trunc(rest % 10)

    rest = Math.trunc(rest / 10)
  }

  return reverse > lowerBound && reverse < upperBound ? reverse : 0
}

console.log(reverseInteger(123))
console.log(reverseInteger(Math.pow(2, 29)))
