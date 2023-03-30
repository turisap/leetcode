const maxSub = require("./max-subarray");
describe("", () => {
  it("should return 6", () => {
    expect(maxSub([-2, 1, -3, 4, -1, 2, 1, -5, 4])).toBe(6);
  });

  it("should return 1", () => {
    expect(maxSub([-2, 1])).toBe(1);
  });

  it("should return -1", () => {
    expect(maxSub([-1, -2])).toBe(-1);
  });
});
