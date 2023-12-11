const levelOrder = require("./level-order");
describe("", () => {
  it("should return ", () => {
    const root = {
      val: 5,
      left: {
        val: 3,
      },
      right: {
        val: 6,
      },
    };

    expect(levelOrder(root)).toEqual([[5], [3, 6]]);
  });
});
