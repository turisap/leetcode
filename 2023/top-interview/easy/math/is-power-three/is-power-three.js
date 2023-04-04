const isPowerThree = (n) => {
  return /^10*$/.test(n.toString(3));
};

module.exports = isPowerThree;
