const removeDuplicates = (nums) => {
  let currentUnique = nums[0];
  let p = 1;

  while (p < nums.length) {
    const curr = nums[p];

    if (currentUnique === curr) {
      nums.splice(p, 1);
    } else {
      currentUnique = nums[p];
      p++;
    }
  }

  return nums;
};

console.log(removeDuplicates([1, 1, 2, 2, 2, 3, 3, 3, 4]));
removeDuplicates([1, 1, 2]);
removeDuplicates([0]);
removeDuplicates([]);
