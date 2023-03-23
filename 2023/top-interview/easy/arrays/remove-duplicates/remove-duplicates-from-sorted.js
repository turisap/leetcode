function removeDuplicatesFromSorted(array) {
  if (!array.length) return 0;

  let p = 0;

  for (let i = 1; i < array.length; i++) {
    if (array[i] !== array[p]) {
      p++;
      array[p] = array[i];
    }
  }

  return p + 1;
}

module.exports = removeDuplicatesFromSorted;
