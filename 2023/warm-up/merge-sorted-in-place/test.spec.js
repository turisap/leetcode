const mergeSortedInPlace = require("./merge-sorted");
describe("it should sort array in  place", () => {
  it("returns [ 1, 2, 2, 3, 5, 6 ]", () => {
    const sorted = mergeSortedInPlace([1, 2, 3, 0, 0, 0], 3, [2, 5, 6], 3);
    expect(sorted).toEqual([1, 2, 2, 3, 5, 6]);
  });

  it("returns [1]", () => {
    const sorted = mergeSortedInPlace([0], 0, [1], 1);
    expect(sorted).toEqual([1]);
  });

  it("returns [1]", () => {
    const sorted = mergeSortedInPlace([1], 1, [], 0);
    expect(sorted).toEqual([1]);
  });

  it("returns [1, 2]", () => {
    const sorted = mergeSortedInPlace([2, 0], 1, [1], 1);
    expect(sorted).toEqual([1, 2]);
  });

  it("returns [1, 2]", () => {
    const sorted = mergeSortedInPlace(
      [-1, 0, 0, 3, 3, 3, 0, 0, 0],
      6,
      [1, 2, 2],
      3
    );
    expect(sorted).toEqual([-1, 0, 0, 1, 2, 2, 3, 3, 3]);
  });
});
