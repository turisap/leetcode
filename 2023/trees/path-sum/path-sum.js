// const pathSum = (root, target) => {
//   if (!root) return false;
//
//   target -= root.val;
//
//   if (!root.left && !root.right) return target === 0;
//
//   return pathSum(root.left, target) || pathSum(root.right, target);
// };

const pathSumIterative = (root, target) => {
  if (!root) return false;

  const stack = [[root, target]];

  while (stack.length) {
    let [node, remain] = stack.pop();

    remain -= node.val;

    if (!node.right && !node.left && remain === 0) {
      return true;
    }

    if (node.right) stack.push([node.right, remain]);
    if (node.left) stack.push([node.left, remain]);
  }

  return false;
};

module.exports = pathSumIterative;
