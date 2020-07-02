const prime_factors = (n) => {
  const factors = [];

  while (n % 2 === 0) {
    n /= 2;
    factors.push(2);
  }

  for (let i = 3; i <= Math.sqrt(n); i += 2) {
    while (n % i === 0) {
      n /= i;
      factors.push(i);
    }
  }

  if (n !== 1) factors.push(n);

  return factors.join(" ");
};

module.exports = prime_factors;
