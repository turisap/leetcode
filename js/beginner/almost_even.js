const almost_even = (number, divider) => {
  const res = [];
  const wholePart = Math.floor(number / divider);

  while (number > 0) {
    if (number >= wholePart) {
      res.push(wholePart);
      number -= wholePart;
    } else {
      break;
    }
  }

  for (let i = res.length - 1; number > 0; i--) {
    res[i]++;
    number--;
  }

  return res;
};

almost_even(20, 6);
