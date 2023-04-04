const isPowerThree = require("./is-power-three");
describe("", () => {
  it("should return true", () => {
    expect(isPowerThree(3)).toBe(true);
  });

  it("should return true ", () => {
    expect(isPowerThree(9)).toBe(true);
  });

  it("should return true ", () => {
    expect(isPowerThree(27)).toBe(true);
  });

  it("should return false ", () => {
    expect(isPowerThree(20)).toBe(false);
  });
});
