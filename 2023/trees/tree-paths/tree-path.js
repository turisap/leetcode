const getPath = (node, path) =>
  path.length ? `${path}->${node.val}` : `${node.val}`;
const binaryTreePaths = (root) => {
  if (!root) return [];

  const res = [];

  const inner = (node, path) => {
    if (!node) return;

    if (!node?.left && !node?.right && node) {
      res.push(getPath(node, path));
      return;
    }

    path = getPath(node, path);

    inner(node.left, path);
    inner(node.right, path);
  };

  inner(root, ``);

  return res;
};

const binaryTreePathIterative = (root) => {
  if (!root) return [];
  const stack = [[root, ""]];
  const res = [];

  while (stack.length) {
    let [node, path] = stack.pop();

    if (!node.left && !node.right) {
      res.push(getPath(node, path));
    }

    path = getPath(node, path);

    node.left && stack.push([node.left, path]);
    node.right && stack.push([node.right, path]);
  }

  return res;
};

module.exports = binaryTreePathIterative;
