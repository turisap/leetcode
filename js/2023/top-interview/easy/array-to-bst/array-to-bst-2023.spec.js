const arrayToBst = require("./array-to-bst");
describe("", () => {
  it("should return ", () => {
    // [-10, -3, 0, 5, 9];

    expect(arrayToBst([-10, -3, 0, 5, 9])).toEqual({
      val: 0,
      left: {
        val: -10,
        left: null,
        right: {
          val: -3,
        },
      },
      right: {
        val: 5,
        left: null,
        right: {
          val: 9,
        },
      },
    });
  });
});
