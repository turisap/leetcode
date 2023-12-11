const { solution, isBadVersion } = require("./bad-version");
describe("", () => {
  it("should return 4", () => {
    const cb = isBadVersion(4);
    const fn = solution(cb);

    expect(fn(10)).toBe(4);
  });

  it("should return 12", () => {
    const cb = isBadVersion(12);
    const fn = solution(cb);

    expect(fn(100)).toBe(12);
  });

  it("should return 1", () => {
    const cb = isBadVersion(1);
    const fn = solution(cb);

    expect(fn(1)).toBe(1);
  });

  it("should return 2", () => {
    const cb = isBadVersion(2);
    const fn = solution(cb);

    expect(fn(2)).toBe(2);
  });
});
