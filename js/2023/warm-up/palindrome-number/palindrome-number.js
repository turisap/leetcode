function isNumberPalindrome(num) {
  if (num === 0) return true;
  // negative
  if (num % 10 === 0 || num < 0) {
    return false;
  }

  let reverted = 0;

  while (num > reverted) {
    reverted = reverted * 10 + Math.floor(num % 10);

    num = Math.floor(num / 10);
  }

  return reverted === num || Math.floor(reverted / 10) === num;
}

module.exports = isNumberPalindrome;
