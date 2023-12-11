const deleteElement = (head, n) => {
  let slow = head;
  let fast = head;

  for (let i = 0; i < n; i++) {
    fast = fast.next;
  }

  if (!fast) return head.next;

  while (fast && fast.next) {
    slow = slow.next;
    fast = fast.next;
  }

  slow.next = slow.next.next;

  return head;
};

module.exports = deleteElement;
