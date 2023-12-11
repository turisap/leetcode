const lastWord = (str) => {
  const trimmed = str.trim();
  let count = 0;
  let i = trimmed.length - 1;

  while (i >= 0) {
    if (trimmed[i] === " ") {
      break;
    }

    count++;
    i--;
  }

  return count;
};

console.log("RESULT", lastWord("Hello World"));
console.log("RESULT", lastWord("   fly me   to   the moon  "));
console.log("RESULT", lastWord("luffy is still joyboy"));
console.log("RESULT", lastWord(""));
