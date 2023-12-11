const longestCommonPrefix = require("./longest-common-prefix");
describe("it calculates lcp", () => {
  it('returns \'\' for["dog", "racecar", "car"] ', () => {
    const res = longestCommonPrefix(["dog", "racecar", "car"]);
    expect(res).toBe("");
  });

  it('returns "fl" for "flower","flow","flight"', () => {
    const res = longestCommonPrefix(["flower", "flow", "flight"]);
    expect(res).toBe("fl");
  });
});
