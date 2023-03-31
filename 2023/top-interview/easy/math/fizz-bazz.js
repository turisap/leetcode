const fizzBazz = (n) => {
  const res = n >= 5 ? ["1", "2", "Fizz", "4", "Buzz"] : [];
  const start = n >= 5 ? 6 : 1;

  for (let i = start; i <= n; i++) {
    if (i % 3 === 0 && i % 5 === 0 && i >= 5) {
      res.push("FizzBuzz");
    } else if (i % 3 === 0) {
      res.push("Fizz");
    } else if (i >= 5 && i % 5 === 0) {
      res.push("Buzz");
    } else {
      res.push(i.toString());
    }
  }

  return res;
};

module.exports = fizzBazz;
