const closestValue = require("./closest-value");
describe("", () => {
  it("should return 4", () => {
    const root = {
      val: 4,
      left: {
        val: 2,
        left: {
          val: 1,
        },
        right: {
          val: 3,
        },
      },
      right: {
        val: 5,
      },
    };

    expect(closestValue(root, 3.714286)).toBe(4);
  });

  it("should return 2", () => {
    const root = {
      val: 1,
      right: { val: 2 },
    };

    expect(closestValue(root, 3.428571)).toBe(2);
  });
});
