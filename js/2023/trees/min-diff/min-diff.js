const minDiff = (root) => {
  if (!root) return 0;
  let min = Infinity,
    prev;

  const dfs = (node) => {
    if (!node) return;
    dfs(node.left);

    if (prev) {
      const currDiff = Math.abs(node.val - prev.val);
      min = currDiff < min ? currDiff : min;
    }

    prev = node;

    dfs(node.right);
  };

  dfs(root);

  return min;
};

module.exports = minDiff;
