const maxTreeDepthRecursive = (root) => {
  if (!root) return 0;

  const leftDepth = maxTreeDepthRecursive(root.left);
  const rightDepth = maxTreeDepthRecursive(root.right);

  return Math.max(leftDepth, rightDepth) + 1;
};

const maxTreeDepthIterative = (root) => {
  if (!root) return 0;

  let maxDepth = 0;
  const stack = [{ node: root, depth: 1 }];

  while (stack.length) {
    const el = stack.pop();
    const left = el.node.left;
    const right = el.node.right;

    maxDepth = Math.max(el.depth, maxDepth);

    left && stack.push({ node: left, depth: el.depth + 1 });
    right && stack.push({ node: right, depth: el.depth + 1 });
  }

  return maxDepth;
};

module.exports = maxTreeDepthIterative;
