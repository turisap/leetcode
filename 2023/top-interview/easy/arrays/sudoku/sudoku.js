const validSudoku = (arr) => {
  const hashSet = {};

  for (let r = 0; r < 9; r++) {
    for (let c = 0; c < 9; c++) {
      const currentEl = arr[r][c];

      if (currentEl === ".") continue;

      const rowS = {
        0: 1,
        1: 1,
        2: 1,
        3: 2,
        4: 2,
        5: 2,
        6: 3,
        7: 3,
        8: 3,
      };
      const square = `${rowS[r]}${rowS[c]}`;

      const hasThisNumber =
        hashSet[`row${r}`]?.[currentEl] ||
        hashSet[`col${c}`]?.[currentEl] ||
        hashSet[`sq${square}`]?.[currentEl];

      if (hasThisNumber) {
        return false;
      }

      hashSet[`row${r}`] = {
        ...hashSet[`row${r}`],
        [currentEl]: true,
      };
      hashSet[`col${c}`] = {
        ...hashSet[`col${c}`],
        [currentEl]: true,
      };
      hashSet[`sq${square}`] = {
        ...hashSet[`sq${square}`],
        [currentEl]: true,
      };
    }
  }

  return true;
};

module.exports = validSudoku;
