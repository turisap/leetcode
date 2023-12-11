const inorderTraversal = require("./inorder");
describe("", () => {
  it("should return ", () => {
    const testTree = {
      val: 1,
      left: {
        val: 2,
      },
      right: {
        val: 3,
      },
    };

    const res = inorderTraversal(testTree);

    expect(res).toEqual([2, 1, 3]);
  });

  it("should return ", () => {
    const testTree = {
      val: 1,
      left: {
        val: 2,
        left: {
          val: 4,
        },
        right: {
          val: 5,
        },
      },
      right: {
        val: 3,
        left: {
          val: 6,
        },
      },
    };

    const res = inorderTraversal(testTree);

    expect(res).toEqual([4, 2, 5, 1, 6, 3]);
  });

  it("should return ", () => {
    const testTree = {
      val: 2,
      left: {
        val: 3,
        left: {
          val: 1,
        },
      },
    };

    const res = inorderTraversal(testTree);

    expect(res).toEqual([1, 3, 2]);
  });
});
