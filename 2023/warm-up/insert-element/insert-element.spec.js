const insertElement = require("./insert-element");
describe("it should return insertion index", () => {
  it("returns 5 for [1, 3, 5, 6]", () => {
    const res = insertElement([1, 3, 5, 6], 5);

    expect(res).toBe(2);
  });

  it("returns 1 for [1, 3, 5, 6] and 2", () => {
    const res = insertElement([1, 3, 5, 6], 2);

    expect(res).toBe(1);
  });

  it("returns for [1, 3, 5, 6] and 7", () => {
    const res = insertElement([1, 3, 5, 6], 7);

    expect(res).toBe(4);
  });

  it("returns for [1,3,5,6]  and 0", () => {
    const res = insertElement([1, 3, 5, 6], 0);

    expect(res).toBe(0);
  });

  it("returns for [1,3]  and 2", () => {
    const res = insertElement([1, 3], 2);

    expect(res).toBe(1);
  });

  it("returns for [1]  and 1", () => {
    const res = insertElement([1], 1);

    expect(res).toBe(0);
  });
});
