const rangeSum = (root, low, high) => {
  if (!root) return 0;
  let sum = 0;

  const inner = (root) => {
    if (!root) return;

    if (root.val >= low && root.val <= high) {
      sum += root.val;
    }

    inner(root.left);
    inner(root.right);
  };

  inner(root);

  return sum;
};

const rangeSumIterative = (root, l, h) => {
  if (!root) return false;

  const stack = [root];
  let res = 0;

  while (stack.length) {
    const curr = stack.pop();

    if (curr.val >= l && curr.val <= h) {
      res += curr.val;
    }

    if (curr.left) stack.push(curr.left);
    if (curr.right) stack.push(curr.right);
  }

  return res;
};

module.exports = rangeSumIterative;
