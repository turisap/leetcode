const isSymmetricRecursive = function (root) {
  const isSame = (left, right) => {
    if (!left && !right) return true;

    if (!left && right) return false;

    if (left && !right) return false;

    if (left.val !== right.val) return false;

    return isSame(left.left, right.right) && isSame(left.right, right.left);
  };

  return isSame(root.left, root.right);
};

const isSymmetricTreeIterative = function (root) {
  const queue = [root.left, root.right];

  while (queue.length) {
    const left = queue.shift();
    const right = queue.pop();

    if (!left && !right) continue;

    if (!left && right) return false;

    if (left && !right) return false;

    if (left.val !== right.val) return false;

    queue.unshift(left.left ?? null);

    queue.push(right.right ?? null);

    queue.unshift(left.right ?? null);

    queue.push(right.left ?? null);
  }

  return !queue.length;
};

module.exports = isSymmetricTreeIterative;
