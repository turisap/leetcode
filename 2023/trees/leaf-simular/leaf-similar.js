const leafSimilar = (r1, r2) => {
  if (!r1 || !r2) return false;

  return getLeafSequence(r1) === getLeafSequence(r2);
};

const getLeafSequence = (root) => {
  const stack = [root];
  let s = "";

  while (stack.length) {
    if (root?.left) {
      stack.push(root.left);
      root = root.left;

      continue;
    }

    root = stack.pop();

    if (!root.left && !root.right) {
      s += root.val;

      continue;
    }

    root.right && stack.push(root.right);
    root = root.right;
  }

  return s;
};

const getLeafSequenceRecur = (root) => {
  let s = "";
  const inner = (node) => {
    if (!node) return;

    if (!node.left && !node.right) s += `${node.val}`;

    inner(node.left);
    inner(node.right);
  };

  inner(root);

  return s;
};

module.exports = {
  getLeafSequenceRecur,
  leafSimilar,
};
