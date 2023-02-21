const reverseLinkedList = (head) => {
  let curr = head;
  let prev = null;

  while (curr) {
    const after = curr.next;
    curr.next = prev;
    prev = curr;

    if (after) {
      curr = after;
    } else {
      break;
    }
  }

  return curr;
};

const reverseLinkedListRecur = (head) => {
  if (!head || !head.next) return head;

  const pointer = reverseLinkedListRecur(head.next);
  head.next.next = head;
  head.next = null;
  return pointer;
};

console.log(
  "RESULT",
  reverseLinkedList({
    val: 1,
    next: {
      val: 2,
      next: {
        val: 3,
      },
    },
  })
);

module.exports = reverseLinkedList;
