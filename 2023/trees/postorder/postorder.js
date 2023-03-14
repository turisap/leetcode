const postorderTraversal = (root) => {
  const stack = [root];
  const res = [];

  while (stack.length) {
    const curr = stack.pop();
    res.push(curr.val);

    if (curr.right) {
      stack.push(curr.right);
    }

    if (curr.left) {
      stack.push(curr.left);
    }
  }

  return res;
};

const postOrderRecursive = (root) => {
  if (!root) return [];
  const res = [];

  const inner = (root) => {
    if (!root) return null;

    res.push(root.val);
    inner(root.left);
    inner(root.right);
  };

  inner(root);

  return res;
};

module.exports = postOrderRecursive;
