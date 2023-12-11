const minDiff = require("./min-diff");
describe("", () => {
  it("should return ", () => {
    const root = {
      val: 5,
      left: {
        val: 1,
      },
      right: {
        val: 7,
      },
    };

    expect(minDiff(root)).toBe(2);
  });

  it("should return ", () => {
    const root = {
      val: 10,
      left: {
        val: 7,
      },
      right: {
        val: 15,
      },
    };

    expect(minDiff(root)).toBe(3);
  });

  it("should return 6", () => {
    const root = {
      val: 27,
      right: {
        val: 34,
        right: {
          val: 58,
          left: {
            val: 50,
            left: {
              val: 44,
            },
          },
        },
      },
    };

    expect(minDiff(root)).toBe(6);
  });
});
