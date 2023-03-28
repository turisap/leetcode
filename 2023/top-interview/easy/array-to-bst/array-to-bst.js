function TreeNode(val, left, right) {
  this.val = val === undefined ? 0 : val;
  this.left = left === undefined ? null : left;
  this.right = right === undefined ? null : right;
}

const arrayToBst = (arr) => {
  if (!arr.length) return null;
  const inner = (left, right) => {
    if (left > right) return null;

    const middle = Math.floor((left + right) / 2);
    const root = new TreeNode(arr[middle]);

    root.left = inner(left, middle - 1);
    root.right = inner(middle + 1, right);

    return root;
  };

  return inner(0, arr.length - 1);
};

module.exports = arrayToBst;
