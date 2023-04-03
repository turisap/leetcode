const primeNumbersCount = (n) => {
  const table = new Array(n).fill(true).fill(false, 0, 2);
  let c = n - 2;

  for (let i = 2; i < Math.sqrt(n); i++) {
    if (table[i]) {
      for (let j = i + i; j < n; j += i) {
        table[j] = false;
        c--;
      }
    }
  }

  return table.filter((val) => val).length;
};

module.exports = primeNumbersCount;
