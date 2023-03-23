const removeDuplicatesFromSorted = require("./remove-duplicates-from-sorted");

describe("it should return the number of unique elements in an array", () => {
  test("should return empty array for an empty array input", () => {
    expect(removeDuplicatesFromSorted([])).toEqual(0);
  });

  test("should filter out [1,1,2]", () => {
    expect(removeDuplicatesFromSorted([1, 1, 2])).toEqual(2);
  });

  test("should filter out [0,0,1,1,1,2,2,3,3,4]", () => {
    expect(removeDuplicatesFromSorted([0, 0, 1, 1, 1, 2, 2, 3, 3, 4])).toEqual(
      5
    );
  });

  test("should filter out [0,0,1,1,1]", () => {
    expect(removeDuplicatesFromSorted([0, 0, 1, 1, 1])).toEqual(2);
  });

  test("should return the same array if no duplicates found", () => {
    expect(removeDuplicatesFromSorted([0, 1, 2, 3])).toEqual(4);
  });
});
