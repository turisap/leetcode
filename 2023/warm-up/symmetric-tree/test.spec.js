const isSymmetric = require("./symmetric-tree");

const symTree = {
  val: 1,
  left: {
    val: 2,
    left: {
      val: 3,
    },
    right: {
      val: 4,
    },
  },
  right: {
    val: 2,
    left: {
      val: 4,
    },
    right: {
      val: 3,
    },
  },
};

const assymTree = {
  val: 1,
  left: {
    val: 2,
    right: {
      val: 3,
    },
  },
  right: {
    val: 2,
    right: {
      val: 3,
    },
  },
};

const shallowTree = {
  val: 1,
  left: {
    val: 2,
  },
  right: {
    val: 3,
  },
};
describe("verifies a symmetric tree", () => {
  it("returns true", () => {
    expect(isSymmetric(symTree)).toBe(true);
  });

  it("returns false", () => {
    expect(isSymmetric(assymTree)).toBe(false);
  });

  it("returns false", () => {
    expect(isSymmetric(shallowTree)).toBe(false);
  });
});
