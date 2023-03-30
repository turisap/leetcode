var Solution = function (nums) {
  this.original = nums;
};

/**
 * @return {number[]}
 */
Solution.prototype.reset = function () {
  return this.original;
};

/**
 * @return {number[]}
 */
Solution.prototype.shuffle = function () {
  this.nums = [...this.original];

  for (let i = 0; i < this.nums.length; i++) {
    const random = Math.floor(Math.random() * this.nums.length);
    const temp = this.nums[random];
    this.nums[random] = this.nums[i];
    this.nums[i] = temp;
  }

  return this.nums;
};

module.exports = Solution;
