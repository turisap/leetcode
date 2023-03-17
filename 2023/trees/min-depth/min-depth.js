const minTreeDepth = (root) => {
  if (!root) return 0;

  if (!root.left && !root.right) return 1;

  if (root.left && !root.right) {
    return 1 + minTreeDepth(root.left);
  }

  if (!root.left && root.right) {
    return 1 + minTreeDepth(root.right);
  }

  return 1 + Math.min(minTreeDepth(root.right), minTreeDepth(root.left));
};

module.exports = minTreeDepth;
