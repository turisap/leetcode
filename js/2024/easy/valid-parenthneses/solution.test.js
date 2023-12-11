const solution = require("./solution.js");

describe("valid parentheses", () => {
  it("should return true", () => {
    const res = solution("()");

    expect(res).toBe(true);
  });

  it("should return true", () => {
    const res = solution("[({})]");

    expect(res).toBe(true);
  });

  it("should return true", () => {
    const res = solution("[]()");

    expect(res).toBe(true);
  });

  it("should return false", () => {
    const res = solution("[()}");

    expect(res).toBe(false);
  });

  it("should return false", () => {
    const res = solution("[");

    expect(res).toBe(false);
  });

  it("should return false", () => {
    const res = solution("((");

    expect(res).toBe(false);
  });
});
