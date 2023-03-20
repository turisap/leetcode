const binaryTreePath = require("./tree-path");
describe("", () => {
  it("should return ", () => {
    const root = {
      val: 1,
      left: {
        val: 2,
      },
      right: {
        val: 3,
      },
    };

    expect(binaryTreePath(root)).toEqual(["1->2", "1->3"]);
  });

  it("should return ", () => {
    const root = {
      val: 1,
      left: {
        val: 2,
        left: {
          val: 5,
        },
        right: {
          val: 6,
        },
      },
      right: {
        val: 3,
      },
    };

    expect(binaryTreePath(root)).toEqual(["1->2->5", "1->2->6", "1->3"]);
  });
});
