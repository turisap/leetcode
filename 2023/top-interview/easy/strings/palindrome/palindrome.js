const isPalindrome = (str) => {
  let escaped = str.replace(/[^A-Za-z0-9]/g, "").toLowerCase();
  let pL = 0;
  let pR = escaped.length - 1;

  while (pL <= pR) {
    if (escaped[pL] !== escaped[pR]) return false;
    pL++;
    pR--;
  }

  return true;
};

module.exports = isPalindrome;
