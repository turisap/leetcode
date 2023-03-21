const leftSum = (root) => {
  if (!root) return 0;
  let sum = 0;

  const inner = (node, rightChild) => {
    if (!node) return;

    if (rightChild && !node.left && !node.right) {
      sum += node.val;
    }

    inner(node.right, false);
    inner(node.left, true);
  };

  inner(root, false);

  return sum;
};

const leftSumIter = (root) => {
  if (!root) return 0;
  const stack = [[root, false]];
  let sum = 0;

  while (stack.length) {
    const [node, isLeftChild] = stack.pop();

    if (isLeftChild && !node.right && !node.left) {
      sum += node.val;
    }

    node.left && stack.push([node.left, true]);
    node.right && stack.push([node.right, false]);
  }

  return sum;
};

module.exports = leftSumIter;
