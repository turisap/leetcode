const inorderTraversal = (root) => {
  if (!root) return [];

  const stack = [root];
  const res = [];
  let curr = root;

  while (stack.length) {
    if (curr.left) {
      stack.push(curr.left);
      curr = curr.left;
      continue;
    }

    curr = stack.pop();
    res.push(curr.val);

    if (curr.right) {
      stack.push(curr.right);
      curr = curr.right;
    }
  }

  return res;
};

module.exports = inorderTraversal;
