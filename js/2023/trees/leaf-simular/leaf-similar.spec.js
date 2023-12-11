const { getLeafSequenceRecur, leafSimilar } = require("./leaf-similar");
describe("", () => {
  it("should return true", () => {
    const root1 = {
      val: 1,
      left: {
        val: 2,
      },
      right: {
        val: 3,
      },
    };

    const root2 = {
      val: 1,
      left: {
        val: 3,
      },
      right: {
        val: 2,
      },
    };

    expect(leafSimilar(root1, root2)).toBe(false);
  });

  it("should return 23", () => {
    const root = {
      val: 1,
      left: {
        val: 2,
      },
      right: {
        val: 3,
      },
    };

    expect(getLeafSequenceRecur(root)).toBe("23");
  });

  it("should return 32", () => {
    const root = {
      val: 1,
      left: {
        val: 3,
      },
      right: {
        val: 2,
      },
    };

    expect(getLeafSequenceRecur(root)).toBe("32");
  });

  it("should return 23", () => {
    const root = {
      val: 1,
      left: {
        val: 2,
      },
      right: {
        val: 3,
      },
    };

    expect(getLeafSequenceRecur(root)).toBe("23");
  });

  it("should return 467", () => {
    const root = {
      val: 1,
      left: {
        val: 2,
        left: {
          val: 4,
        },
        right: {
          val: 5,
          left: {
            val: 6,
          },
        },
      },
      right: {
        val: 3,
        right: {
          val: 7,
        },
      },
    };

    expect(getLeafSequenceRecur(root)).toBe("467");
  });

  //3,5,1,6,2,9,8,null,null,7,14]
  it("should return ", () => {
    const root = {
      val: 3,
      right: {
        val: 5,
        right: {
          val: 6,
        },
        left: {
          val: 2,
          right: {
            val: 7,
          },
          left: {
            val: 14,
          },
        },
      },
      left: {
        val: 1,
        right: {
          val: 9,
        },
        left: {
          val: 8,
        },
      },
    };

    expect(getLeafSequenceRecur(root)).toBe("891476");
  });
});
