const primeNumbersCount = require("./prime-numbers");
describe("", () => {
  it("should return 0 for 0", () => {
    expect(primeNumbersCount(0)).toBe(0);
  });

  it("should return 0 for 1", () => {
    expect(primeNumbersCount(1)).toBe(0);
  });

  it("should return 0 for 2", () => {
    expect(primeNumbersCount(2)).toBe(0);
  });

  it("should return 1 for 3", () => {
    expect(primeNumbersCount(3)).toBe(1);
  });

  it("should return 2", () => {
    expect(primeNumbersCount(4)).toBe(2);
  });

  it("should return 4 for 10", () => {
    expect(primeNumbersCount(10)).toBe(4);
  });

  it("should return 4 for 11", () => {
    expect(primeNumbersCount(11)).toBe(4);
  });

  it("should return  1229 for 10000 ", () => {
    expect(primeNumbersCount(10000)).toBe(1229);
  });
});
