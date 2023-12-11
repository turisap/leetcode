const squareRoot = require("./square-root");
describe("should find the clothest square root", () => {
  it("returns 2 for 4", () => {
    const res = squareRoot(4);

    expect(res).toBe(2);
  });

  it("calculates for 2", () => {
    const res = squareRoot(2);

    expect(res).toBe(2);
  });

  it("calculates for 5", () => {
    const res = squareRoot(5);

    expect(res).toBe(Math.floor(Math.sqrt(5)));
  });

  it("calculates for 6", () => {
    const res = squareRoot(6);

    expect(res).toBe(Math.floor(Math.sqrt(6)));
  });

  it("calculates for 10", () => {
    const res = squareRoot(10);

    expect(res).toBe(Math.floor(Math.sqrt(10)));
  });
});
