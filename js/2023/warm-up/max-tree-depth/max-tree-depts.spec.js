const maxTreeDepth = require("./max-tree-depth");

const testTree = {
  val: 1,
  left: {
    val: 2,
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

const onlyRoot = {
  val: 1,
};

const depthTwo = {
  val: 1,
  left: {
    val: 2,
  },
};
describe("should calculate max tree depth", () => {
  it("should return 3", () => {
    expect(maxTreeDepth(testTree)).toBe(3);
  });

  it("should return 3", () => {
    expect(maxTreeDepth(onlyRoot)).toBe(1);
  });

  it("should return 2", () => {
    expect(maxTreeDepth(depthTwo)).toBe(2);
  });
});
