const hammingDistance = (a, b) => {
  let res = a ^ b;
  let mask = 1;
  let c = 0;

  for (let i = 0; i < 32; i++) {
    if ((mask & res) !== 0) c++;
    mask = mask << 1;
  }

  return c;
};

module.exports = hammingDistance;
