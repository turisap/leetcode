const validateBST = (root) => {
  if (!root) return false;
  const vals = [];

  const dfs = (root) => {
    if (!root) return true;

    dfs(root.left);
    vals.push(root.val);
    dfs(root.right);
  };

  dfs(root);

  for (let i = 1; i < vals.length; i++) {
    if (vals[i - 1] >= vals[i]) return false;
  }

  return true;
};

module.exports = validateBST;
