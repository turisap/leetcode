const climbStairs = (n) => {
  // const inner = (n) => {
  //   if (n === 1) return 1;
  //   if (n === 2) return 2;
  //   if (n === 3) return 3;
  //
  //   return inner(n - 1) + inner(n - 2);
  // };
  //
  // return inner(n);

  if (n <= 3) return n;

  let prev = 2;
  let curr = 3;

  for (let i = 4; i <= n; i++) {
    let temp = curr;
    curr = prev + curr;
    prev = temp;
  }

  return curr;
};

module.exports = climbStairs;
