const beeramid = (bonus, price) => {
  if (bonus === 0 || bonus < price) return 0;

  const maxCaps = Math.floor(bonus / price);
  let totalCaps = 1;
  let level = 1;

  do {
    level++;
    totalCaps += Math.pow(level, 2);
  } while (totalCaps <= maxCaps);

  return level - 1;
};
