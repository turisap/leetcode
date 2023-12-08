const solution = require("./solution.js");

describe("", () => {
  it("should return [1,2,3,4,5,6]", () => {
    const res = solution([1, 3, 5], [2, 4, 6]);

    expect(res).toEqual([1, 2, 3, 4, 5, 6]);
  });

  // write the following tests:
  it("should return []", () => {
    const res = solution([], []);

    expect(res).toEqual([]);
  });

  //write test for the following cases:
  it("should return [1,2,3]", () => {
    const res = solution([1, 2, 3], []);

    expect(res).toEqual([1, 2, 3]);
  });

  it("should return [1,2,3]", () => {
    const res = solution([], [1, 2, 3]);

    expect(res).toEqual([1, 2, 3]);
  });

  it("should return [1,2,3,4,5]", () => {
    const res = solution([1, 2, 3], [4, 5]);

    expect(res).toEqual([1, 2, 3, 4, 5]);
  });

  it("should return [1,2,4,5,7,9,15]", () => {
    const res = solution([1, 5, 9], [2, 4, 7, 15]);

    expect(res).toEqual([1, 2, 4, 5, 7, 9, 15]);
  });

  it("should return [1,2,3,4,4]", () => {
    const res = solution([1, 2, 4], [1, 3, 4]);

    expect(res).toEqual([1, 1, 2, 3, 4, 4]);
  });

  it("should return [1,2,3,4,4]", () => {
    const res = solution([1, 3, 4], [1, 2, 4]);

    expect(res).toEqual([1, 1, 2, 3, 4, 4]);
  });
});
