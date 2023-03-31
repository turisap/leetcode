const fizzBazz = require("./fizz-bazz");
describe("", () => {
  it('should return ["1","2","Fizz"]', () => {
    expect(fizzBazz(3)).toEqual(["1", "2", "Fizz"]);
  });

  it("should return for 5", () => {
    expect(fizzBazz(5)).toEqual(["1", "2", "Fizz", "4", "Buzz"]);
  });

  it("should return for 15", () => {
    expect(fizzBazz(15)).toEqual([
      "1",
      "2",
      "Fizz",
      "4",
      "Buzz",
      "Fizz",
      "7",
      "8",
      "Fizz",
      "Buzz",
      "11",
      "Fizz",
      "13",
      "14",
      "FizzBuzz",
    ]);
  });
});
