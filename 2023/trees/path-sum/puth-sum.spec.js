const pathSum = require("./path-sum");
describe("", () => {
  it("should return", () => {
    const root = {
      val: 1,
    };

    expect(pathSum(root, 1)).toBe(true);
  });

  it("should return", () => {
    const root = {
      val: 1,
      left: {
        val: 2,
      },
    };

    expect(pathSum(root, 1)).toBe(false);
  });

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

    expect(pathSum(root, 5)).toBe(false);
  });

  it("should return true for a small tree", () => {
    const root = {
      val: 1,
      left: {
        val: 2,
      },
      right: {
        val: 3,
      },
    };

    expect(pathSum(root, 4)).toBe(true);
  });

  it("should return ", () => {
    const root = {
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
        right: {
          val: 7,
          right: {
            val: 9,
          },
        },
      },
    };

    expect(pathSum(root, 20)).toBe(true);
  });
});
