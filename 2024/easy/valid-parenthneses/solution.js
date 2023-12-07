// CHECKLIST
// Write a few test examples. This ensures that I understood the task

/*
- Think out loud the algorithm and the way (sliding window, two pointers) you want to implement it.
    - declare all variables you need
    - body of the algorithm
    - returning result
*/

// Think about bad input (at least ask about this)

const closeMap = {
  "]": "[",
  "}": "{",
  ")": "(",
};

const openSet = new Set(["[", "{", "("]);

const isValid = (str) => {
  const stack = [];

  if (str.length % 2 !== 0) return false;

  for (const p of str) {
    if (openSet.has(p)) {
      stack.push(p);
      continue;
    }

    const curr = stack.pop();

    if (curr !== closeMap[p]) {
      return false;
    }
  }

  if (stack.length) return false;

  return true;
};

module.exports = isValid;
