const singleEl = require("./single-element");
describe("", () => {
  it("should return 1", () => {
    expect(singleEl([2, 2, 1])).toBe(1);
  });
  //
  // it("should return 4", () => {
  //   expect(singleEl([4, 1, 2, 1, 2])).toBe(4);
  // });
  //
  // it("should return 1", () => {
  //   expect(singleEl([1])).toBe(1);
  // });
});
