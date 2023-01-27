var removeElement = function (nums, val) {
  let p = 0;

  while (p < nums.length) {
    if (nums[p] === val) {
      nums.splice(p, 1);
    } else {
      p++;
    }
  }
};

removeElement([3, 2, 2, 3], 3);
removeElement([0, 1, 2, 2, 3, 0, 4, 2], 2);
