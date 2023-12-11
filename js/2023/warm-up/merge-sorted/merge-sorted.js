mergeTwoLists = function (arr1, arr2) {
  const res = [];
  let p1 = 0;
  let p2 = 0;

  while (p1 < arr1.length && p2 < arr2.length) {
    const curr1 = arr1[p1];
    const curr2 = arr2[p2];

    if (curr1 > curr2) {
      res.push(curr2);

      p2++;
    } else if (curr1 < curr2) {
      res.push(curr1);

      p1++;
    } else if (curr1 === curr2) {
      res.push(curr1);
      res.push(curr2);

      p1++;
      p2++;
    }
  }

  res.push(...arr2.slice(p2), ...arr1.slice(p1));

  return res;
};

// console.log(mergeTwoLists([1, 2, 4], [1, 3, 4]));
// console.log(mergeTwoLists([1, 2, 4], [1, 3, 4, 5, 6]));
// console.log(mergeTwoLists([], [0]));
console.log(mergeTwoLists([1, 2, 4], [1, 3, 4]));
