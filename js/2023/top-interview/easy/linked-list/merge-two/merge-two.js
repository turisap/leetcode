function ListNode(val, next) {
  this.val = val === undefined ? 0 : val;
  this.next = next === undefined ? null : next;
}
const mergeTwo = (h1, h2) => {
  if (!h1 && !h2) return null;
  if (!h1) return h2;
  if (!h2) return h1;

  let curr = new ListNode(-500000);
  let head = curr;

  while (h1 && h2) {
    if (h1.val < h2.val) {
      curr.next = h1;
      h1 = h1.next;
    } else {
      curr.next = h2;
      h2 = h2.next;
    }

    curr = curr.next;
  }

  if (!h1) curr.next = h2;
  else curr.next = h1;

  return head.next;
};

module.exports = mergeTwo;
