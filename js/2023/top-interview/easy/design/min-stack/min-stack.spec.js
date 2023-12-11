const MinStack = require("./min-stack");
describe("", () => {
  it("should return ", () => {
    const s = new MinStack();

    s.push(null);
    s.push(-2);
    s.push(0);
    s.push(-3);

    expect(s.getMin()).toBe(-3);

    s.pop();
    expect(s.top()).toBe(0);
    expect(s.getMin()).toBe(-2);
  });

  it("should return ", () => {
    const s = new MinStack();

    s.push(0);
    s.push(1);
    s.push(0);

    expect(s.getMin()).toBe(0);
    s.pop();
    expect(s.getMin()).toBe(0);
  });
});
