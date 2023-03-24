const validAnagram = require("./valid-anagram");
describe("", () => {
  it("should return ", () => {
    expect(validAnagram("anagram", "nagaram")).toBe(true);
  });

  it("should return ", () => {
    expect(validAnagram("rat", "cat")).toBe(false);
  });

  it("should return ", () => {
    expect(validAnagram("ab", "a")).toBe(false);
  });
});
