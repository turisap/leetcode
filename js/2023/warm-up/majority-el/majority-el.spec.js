const majorityEl = require("./majority-el");
describe("find majority element", () => {
  it("should return 3", () => {
    expect(majorityEl([3, 2, 3])).toBe(3);
  });

  it("should return 2", () => {
    expect(majorityEl([2, 2, 1, 1, 1, 2, 2])).toBe(2);
  });

  it("should return 3", () => {
    expect(majorityEl([3, 3, 4])).toBe(3);
  });
});
