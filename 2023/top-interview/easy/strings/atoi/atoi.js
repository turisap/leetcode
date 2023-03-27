const myAtoi = (str) => {
  // if (!str) return 0;
  let toParse = "";
  const trimmed = str.trimStart();
  const negative = trimmed[0] === "-";
  const escaped = trimmed.replace(/^[-,+]/, "");

  for (let char of escaped) {
    if (/^\d/.test(char)) toParse += char;
    else break;
  }

  if (negative) {
    return Number(toParse) > Math.pow(2, 31)
      ? -Math.pow(2, 31)
      : -Number(toParse);
  } else {
    return Number(toParse) > Math.pow(2, 31) - 1
      ? Math.pow(2, 31) - 1
      : Number(toParse);
  }
};

module.exports = myAtoi;
