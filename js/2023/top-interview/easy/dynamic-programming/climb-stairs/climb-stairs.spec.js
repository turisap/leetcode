const climbStairs = require("./climb-stairs");
describe("", () => {
  it("should return 1", () => {
    expect(climbStairs(1)).toBe(1);
  });

  it("should return 2", () => {
    expect(climbStairs(2)).toBe(2);
  });

  it("should return 3", () => {
    expect(climbStairs(3)).toBe(3);
  });

  it("should return 5", () => {
    expect(climbStairs(4)).toBe(5);
  });

  it("should return 8", () => {
    expect(climbStairs(5)).toBe(8);
  });

  it("should return 1836311903", () => {
    expect(climbStairs(45)).toBe(1836311903);
  });
});
