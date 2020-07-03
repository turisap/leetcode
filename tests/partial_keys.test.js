const partial_keys = require("../medium/proxy_keys.js");

describe("returns value for a matched partial key sorted in alphabetic order", () => {
  let testObj = { aaa: 1, abc: 3, dfg: 3, def: 4 };
  let proxy = null;

  beforeEach(() => (proxy = partial_keys(testObj)));

  test("should return 1 for 'a' and 'aa'", () => {
    expect(proxy.a).toBe(1);
  });

  test("testObj.abc should return 3", () => {
    expect(proxy.ab).toBe(3);
  });

  test("should return 4 for 'd' and 'de' ", () => {
    expect(proxy.d).toBe(4);
    expect(proxy.de).toBe(4);
  });

  test("should  return null for 'l' key", () => {
    expect(proxy.l).toBe(null);
  });
});
