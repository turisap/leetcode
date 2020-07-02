const merge_two_sorted = (arr1, arr2) => {
  const res = [];

  while (arr1.length && arr2.length) {
    if (arr1[0] > arr2[0]) {
      res.push(arr2.shift());
      continue;
    }

    res.push(arr1.shift());
  }

  return res.concat(arr1).concat(arr2);
};

console.log(merge_two_sorted([1, 2, 4, 5, 7, 9, 15, 19, 21], [1, 3, 8]));
