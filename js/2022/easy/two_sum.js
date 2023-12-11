const two_sum = (nums, target) => {
  const diffMap = {};

  for (let i = 0; i < nums.length; i++) {
    const current = nums[i];
    const secondPart = target - current;

    if (diffMap[current] !== undefined) {
      return [i, diffMap[current]];
    }

    diffMap[secondPart] = i;
  }
};

console.log(two_sum([2, 7, 11, 15], 9));
console.log(two_sum([3, 3], 6));
