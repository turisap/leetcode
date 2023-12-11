const alphabetic_war = (string) => {
  if (!/#/.test(string)) {
    return string.replace(/[#,\],\[]/g, "");
  }

  if (!/\]/.test(string)) {
    return "";
  }

  if (string.match(/#/g).length === 1) {
    return string
      .match(/\[\w+]/g)
      .toString()
      .replace(/[\[\]]/g, "")
      .replace(/,/g, "");
  }

  let groups = [];
  let brCount = 0;
  let startNext = 0;
  let startThis = 0;

  for (let i = 0; i < string.length; i++) {
    if (string[i] === "[") brCount++;
    if (string[i] === "]") startNext = i;

    if (brCount === 2) {
      groups.push(string.substring(startThis, i));
      startThis = startNext + 1;
      brCount = 0;
      i = startThis;
    }

    if (i === string.length - 1) {
      groups.push(string.substring(startThis, i + 1));
    }
  }


  const alive = groups.reduce((tail, current) => {
    if ((current.match(/#/g) || []).length > 1) {
      return tail;
    } else {
      return (tail += current.substring(
        current.indexOf("[") + 1,
        current.indexOf("]")
      ));
    }
  }, "");

  return alive;
};

module.exports = alphabetic_war;
