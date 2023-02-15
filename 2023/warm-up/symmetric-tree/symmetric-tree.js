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

module.exports = isSymmetricRecursive;
