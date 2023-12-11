const pascalTriangle = (n) => {
  if (n === 1) {
    return [[1]];
  }

  const res = [[1], [1, 1]];

  for (let i = 2; i < n; i++) {
    const row = res[i - 1].concat(1);

    for (let c = 1; c < i; c++) {
      const prevRow = res[i - 1];

      row[c] = prevRow[c - 1] + prevRow[c];
    }

    res.push(row);
  }

  return res;
};

module.exports = pascalTriangle;
