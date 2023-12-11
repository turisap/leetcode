const twoSum = require("./two-sum");
describe("it should calculate the target number", () => {
  it("returns [0,1]", () => {
    const res = twoSum([2, 7, 11, 15], 9);

    expect(res).toEqual([0, 1]);
  });

  it("returns [1,2]", () => {
    const res = twoSum([3, 2, 4], 6);

    expect(res).toEqual([1, 2]);
  });

  it("returns [0,1]", () => {
    const res = twoSum([3, 3], 6);

    expect(res).toEqual([0, 1]);
  });

  it("returns [0,2]", () => {
    const res = twoSum([3, 2, 3], 6);

    expect(res).toEqual([0, 2]);
  });
});
