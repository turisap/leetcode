const buyStock = require("./buy-stock");
describe("it should max profit", () => {
  it("returns 5", () => {
    const res = buyStock([7, 1, 5, 3, 6, 4]);
    expect(res).toBe(5);
  });

  it("returns 0", () => {
    const res = buyStock([7, 6, 4, 3, 1]);
    expect(res).toBe(0);
  });

  it("returns 6", () => {
    const res = buyStock([1, 2, 3, 4, 5, 6, 7]);
    expect(res).toBe(6);
  });

  it("returns 0", () => {
    const res = buyStock([]);
    expect(res).toBe(0);
  });
});
