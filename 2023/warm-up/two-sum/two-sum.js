const twoSum = (nums, target) => {
  const diffMap = {};

  for (let i = 0; i < nums.length; i++) {
    const current = nums[i];
    const diff = target - current;

    if (diffMap[diff] !== undefined) {
      return [diffMap[diff], i];
    }

    diffMap[current] = i;
  }
};

module.exports = twoSum;
