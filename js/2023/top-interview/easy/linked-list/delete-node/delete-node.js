const deleteNode = (node) => {
  const temp = node.next;
  node.val = temp.val;
  node.next = temp.next;
};

module.exports = deleteNode;
