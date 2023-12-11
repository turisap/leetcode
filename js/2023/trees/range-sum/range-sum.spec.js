const rangeSum = require("./range-sum");
describe("", () => {
  it("should return ", () => {
    const root = {
      val: 10,
      left: {
        val: 5,
        left: {
          val: 2,
        },
        right: {
          val: 7,
        },
      },
      right: {
        val: 15,
        left: {
          val: 18,
        },
      },
    };

    expect(rangeSum(root, 7, 15)).toBe(32);
  });
});
