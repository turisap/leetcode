const deleteEl = require("./delete-element");
describe("", () => {
  it("should return ", () => {
    const list = {
      val: 1,
      next: {
        val: 2,
        next: {
          val: 3,
          next: {
            val: 4,
            next: {
              val: 5,
            },
          },
        },
      },
    };

    expect(deleteEl(list, 2)).toEqual({
      val: 1,
      next: {
        val: 2,
        next: {
          val: 3,
          next: {
            val: 5,
          },
        },
      },
    });
  });

  it("should return ", () => {
    const list = {
      val: 1,
      next: {
        val: 2,
      },
    };

    expect(deleteEl(list, 1)).toEqual({
      val: 1,
    });
  });

  it("should return ", () => {
    const list = {
      val: 1,
      next: {
        val: 2,
      },
    };

    expect(deleteEl(list, 2)).toEqual({
      val: 2,
    });
  });
});
