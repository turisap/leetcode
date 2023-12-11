// CHECKLIST
// Write a few test examples. This ensures that I understood the task

/*
- Think out loud the algorithm and the way (sliding window, two pointers) you want to implement it.
    - declare all variables you need
    - body of the algorithm
    - returning result
*/

// Think about bad input (at least ask about this)

const mergeTwoLists = (list1, list2) => {
  let c1 = 0;
  let c2 = 0;
  const res = [];

  if (!list2?.length) return list1;
  if (!list1?.length) return list2;

  while (c1 < list1.length || c2 < list2.length) {
    const curr1 = list1[c1];
    const curr2 = list2[c2];

    if (curr1 <= curr2) {
      res.push(curr1);
      c1++;
      continue;
    }

    res.push(curr2);
    c2++;
  }

  return res;
};

module.exports = mergeTwoLists;
