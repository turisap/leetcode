const prime_factors = require("../interview/prime_factors");

describe("should compute the sequence of prime factors of a given number", () => {
  test("should return 2 2 5 5 for 100", () => {
    expect(prime_factors(100)).toEqual("2 2 5 5");
  });

  test("should return 11 13 for 143 ", () => {
    expect(prime_factors(143)).toEqual("11 13");
  });

  test("should return 11 71 781", () => {
    expect(prime_factors(781)).toEqual("11 71");
  });

  test("should return 2 2 2 5 5 5 for 1000", () => {
    expect(prime_factors(1000)).toEqual("2 2 2 5 5 5");
  });
});
