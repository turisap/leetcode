const decimal2Romanian = require("../beginner/romanian_numbers.js");

// simple tests
test("converts 2 to II", () => {
  expect(decimal2Romanian(2)).toBe("II");
});

test("converts 5 to V", () => {
  expect(decimal2Romanian(5)).toBe("V");
});

test("converts 7 to VII", () => {
  expect(decimal2Romanian(7)).toBe("VII");
});

test("converts 13 to XIII", () => {
  expect(decimal2Romanian(13)).toBe("XIII");
});

test("converts 50 to L", () => {
  expect(decimal2Romanian(50)).toBe("L");
});

test("converts 123 to CXXIII", () => {
  expect(decimal2Romanian(123)).toBe("CXXIII");
});

test("converts 500 to D", () => {
  expect(decimal2Romanian(500)).toBe("D");
});

test("converts 1000 to M", () => {
  expect(decimal2Romanian(1000)).toBe("M");
});

// tests with 4 same letters in a row (substraction case)
test("converts 4 to IV", () => {
  expect(decimal2Romanian(4)).toBe("IV");
});

test("converts 9 to IX", () => {
  expect(decimal2Romanian(9)).toBe("IX");
});

test("converts 19  to XIX", () => {
  expect(decimal2Romanian(19)).toBe("XIX");
});

test("converts 40  to XL", () => {
  expect(decimal2Romanian(40)).toBe("XL");
});

test("converts 50  to XIX", () => {
  expect(decimal2Romanian(50)).toBe("L");
});

test("converts 90  to XC", () => {
  expect(decimal2Romanian(90)).toBe("XC");
});

test("converts 100  to C", () => {
  expect(decimal2Romanian(100)).toBe("C");
});

test("converts 1990 t0 MCMXC", () => {
  expect(decimal2Romanian(1990)).toBe("MCMXC");
});

test("converts 2008 to MMVIII", () => {
  expect(decimal2Romanian(2008)).toBe("MMVIII");
});
