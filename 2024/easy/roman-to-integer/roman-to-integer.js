// CHECKLIST
// Write a few test examples. This ensures that I understood the task

/*
- Think out loud the algorithm and the way (sliding window, two pointers) you want to implement it.
    - declare all variables you need
    - body of the algorithm
    - returning result
*/

// Think about bad input (at least ask about this)

const singlesMap = {
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

const rTI = (roman) => {
  let res = 0;
  let c = 0;

  while (c < roman.length) {
    if (doublesMap[roman[c] + roman[c + 1]]) {
      res += doublesMap[roman[c] + roman[c + 1]];
      c += 2;
      continue;
    }

    res += singlesMap[roman[c]];
    c++;
  }

  return res;
};

module.exports = rTI;
