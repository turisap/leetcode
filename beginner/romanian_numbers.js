function decimal2Romanian(number) {
  let result = "";

  const NUMBERS = [
    ["M", 1000],
    ["CM", 900],
    ["D", 500],
    ["CD", 400],
    ["C", 100],
    ["XC", 90],
    ["L", 50],
    ["XL", 40],
    ["X", 10],
    ["IX", 9],
    ["V", 5],
    ["IV", 4],
    ["I", 1],
  ];

  NUMBERS.forEach((digit, i) => {
    const [romanian, decimal] = digit;

    const lettersNumber = Math.floor(number / decimal);
    number = number % decimal;

    if (lettersNumber > 0) {
      result += romanian.repeat(lettersNumber);
    }
  });

  return result;
}

module.exports = decimal2Romanian;
