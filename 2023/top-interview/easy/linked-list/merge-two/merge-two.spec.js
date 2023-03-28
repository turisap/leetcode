const mergeTwo = require("./merge-two");
describe("", () => {
  it("should return ", () => {
    const l1 = {
      val: 1,
      next: {
        val: 2,
        next: {
          val: 3,
        },
      },
    };

    const l2 = {
      val: 1,
      next: {
        val: 3,
        next: {
          val: 4,
        },
      },
    };

    expect(mergeTwo(l1, l2)).toEqual({
      val: 1,
      next: {
        val: 1,
        next: {
          val: 2,
          next: {
            val: 3,
            next: {
              val: 3,
              next: {
                val: 4,
              },
            },
          },
        },
      },
    });
  });
});
