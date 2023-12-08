const solution = require("./solution.js");

describe("", () => {
  it("should return 2", () => {
    const res = solution([1, 2, 5, 6], 5);

    expect(res).toBe(2);
  });

  it("should return 0", () => {
    const res = solution([], 5);

    expect(res).toBe(0);
  });

  it("should return 4", () => {
    const res = solution([1, 2, 3, 4, 7, 10], 9);

    expect(res).toBe(5);
  });

  it("should return ?", () => {
    const res = solution([1, 3, 5, 6], 5);

    expect(res).toBe(2);
  });
});
