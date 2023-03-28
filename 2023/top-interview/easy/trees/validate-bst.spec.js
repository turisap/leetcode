const validateBst = require("./validate-bst");
describe("", () => {
  it("should return ", () => {
    const root = {
      val: 2,
      left: {
        val: 1,
      },
      right: {
        val: 3,
      },
    };

    expect(validateBst(root)).toBe(true);
  });

  it("should return ", () => {
    const root = {
      val: 5,
      left: {
        val: 1,
      },
      right: {
        val: 4,
        left: {
          val: 3,
        },
        right: {
          val: 6,
        },
      },
    };

    expect(validateBst(root)).toBe(false);
  });

  it("should return ", () => {
    const root = {
      val: 2,
      right: {
        val: 2,
      },
      left: {
        val: 2,
      },
    };

    expect(validateBst(root)).toBe(false);
  });
});
