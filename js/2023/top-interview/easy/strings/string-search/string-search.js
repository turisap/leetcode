const stringSearch = (haystack, needle) => {
  for (let i = 0; i < haystack.length; i++) {
    if (haystack[i] !== needle[0]) {
      continue;
    }

    const sub = haystack.substring(i, i + needle.length);

    if (sub === needle) return i;
  }

  return -1;
};

module.exports = stringSearch;
