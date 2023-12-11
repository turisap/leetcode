const minDepth = require("./min-depth");
describe("", () => {
  it("should return ", () => {
    const root = {
      val: 1,
      left: {
        val: 2,
      },
      right: {
        val: 3,
        right: {
          val: 4,
        },
      },
    };

    expect(minDepth(root)).toBe(2);
  });

  it("should return 2", () => {
    const root = {
      val: 1,
      left: {
        val: 2,
        right: {
          val: 7,
          right: {
            val: 7,
          },
        },
      },
      right: {
        val: 3,
        right: {
          val: 4,
          left: {
            val: 5,
          },
        },
      },
    };

    expect(minDepth(root)).toBe(4);
  });
});
