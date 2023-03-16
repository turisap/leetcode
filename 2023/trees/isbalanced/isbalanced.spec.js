const isBalanced = require("./isbalanced");
describe("", () => {
  it("should return ", () => {
    const root = {
      val: 3,
      left: {
        val: 9,
      },
      right: {
        val: 20,
        left: {
          val: 15,
        },
        right: {
          val: 99,
        },
      },
    };

    expect(isBalanced(root)).toBe(true);
  });

  it("should return ", () => {
    const root = {
      val: 3,
      left: {
        val: 9,
      },
      right: {
        val: 20,
        left: {
          val: 15,
        },
        right: {
          val: 99,
          right: {
            val: 5,
          },
        },
      },
    };

    expect(isBalanced(root)).toBe(false);
  });
});
