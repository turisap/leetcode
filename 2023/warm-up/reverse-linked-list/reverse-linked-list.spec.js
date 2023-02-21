const reverseLinkedList = require("./reverse-linked-list");

const list = {
  val: 1,
  next: {
    val: 2,
    next: {
      val: 3,
    },
  },
};

const reversed = {
  val: 3,
  next: {
    val: 2,
    next: {
      val: 1,
    },
  },
};

describe("returns reversed list", () => {
  it("should return ", () => {
    expect(reverseLinkedList(list)).toEqual(reversed);
  });
});
