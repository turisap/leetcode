const linkedListIntersection =
  require("./linked-list-intersection").linkedListIntersection;
const linkedListIntersectionTwoPointers =
  require("./linked-list-intersection").linkedListIntersectionTwoPointers;

describe("should return  list intersection", () => {
  it("should return 3", () => {
    const l1 = {
      value: 3,
    };

    const l2 = {
      val: 2,
      next: l1,
    };

    // const res = linkedListIntersection(l1, l2);
    const res2 = linkedListIntersectionTwoPointers(l1, l2);
    expect(res2).toBe(l1);
  });
});
