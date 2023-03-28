const levelOrder = (root) => {
  if (!root) return [];
  const res = [];
  const stack = [[root, 0]];

  while (stack.length) {
    const [node, level] = stack.pop();

    res[level] ? res[level].push(node.val) : (res[level] = [node.val]);

    node.right && stack.push([node.right, level + 1]);
    node.left && stack.push([node.left, level + 1]);
  }

  return res;
};

module.exports = levelOrder;
