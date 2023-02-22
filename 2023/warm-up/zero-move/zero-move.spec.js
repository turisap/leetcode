const moveZerosSort = require("./zero-move").zeroMoveSort;
const zeroTwoPointers = require("./zero-move").zeroMoveTwoPointers;
describe("", () => {
  it("should return [1, 3, 12, 0, 0]", () => {
    expect(zeroTwoPointers([0, 1, 0, 3, 12])).toEqual([1, 3, 12, 0, 0]);
  });

  it("should return [0]", () => {
    expect(zeroTwoPointers([0])).toEqual([0]);
  });

  it("should return [1,0,0]", () => {
    expect(zeroTwoPointers([0, 0, 1]).toString()).toEqual([1, 0, 0].toString());
  });

  it("should return [1,3,12,0,0]", () => {
    const res1 = moveZerosSort([0, 1, 0, 3, 12]);
    const res2 = zeroTwoPointers([0, 1, 0, 3, 12]);

    expect(res1).toEqual(res2);
  });

  it("should return [1,3,12,0,0]", () => {
    const res1 = moveZerosSort([0, 1, 0, 3, 12]);
    const res2 = zeroTwoPointers([0, 1, 0, 3, 12]);

    expect(res1).toEqual(res2);
  });
});
