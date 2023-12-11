const binarySum = require("./binary-sum");
describe("it sums up two binary strings", () => {
  it("sums 11 and 1", () => {
    const res = binarySum("11", "1");

    expect(res).toBe("100");
  });

  it("sums 1010 and 1011", () => {
    const res = binarySum("1010", "1011");

    expect(res).toBe("10101");
  });

  it("sums 1111 and 1111", () => {
    const res = binarySum("1111", "1111");

    expect(res).toBe("11110");
  });
});

// 1010
// 1011
//    11101
