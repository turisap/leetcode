const inorderTraversalIter = (root) => {
  if (!root) return [];

  const stack = [];
  const res = [];
  let curr = root;

  while (curr || stack.length) {
    if (curr) {
      stack.push(curr);
      curr = curr.left;
    } else {
      curr = stack.pop();
      res.push(curr.val);
      curr = curr.right;
    }
  }

  return res;
};
const inorderTraversal = (root) => {
  if (!root) return [];
  const res = [];

  const dfs = (root) => {
    if (!root) {
      return null;
    }

    dfs(root.left);
    res.push(root.val);
    dfs(root.right);
  };

  dfs(root);
  return res;
};

module.exports = inorderTraversal;
