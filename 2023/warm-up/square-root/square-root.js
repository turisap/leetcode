const squareRoot = (num) => {
  if (num <= 2) {
    return num;
  }

  let left = 2;
  let right = Math.floor(num / 2);

  while (left <= right) {
    const middle = Math.floor((left + right) / 2);
    const res = middle * middle;

    if (res === num) {
      return middle;
    }

    if (res < num) {
      left = middle + 1;
    }

    if (res > num) {
      right = middle - 1;
    }
  }

  return Math.floor(right);
};

module.exports = squareRoot;
