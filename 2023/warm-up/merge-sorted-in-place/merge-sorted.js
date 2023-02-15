const mergeSortedInPlace = (arr1, m, arr2, n) => {
  // make a copy of arr1 meat values
  // use three pointers to rewrite arr1
  if (!arr2.length) return arr1;

  const copy1 = arr1.slice(0, m);
  const copy2 = arr2;
  let p1 = 0;
  let p2 = 0;
  let p3 = 0;

  while (p3 < arr1.length) {
    let val1 = copy1[p1];
    let val2 = copy2[p2];

    switch (true) {
      case typeof val1 === "undefined": {
        arr1[p3] = val2;
        p2++;

        break;
      }
      case typeof val2 === "undefined": {
        arr1[p3] = val1;
        p1++;

        break;
      }
      case val1 > val2: {
        arr1[p3] = val2;
        p2++;

        break;
      }
      default: {
        arr1[p3] = val1;
        p1++;
      }
    }

    p3++;
  }

  return arr1;
};

module.exports = mergeSortedInPlace;
