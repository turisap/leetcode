const lib = require("./romant-to-integer");

const { isDouble, romanToInteger } = lib;
describe("testes isDouble", () => {
  it("returns false for I and I", () => {
    expect(isDouble("I", "I")).toBe(false);
  });

  it("returns true for I and V", () => {
    expect(isDouble("I", "V")).toBe(true);
  });
});
