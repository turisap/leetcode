const longestCommonPrefix = (strs) => {
  const sorted = strs.sort((a, b) => a.length - b.length);
  let lcp = sorted[0];

  for (let i = 1; i < sorted.length; i++) {
    const curr = sorted[i];
    const minLength = Math.min(lcp.length, curr.length);

    for (let c = 0; c < minLength; c++) {
      if (lcp[c] !== curr[c]) {
        lcp = lcp.slice(0, c);
        break;
      }
    }

    if (!lcp) {
      return "";
    }
  }

  return lcp;
};

// longestCommonPrefix(["dog", "racecar", "car"]);

module.exports = longestCommonPrefix;
