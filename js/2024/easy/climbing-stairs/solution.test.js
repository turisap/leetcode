const solution = require("./solution.js");

describe("", () => {
  it("should return 1", () => {
    const res = solution(1);

    expect(res).toBe(1);
  });

  it("should return 2", () => {
    const res = solution(2);

    expect(res).toBe(2);
  });

  it("should return 3", () => {
    const res = solution(3);

    expect(res).toBe(3);
  });

  it("should return 4", () => {
    const res = solution(4);

    expect(res).toBe(5);
  });
});
