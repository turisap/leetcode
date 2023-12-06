const solution = require("./solution.js");

describe("longest common prefix", () => {
  it("should return fl", () => {
    const res = solution(["flower", "flow", "fl"]);

    expect(res).toBe("fl");
  });

  it("should return fl", () => {
    const res = solution(["flower", "flow", "flight"]);

    expect(res).toBe("fl");
  });

  it("should return cat", () => {
    const res = solution(["cat"]);

    expect(res).toBe("cat");
  });

  it("should return cat", () => {
    const res = solution(["cat", "job"]);

    expect(res).toBe("");
  });
});
