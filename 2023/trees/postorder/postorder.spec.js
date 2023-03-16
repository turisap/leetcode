const postorderTraversal = require("./postorder");
describe("", () => {
  it("should return ", () => {
    const root = {
      val: 1,
      right: {
        val: 2,
        left: {
          val: 3,
        },
      },
    };

    const res = postorderTraversal(root);

    expect(res).toEqual([3, 2, 1]);
  });

  it("should return ", () => {
    const root = {
      val: 3,
      left: {
        val: 1,
      },
      right: {
        val: 2,
      },
    };

    const res = postorderTraversal(root);

    expect(res).toEqual([1, 2, 3]);
  });
});
