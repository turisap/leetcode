// CHECKLIST
// Write a few test examples. This ensures that I understood the task

/*
- Think out loud the algorithm and the way (sliding window, two pointers) you want to implement it.
    - declare all variables you need
    - body of the algorithm
    - returning result
*/

// Think about bad input (at least ask about this)

const solution = (words) => {
  words.sort();

  let res = words[0];
  let c = 1;

  while (c < words.length) {
    const curr = words[c];

    let i = 0;
    while (i < res.length) {
      if (res[i] !== curr[i]) {
        res = res.substring(0, i);
        break;
      }

      i++;
    }

    c++;

    if (!res) return "";
  }

  return res;
};

module.exports = solution;
