const twoSum = require("./two-sum");
describe("Two sum", () => {
  it("should return [1,3]", () => {
    const input = [1, 2, 5, 7];

    expect(twoSum(input, 9)).toEqual([1, 3]);
  });

  it("should return [1,3] for 9", () => {
    const input = [0, 2, 5, 7];

    expect(twoSum(input, 9)).toEqual([1, 3]);
  });

  it("should return [0,4]", () => {
    const input = [1, 2, 3, 4, 5];

    expect(twoSum(input, 6)).toEqual([1, 3]);
  });

  it("should return [1,2]", () => {
    const input = [3, 4, 5];

    expect(twoSum(input, 90)).toEqual(null);
  });

  it("should return [0,1]", () => {
    const input = [3, 3];

    expect(twoSum(input, 6)).toEqual([0, 1]);
  });
});
