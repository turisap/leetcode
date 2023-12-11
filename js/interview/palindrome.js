const palindrome = (string) => {
  if (string.length % 2 === 0) return false;
  const lastI = string.length - 1;

  for (let i = 0; i < string.length; i++) {
    if (string[i] !== string[lastI - i]) return false;
  }

  return true;
};

palindrome("mom");
palindrome("adjf");
palindrome("refer");
palindrome("mama");
