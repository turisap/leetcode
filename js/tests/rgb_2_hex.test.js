const rgb_2_hex = require("../beginner/rgb_2_hex.js");

describe("converts ligid rgb values to hex", () => {
  test("converst rgb(255, 255, 255) to FFFFFF", () => {
    expect(rgb_2_hex(255, 255, 255)).toBe("FFFFFF");
  });

  test("convers rgb(238, 238, 238) to EEEEEE", () => {
    expect(rgb_2_hex(238, 238, 238)).toBe("EEEEEE");
  });

  test("converts rgb(0, 0, 0) to 000000", () => {
    expect(rgb_2_hex(0, 0, 0)).toBe("000000");
  });

  test("converts rgb(0, 155, 216) to 009BD8", () => {
    expect(rgb_2_hex(0, 155, 216)).toBe("009BD8");
  });

  test("converts rgb(5, 20, 199) to 0514C7", () => {
    expect(rgb_2_hex(5, 20, 199)).toBe("0514C7");
  });
});

describe("converts illigal rgb values to hex", () => {
  test("converts rgb(-10, -10, -10) to 000000", () => {
    expect(rgb_2_hex(-10, -10, -10)).toBe("000000");
  });

  test("converts rgb(300, 500, 290) to FFFFFF", () => {
    expect(rgb_2_hex(300, 500, 290)).toBe("FFFFFF");
  });

  test("converts rgb(-100, 280, 7) to 00FF07", () => {
    expect(rgb_2_hex(-100, 280, 7)).toBe("00FF07");
  });
});
