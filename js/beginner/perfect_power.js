const can_be_divided = (number, d) => {
  const res = { success: false, pow: 0 };
  let count = 0;

  while (number >= 1) {
    if (number === 1) {
      res.success = true;
      res.pow = count;
      break;
    }

    if (number / d < 1) res.success = false;

    count++;
    number = number / d;
  }

  return res;
};

const perfect_power = (n) => {
  let c = 2;
  let res = null;

  while (c < n) {
    const check = can_be_divided(n, c);
    if (check.success) {
      res = [c, check.pow];
      break;
    }
    c++;
  }

  return res;
};

module.exports = { can_be_divided, perfect_power };
