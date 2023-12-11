const robHouse = (arr) => {
  if (!arr.length) return 0;
  if (arr.length === 1) return arr[0];
  if (arr.length === 2) return Math.max(arr[0], arr[1]);

  const table = [];

  table[0] = 0;
  table[1] = arr.shift();

  while (arr.length) {
    const curr = arr.shift();
    const nextMax = Math.max(curr + table.at(-2), table.at(-1));

    table.push(nextMax);
  }

  return table.at(-1);
};

module.exports = robHouse;
