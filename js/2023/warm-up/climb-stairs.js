const climbStairs = (n) => {
  const memoResults = [];
  const baseClimbStairs = (a, b) => {
    console.log(memoResults);
    if (a > b) {
      return 0;
    }

    if (a === b) {
      return 1;
    }

    if (memoResults[a]) return memoResults[a];

    const next = baseClimbStairs(a + 1, b);
    const nextPlusOne = baseClimbStairs(a + 2, b);

    memoResults[a + 1] = next;

    return next + nextPlusOne;
  };

  return baseClimbStairs(0, n);
};

const climStairsDynamic = (n) => {
  const cache = {
    0: 1,
    1: 1,
  };

  let c = 2;

  while (c <= n) {
    cache[c] = cache[c - 1] + cache[c - 2];

    c++;
  }

  return cache[c - 1];
};

console.log("RESULT", climbStairs(6));
console.log("RESULT", climStairsDynamic(6));
