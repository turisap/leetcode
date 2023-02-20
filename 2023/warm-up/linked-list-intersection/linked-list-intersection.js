// naive
const linkedListIntersection = (headA, headB) => {
  if (headA === headB) {
    return headA;
  }

  let currA = headA;

  while (currA) {
    let currB = headB;

    while (currB) {
      if (currA === currB) {
        return currB;
      }

      currB = currB.next;
    }

    currA = currA.next;
  }

  return null;
};

// two-pointers
const linkedListIntersectionTwoPointers = (headA, headB) => {
  let lengthA = 0;
  let lengthB = 0;
  let pA = headA;
  let pB = headB;

  while (pA) {
    lengthA++;
    pA = pA.next;
  }

  while (pB) {
    lengthB++;
    pB = pB.next;
  }

  const longerList = lengthA > lengthB ? headA : headB;
  const shorterList = lengthA > lengthB ? headB : headA;
  const lengthDiff = Math.abs(lengthA - lengthB);
  let longerStart = longerList;
  let shorterStart = shorterList;

  for (let i = 1; i <= lengthDiff; i++) {
    longerStart = longerStart.next;
  }

  while (longerStart && shorterStart) {
    if (longerStart === shorterStart) return longerStart;

    longerStart = longerStart.next;
    shorterStart = shorterStart.next;
  }

  return null;
};

module.exports = {
  linkedListIntersection,
  linkedListIntersectionTwoPointers,
};
