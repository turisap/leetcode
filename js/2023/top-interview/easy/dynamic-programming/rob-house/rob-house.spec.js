const robHouse = require("./rob-house");
describe("", () => {
  it("should return 4", () => {
    expect(robHouse([1, 2, 3, 1])).toBe(4);
  });

  it("should return 12", () => {
    expect(robHouse([2, 7, 9, 3, 1])).toBe(12);
  });

  it("should return 3", () => {
    expect(robHouse([2, 1, 1, 2])).toBe(4);
  });
});

// 2, 7, 9, 3, 1
// 0, 2,
