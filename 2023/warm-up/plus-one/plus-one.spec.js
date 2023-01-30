const plusOne = require("./plus-one");
describe("it should add one properly", () => {
  it("adds to [1,2,3]", () => {
    const res = plusOne([1, 2, 3]);

    expect(res).toEqual([1, 2, 4]);
  });

  it("adds to [1,2,9]", () => {
    const res = plusOne([1, 2, 9]);

    expect(res).toEqual([1, 3, 0]);
  });

  it("adds to [1,8,9,9]", () => {
    const res = plusOne([1, 8, 9, 9]);

    expect(res).toEqual([1, 9, 0, 0]);
  });

  it("adds to [0]", () => {
    const res = plusOne([0]);

    expect(res).toEqual([1]);
  });

  it("adds to [9,9,9]", () => {
    const res = plusOne([9, 9, 9]);

    expect(res).toEqual([1, 0, 0, 0]);
  });
});
