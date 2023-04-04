const binaryOnesCount = (n) => {
  // ============== 1 ================
  // const rg = new RegExp(/1/, "g");
  //
  // return n.toString(2).match(rg)?.length || 0;
  // ============== 2 ================
  // return n.toString(2).replace(/0/g, "")?.length || 0;

  // =======================  bits manipulation =============
  // let bits = 0;
  // let mask = 1;
  //
  // for (let i = 0; i < 32; i++) {
  //   if ((mask & n) !== 0) bits++;
  //
  //   mask = mask << 1;
  // }
  //
  // return bits;

  let bits = 0;

  while (n !== 0) {
    bits++;
    n &= n - 1;
  }

  return bits;
};

module.exports = binaryOnesCount;
