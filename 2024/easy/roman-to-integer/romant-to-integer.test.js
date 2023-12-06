const rTI = require("./roman-to-integer");
describe("roman to integer", () => {
  it("should return 3", () => {
    const res = rTI("III");

    expect(res).toBe(3);
  });

  it("should return 4", () => {
    const res = rTI("IV");

    expect(res).toBe(4);
  });

  it("should return 9", () => {
    const res = rTI("IX");

    expect(res).toBe(9);
  });

  it("should return 40", () => {
    const res = rTI("XL");

    expect(res).toBe(40);
  });

  it("should return 90", () => {
    const res = rTI("XC");

    expect(res).toBe(90);
  });

  it("should return 400", () => {
    const res = rTI("CD");

    expect(res).toBe(400);
  });

  it("should return 900", () => {
    const res = rTI("CM");

    expect(res).toBe(900);
  });
});
