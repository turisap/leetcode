const listPalindrome = require("./list-palindrome");
describe("", () => {
  it("should return ", () => {
    const head = {
      val: 1,
      next: {
        val: 2,
        next: {
          val: 2,
          next: {
            val: 1,
          },
        },
      },
    };
    expect(listPalindrome(head)).toBe(true);
  });

  it("should return ", () => {
    const head = {
      val: 1,
      next: {
        val: 2,
      },
    };
    expect(listPalindrome(head)).toBe(false);
  });
});
