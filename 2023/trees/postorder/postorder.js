const postorderTraversal = (root) => {
  const stack = [root];
  const res = [];

  while (stack.length) {
    const curr = stack.pop();
    res.push(curr.val);

    if (curr.right) {
      stack.push(curr.right);
    }

    if (curr.left) {
      stack.push(curr.left);
    }
  }

  return res;
};

module.exports = postorderTraversal;
