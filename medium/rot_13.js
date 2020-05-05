const rot_13 = (input) => {
  let res = "";
  let code = null;

  const getCode = (char, base) =>
    ((char.charCodeAt(0) - base + 13) % 26) + base;

  input.split("").forEach((char) => {
    if (/[A-Z]/.test(char)) {
      code = getCode(char, 65);
      res += String.fromCharCode(code);
    } else if (/[a-z]/.test(char)) {
      code = getCode(char, 97);
      res += String.fromCharCode(code);
    } else {
      res += char;
    }
  });
  return res;
};
