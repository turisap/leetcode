const arrayIntersection = (num1, num2) => {
  num1.sort();
  num2.sort();

  const map = {};
  const res = [];
  const ceil = num1.at(-1);

  for (let i = 0; i < num1.length; i++) {
    const curr = num1[i];

    if (map[curr]) {
      map[curr]++;
    } else {
      map[curr] = 1;
    }
  }

  for (let i = 0; i < num2.length; i++) {
    const curr = num2[i];

    if (curr > ceil) break;

    if (map[curr]) {
      res.push(curr);
      map[curr]--;
    }
  }

  return res;
};

module.exports = arrayIntersection;
