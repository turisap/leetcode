const leftSum = require("./left-sum");
describe("", () => {
  it("should return 24", () => {
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
          val: 7,
        },
      },
    };

    expect(leftSum(root)).toBe(24);
  });
});
