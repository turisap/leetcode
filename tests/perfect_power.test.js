const {
  can_be_divided,
  perfect_power,
} = require("../beginner/perfect_power.js");

describe("should establish whether a number can be divided fully by a divisor", () => {
  test("should return true for 9 and 3", () => {
    expect(can_be_divided(9, 3)).toEqual({ success: true, pow: 2 });
  });

  test("should return true, 2 for 25 and 5", () => {
    expect(can_be_divided(25, 5)).toEqual({ success: true, pow: 2 });
  });

  test("should return true, 3 for 27 and 3", () => {
    expect(can_be_divided(27, 3)).toEqual({ success: true, pow: 3 });
  });

  test("should return true, 2 for 4 and 2", () => {
    expect(can_be_divided(4, 2)).toEqual({ success: true, pow: 2 });
  });

  test("should return true for 60466176 and 6", () => {
    expect(can_be_divided(60466176, 6)).toEqual({ success: true, pow: 10 });
  });

  test("should return true for 16 and 4", () => {
    expect(can_be_divided(16, 4)).toEqual({ success: true, pow: 2 });
  });

  test("should return false for 11 and 4", () => {
    expect(can_be_divided(11, 4)).toEqual({ success: false, pow: 0 });
  });

  test("should return false for 12 and 6", () => {
    expect(can_be_divided(12, 6)).toEqual({ success: false, pow: 0 });
  });

  test("should return false for 16 and 5", () => {
    expect(can_be_divided(16, 5)).toEqual({ success: false, pow: 0 });
  });
});

describe("should be calculate the power and number pair for perfect numbers", () => {
  test("should return [2, 2] for 4", () => {
    expect(perfect_power(4)).toEqual([2, 2]);
  });

  test("should return [3, 2] for 9 ", () => {
    expect(perfect_power(9)).toEqual([3, 2]);
  });

  test("should return [5, 2] for 25", () => {
    expect(perfect_power(25)).toEqual([5, 2]);
  });

  test("should return [30, 30] for 900", () => {
    expect(perfect_power(900)).toEqual([30, 2]);
  });
});
