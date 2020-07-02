const is_prime = require("../interview/prime_number.js");

describe("should validate prime number", () => {
  test("should return true for 3", () => {
    expect(is_prime(3)).toBe(true);
  });
  test("should return true for 5", () => {
    expect(is_prime(5)).toBe(true);
  });
  test("should return true for 7", () => {
    expect(is_prime(7)).toBe(true);
  });
  test("should return true for 11", () => {
    expect(is_prime(11)).toBe(true);
  });
  test("should return true for 2", () => {
    expect(is_prime(2)).toBe(true);
  });
  test("should return false for 6", () => {
    expect(is_prime(6)).toBe(false);
  });
  test("should return false for 99", () => {
    expect(is_prime(99)).toBe(false);
  });
});
