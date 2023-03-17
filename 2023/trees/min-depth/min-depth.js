const minTreeDepthRecur = (root) => {
  if (!root) return 0;

  if (!root.left && !root.right) return 1;

  if (root.left && !root.right) {
    return 1 + minTreeDepthRecur(root.left);
  }

  if (!root.left && root.right) {
    return 1 + minTreeDepthRecur(root.right);
  }

  return (
    1 + Math.min(minTreeDepthRecur(root.right), minTreeDepthRecur(root.left))
  );
};

const minDepth = (root) => {
  if (!root) return 0;

  const stack = [[root, 1]];
  const minDepth = Number.MAX_SAFE_INTEGER;

  while (stack.length) {
    const [node, depth] = stack.pop();

    if (!node.right && !node.left) {
      return Math.min(minDepth, depth);
    }

    if (node.right) stack.push([node.right, depth + 1]);
    if (node.left) stack.push([node.left, depth + 1]);
  }

  return minDepth;
};

module.exports = minDepth;
