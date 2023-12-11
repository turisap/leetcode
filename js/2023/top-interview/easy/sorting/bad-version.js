const solution = (cb) => {
  return function (n) {
    let left = 1;
    let right = n;

    if (left === right) return left;

    while (left <= right) {
      const middle = Math.floor((left + right) / 2);
      const isBad = cb(middle);

      if (isBad && !cb(middle - 1)) return middle;

      if (isBad) {
        right = middle - 1;
      } else {
        left = middle + 1;
      }
    }
  };
};

const isBadVersion = (bad) => (search) => {
  return search >= bad;
};

module.exports = { solution, isBadVersion };
