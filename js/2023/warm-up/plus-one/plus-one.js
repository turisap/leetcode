const plusOne = (digits) => {
  let i = digits.length - 1;
  digits[i] = digits[i] + 1;

  while (digits[i] >= 10) {
    digits[i] = digits[i] % 10;

    if (i === 0) {
      digits.unshift(1);
      break;
    }

    digits[i - 1] = digits[i - 1] + 1;
    i--;
  }

  return digits;
};

module.exports = plusOne;
