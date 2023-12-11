const listPalindrome = (head) => {
  if (!head) return null;
  let start = head;

  const values = [];

  while (head) {
    values.push(head.val);
    head = head.next;
  }

  values.reverse();

  for (let val of values) {
    if (val !== start.val) {
      return false;
    }
    start = start.next;
  }

  return true;
};

module.exports = listPalindrome;
