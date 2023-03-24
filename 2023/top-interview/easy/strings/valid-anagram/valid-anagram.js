const validAnagram = (s1, s2) => {
  if (s1.length !== s2.length) return false;

  const map = new Map();

  for (let str of s1) {
    if (map.has(str)) map.set(str, map.get(str) + 1);
    else map.set(str, 1);
  }

  for (let str of s2) {
    if (!map.get(str)) {
      return false;
    }

    map.set(str, map.get(str) - 1);
  }

  return true;
};

module.exports = validAnagram;
