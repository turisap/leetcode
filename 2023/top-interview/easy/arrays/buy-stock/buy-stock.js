const buyStock = (arr) => {
  if (!arr.length) return 0;
  let diffs = 0;

  for (let i = 0; i < arr.length; i++) {
    if (arr[i] > arr[i - 1]) {
      diffs += arr[i] - arr[i - 1];
    }
  }

  return diffs;
};

module.exports = buyStock;
