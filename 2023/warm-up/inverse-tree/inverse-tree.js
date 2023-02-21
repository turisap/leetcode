const inverseTree = (root) => {
  const inverseRec = (node) => {
    if (!node) return;

    const tempLeft = node.left;

    node.left = node.right;
    node.right = tempLeft;

    inverseRec(node.right);
    inverseRec(node.left);

    return node;
  };

  inverseRec(root);

  console.log(root);
  return root;
};

const inverseTreeIterative = (root) => {
  if (!root) return root;

  const stack = [root];

  while (stack.length) {
    const curr = stack.pop();
    const left = curr?.left;
    const right = curr?.right;

    curr.right = left;
    curr.left = right;

    if (left) stack.push(left);
    if (right) stack.push(right);
  }

  return root;
};

module.exports = inverseTree;
