const isNumberPolyndrome = require("./palindrome-number");
describe("validates palindrome number", () => {
  it("returns true for 121", () => {
    const res = isNumberPolyndrome(121);

    expect(res).toBe(true);
  });

  it("returns true for 12321", () => {
    const res = isNumberPolyndrome(12321);

    expect(res).toBe(true);
  });

  it("returns true for 5665", () => {
    const res = isNumberPolyndrome(5665);

    expect(res).toBe(true);
  });

  it("returns true for 0", () => {
    const res = isNumberPolyndrome(0);

    expect(res).toBe(true);
  });

  it("returns false -4554", () => {
    const res = isNumberPolyndrome(-4554);

    expect(res).toBe(false);
  });

  it("returns false 100", () => {
    const res = isNumberPolyndrome(100);

    expect(res).toBe(false);
  });
});
