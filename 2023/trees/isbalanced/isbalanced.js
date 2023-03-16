const isBalanced = (root) => {
  if (!root) return true;

  const hL = (root.left && height(root.left)) || 0;
  const hR = (root.right && height(root.right)) || 0;

  return (
    Math.abs(hL - hR) <= 1 && isBalanced(root.left) && isBalanced(root.right)
  );
};

const height = (root) => {
  if (!root.left && !root.right) return 1;

  const heightRight = (root.right && height(root.right)) || 0;
  const heightLeft = (root.left && height(root.left)) || 0;

  return 1 + Math.max(heightLeft, heightRight);
};

module.exports = isBalanced;
