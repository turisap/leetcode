const reverseIntegerStr = (num) => {
  const reversed = num.toString().split("").reverse();
  const number =
    reversed.at(-1) === "-"
      ? -Number(reversed.slice(0, -1).join(""))
      : Number(reversed.join(""));

  return Math.abs(number) > Math.pow(2, 31) - 1 ? 0 : number;
};

const reverseInteger = (num) => {
  let res = 0;
  let multi = 10;
  let positive = num >= 0;

  num = Math.abs(num);

  while (num > 0) {
    res = res * multi + (num % 10);

    num = Math.floor(num / 10);

    if (Math.abs(res) > Math.pow(2, 31)) {
      return 0;
    }
  }

  return positive ? res : -res;
};

module.exports = reverseInteger;
