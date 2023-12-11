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

//[2,3,3,4,5,null,4]
const failedTree = {
  val: 2,
  left: {
    val: 3,
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
      val: null,
    },
    right: {
      val: 4,
    },
  },
};
describe("verifies a symmetric tree", () => {
  it("verifies a symmetric tree", () => {
    expect(isSymmetric(symTree)).toBe(true);
  });

  it("returns verifies an asymmetric tree", () => {
    expect(isSymmetric(assymTree)).toBe(false);
  });

  it("verifies a shallow tree", () => {
    expect(isSymmetric(shallowTree)).toBe(false);
  });

  it("verifies an edge case", () => {
    expect(isSymmetric(failedTree)).toBe(false);
  });
});
