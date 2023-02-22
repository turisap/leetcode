const treeDiameter = (root) => {
  let diameter = 0;

  const longestPath = (node) => {
    if (!node) return 0;

    const left = longestPath(node.left);
    const right = longestPath(node.right);

    diameter = Math.max(diameter, left + right);

    return Math.max(left, right) + 1;
  };

  longestPath(root);

  return diameter;
};

//[2,3,null,1]

const tree = {
  val: 2,
  left: {
    val: 3,
    left: {
      val: 1,
    },
  },
};

console.log("RESULT", treeDiameter(tree));

module.exports = treeDiameter;
