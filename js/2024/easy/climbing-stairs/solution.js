// CHECKLIST
// Write a few test examples. This ensures that I understood the task

/*
- Think out loud the algorithm and the way (sliding window, two pointers) you want to implement it.
    - declare all variables you need
    - body of the algorithm
    - returning result
*/

// Think about bad input (at least ask about this)

const memo = {
  1: 1,
  2: 2,
};
const climbingStairsRecursive = (n) => {
  if (memo[n]) return memo[n];

  memo[n - 1] = climbingStairsRecursive(n - 1);
  memo[n - 2] = climbingStairsRecursive(n - 2);

  return memo[n - 1] + memo[n - 2];
};

const climbingStairsIterative = (n) => {
  if (n < 2) return n;
  const cache = [1, 2];
  let c = 2;

  while (c < n) {
    cache[c] = cache[c - 1] + cache[c - 2];
    c++;
  }

  return cache.at(-1);
};

module.exports = climbingStairsRecursive;
