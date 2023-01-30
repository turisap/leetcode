// const binarySum = (a1, a2) => {
//   const arr1 = a1.split("");
//   const arr2 = a2.split("");
//   let carry = 0;
//   let res = [];
//
//   while (arr1.length || arr2.length) {
//     const curr1 = arr1.pop() || 0;
//     const curr2 = arr2.pop() || 0;
//
//     carry += parseInt(curr1);
//     carry += parseInt(curr2);
//
//     res.push(carry % 2);
//
//     carry = Math.floor(carry / 2);
//   }
//
//   if (carry === 1) res.push("1");
//
//   console.log("RESULT", res);
//
//   return res.reverse().join("");
// };

const binarySum = (a, b) => {
  const sum = BigInt(`0b${a}`) + BigInt(`0b${b}`);

  return sum.toString(2);
};

module.exports = binarySum;
