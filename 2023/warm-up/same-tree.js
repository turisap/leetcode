const isSameTreeRecursive = function (a, b) {
  if (!a && !b) return true;

  if ((!a && b) || (a && !b) || a.val !== b.val) return false;

  return (
    isSameTreeRecursive(a.left, b.left) && isSameTreeRecursive(a.right, b.right)
  );
};

const isSameTreeIterative = function (a, b) {
  const q = [a, b];

  while (q.length) {
    const left = q.pop();
    const right = q.pop();

    if (!left && !right) {
      continue;
    }

    if ((!left && right) || (left && !right)) return false;

    if (left.val !== right.val) return false;

    q.push(left.left, right.left, left.right, right.right);
  }

  return true;
};

const a = {
  val: 1,
  left: {
    val: 2,
  },
  right: {
    val: 1,
  },
};
const b = {
  val: 1,
  left: {
    val: 1,
  },
  right: {
    val: 2,
  },
};

const a1 = {
  val: 0,
  left: {
    val: -5,
  },
};

const b1 = {
  val: 0,
  left: {
    val: -8,
  },
};

// console.log(isSameTreeIterative(a, b));
console.log(isSameTreeIterative(a1, b1));
