const buyStock = (days) => {
  let currMin = days[0];
  let maxDiff = 0;

  for (let day of days) {
    const diff = day - currMin;
    if (diff > maxDiff) maxDiff = diff;

    if (day < currMin) currMin = day;
  }

  return maxDiff;
};

module.exports = buyStock;
