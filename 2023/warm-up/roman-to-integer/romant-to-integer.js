const numbersMap = {
  I: 1,
  V: 5,
  X: 10,
  L: 50,
  C: 100,
  D: 500,
  M: 1000,
};

const doublesMap = {
  IV: 4,
  IX: 9,
  XL: 40,
  XC: 90,
  CD: 400,
  CM: 900,
};
const romanToInteger = (val) => {
  const chars = val.split("");
  let pointer = 0;
  let res = 0;

  while (pointer < chars.length) {
    const curr = chars[pointer];
    const next = chars[pointer + 1];

    if (isDouble(curr, next)) {
      res += doublesMap[`${curr}${next}`];
      pointer += 2;
    } else {
      res += numbersMap[chars[pointer]];
      pointer++;
    }
  }

  return res;
};

const isDouble = (curr, next) => {
  return Boolean(doublesMap[curr + next]);
};

romanToInteger("MCMXCIV");
// romanToInteger("CIV");

module.exports = {
  romanToInteger,
  isDouble,
};
