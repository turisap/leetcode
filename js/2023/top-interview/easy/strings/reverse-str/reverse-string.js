const reverseString = (str) => {
  let pL = 0;
  let pR = str.length - 1;

  while (pL < pR) {
    const temp = str[pL];
    str[pL] = str[pR];
    str[pR] = temp;

    pL++;
    pR--;
  }

  return str;
};

module.exports = reverseString;
