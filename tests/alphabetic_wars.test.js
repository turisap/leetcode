const alphabetic_war = require("../medium/alphabetic_war.js");

describe('it should return all chars bar "]", "[", "[" if there is no "#" char', () => {
  test('should return "abc" for "abc"', () => {
    expect(alphabetic_war("abc")).toBe("abc");
  });

  test('should return "abc" for "[abc]"', () => {
    expect(alphabetic_war("[abc]")).toBe("abc");
  });

  test('should return "abc" for "a[bc]"', () => {
    expect(alphabetic_war("a[bc]")).toBe("abc");
  });

  test('should return "abcde" for "[a]bc[de]"', () => {
    expect(alphabetic_war("[a]bc[de]")).toBe("abcde");
  });
});

describe('it should return an empty string if there is no "[" and at least one "#"', () => {
  test('should return "" for "abcdegkd#"', () => {
    expect(alphabetic_war("abcdegkd#")).toBe("");
  });
});

describe('it should return chars in square brackets only if there is only one "#"', () => {
  test('should return "abc" for g#ht[abc]tre', () => {
    expect(alphabetic_war("g#ht[abc]tre")).toBe("abc");
  });

  test('should return "a"  for ght[a]bfre#', () => {
    expect(alphabetic_war("ght[a]bfre#")).toBe("a");
  });
});
