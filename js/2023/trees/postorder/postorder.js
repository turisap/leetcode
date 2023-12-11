const postOrderIterative = (root) => {
  if (!root) return;
  const stack = [root];
  const res = [];

  while (stack.length) {
    root = stack.pop();
    res.push(root.val);

    if (root.left) stack.push(root.left);
    if (root.right) stack.push(root.right);
  }

  return res.reverse();
};

const postOrderRec = (root) => {
  if (!root) return [];
  const res = [];

  const inner = (root) => {
    if (!root) return;

    inner(root.left);
    inner(root.right);
    res.push(root.val);
  };

  inner(root);

  return res;
};

module.exports = postOrderRec;
