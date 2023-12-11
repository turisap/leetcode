const openMap = {
  "{": true,
  "[": true,
  "(": true,
};

const matchMap = {
  "{": "}",
  "[": "]",
  "(": ")",
};

const validParentheses = (str) => {
  if (str.length % 2 !== 0) {
    return false;
  }

  const stack = [];

  for (let char of str) {
    if (openMap[char]) {
      stack.push(char);
    } else {
      const open = stack.pop();

      if (matchMap[open] !== char) {
        return false;
      }
    }
  }

  return Boolean(!stack.length);
};

console.log(validParentheses("()"));
console.log(validParentheses("([{}])"));
console.log(validParentheses("()[]{}"));
console.log(validParentheses("(]"));
console.log(validParentheses("["));
