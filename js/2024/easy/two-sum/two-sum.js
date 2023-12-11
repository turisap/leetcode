// CHECKLIST
// Think about bad input (at least ask about this)
// Write a few test examples. This ensures that I understood the task
/*
- Think out loud the algorithm and the way (sliding window, two pointers) you want to implement it.
    - declare all variables you need
    - body of the algorithm
    - returning result
*/

const TwoSum = (input, target) => {
  const hashMap = {};

  for (let i = 0; i < input.length; i++) {
    const diff = target - input[i];

    if (typeof hashMap[diff] !== "undefined") {
      console.log(hashMap[diff], i);
      return [hashMap[diff], i];
    }

    hashMap[input[i]] = i;
  }

  return null;
};

/*
  {
      "3": 2,
      "4": 1,
      "5": 0
    }

 */
module.exports = TwoSum;
