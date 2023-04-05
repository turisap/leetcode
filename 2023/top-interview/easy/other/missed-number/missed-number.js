const missedNumber = (nums) => {
  // let s = nums.length;
  //
  // for (let i = 0; i < nums.length; i++) {
  //   s ^= nums[i];
  //   s ^= i;
  // }
  //
  // return s;

  nums.sort((a, b) => a - b);

  for (let i = 0; i < nums.length; i++) {
    if (nums[i] !== i) return i;
  }

  return nums.length;
};

module.exports = missedNumber;
