const inverseTree = require("./inverse-tree");
describe("should inverse tree", () => {
  // it("should return ", () => {
  //   const root = {
  //     val: 2,
  //     left: {
  //       val: 1,
  //     },
  //     right: {
  //       val: 3,
  //     },
  //   };
  //
  //   const inverseRoot = {
  //     val: 2,
  //     left: {
  //       val: 3,
  //     },
  //     right: {
  //       val: 1,
  //     },
  //   };
  //
  //   expect(inverseTree(root)).toEqual(inverseRoot);
  // });

  it("should return ", () => {
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
        val: 7,
        left: {
          val: 6,
        },
        right: {
          val: 9,
        },
      },
    };

    const inverseRoot = {
      val: 4,
      left: {
        val: 7,
        left: {
          val: 3,
        },
        right: {
          val: 1,
        },
      },
      right: {
        val: 2,
        left: {
          val: 9,
        },
        right: {
          val: 6,
        },
      },
    };

    expect(inverseTree(root)).toEqual(inverseRoot);
  });
});
