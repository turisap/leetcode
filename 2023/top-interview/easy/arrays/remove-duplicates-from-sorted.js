function removeDuplicatesFromSorted(array) {
  let count = 1;

  if (!array.length) {
    return 0;
  }

  for (let i = 1; i < array.length; i++) {
    const current = array[i];
    const prev = array[i - 1];

    if (current !== prev) {
      count++;
    }
  }

  return count;
}

module.exports = removeDuplicatesFromSorted;
