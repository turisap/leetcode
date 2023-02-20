// naive
const hasCycle = (head) => {
  let curr = head.next;
  head.visited = true;

  while (curr) {
    if (curr.visited) {
      return true;
    }

    curr.visited = true;
    curr = curr.next;
  }

  return false;
};

// slow - fast
const hasCycleTwoPointers = (head) => {
  let slow = head;
  let fast = head.next;

  while (slow !== fast) {
    if (!slow || !fast) {
      return false;
    }

    slow = slow.next;
    fast = fast?.next?.next;
  }

  return true;
};

module.exports = hasCycle;
