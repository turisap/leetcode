const firstUnique = (str) => {
  const hashMap = {};

  for (let s of str) {
    hashMap[s] = hashMap[s] ? (hashMap[s] += 1) : 1;
  }

  for (let i = 0; i < str.length; i++) {
    if (hashMap[str[i]] === 1) {
      return i;
    }
  }

  return -1;
};

module.exports = firstUnique;
