const closestValue = (root, target) => {
  let closest = root.val;
  let diff = Math.abs(target - root.val);
  const searchSubtree = (node) => {
    if (!node) return;

    const currDiff = Math.abs(target - node.val);

    if (currDiff < diff) {
      diff = currDiff;
      closest = node.val;
    }

    if (target > node.val) {
      searchSubtree(node.right);
    } else {
      searchSubtree(node.left);
    }
  };

  searchSubtree(root);

  return closest;
};

const closestValueIterative = (root, target) => {
  if (!root) return 0;

  const stack = [root];
  let closest = root.val;
  let diff = Math.abs(target - root.val);

  while (stack.length) {
    const curr = stack.pop();
    const currDiff = Math.abs(target - curr.val);

    if (currDiff < diff) {
      closest = curr.val;
      diff = currDiff;
    }

    curr.left && stack.push(curr.left);
    curr.right && stack.push(curr.right);
  }

  return closest;
};

module.exports = closestValueIterative;
