const first_unique_char = (string) => {
  const chars = [...string.split("")];

  for (let i = 0; i < chars.length; i++) {
    const char = chars[i];
    const matches = string.match(RegExp(char, "g"));

    if (matches.length === 1) {
      return char;
      break;
    }
  }
};

const char = first_unique_char("sisi lal o e");
console.log(char);
