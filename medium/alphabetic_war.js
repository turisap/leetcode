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
      .replace(/[\[\]]/g, "");
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
  }

  console.log(groups);
};

alphabetic_war("[ab]adfd[dd]##[abe]dedf[ijk]d#d[h]#");

module.exports = alphabetic_war;
