const myAtoi = require("./atoi");
describe("", () => {
  it("should return 42", () => {
    expect(myAtoi("42")).toBe(42);
  });

  it("returns negative -42", () => {
    expect(myAtoi("               -42")).toBe(-42);
  });

  it("trim leading whitespace", () => {
    expect(myAtoi("0052")).toBe(52);
  });

  it("returns negative", () => {
    expect(myAtoi("-0032")).toBe(-32);
  });

  it("ignores non-digits", () => {
    expect(myAtoi("512f34")).toBe(512);
  });

  it("clamps exceeding 32bits from the positive side", () => {
    expect(myAtoi(Math.pow(2, 32).toString())).toBe(Math.pow(2, 31) - 1);
  });

  it("clamps exceeding 32bits from the negative side", () => {
    expect(myAtoi("-91283472332")).toBe(-2147483648);
  });
});
