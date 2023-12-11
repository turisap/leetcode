const containsDuplicates = (arr) => {
  const sorted = arr.sort();

  for (let i = 0; i < sorted.length; i++) {
    if (sorted[i] === sorted[i - 1]) {
      return true;
    }
  }

  return false;
};

module.exports = containsDuplicates;
