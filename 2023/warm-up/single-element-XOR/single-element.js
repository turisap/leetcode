// @TODO good O(n), bad memory
// const singleElement = (arr) => {
//   const map = new Map();
//
//   for (let number of arr) {
//     if (map.has(number)) {
//       map.delete(number);
//     } else {
//       map.set(number, true);
//     }
//   }
//
//   return Array.from(map.keys())[0];
// };

// @TODO good memory usage, bad O(n)
// const singleElement = (arr) => {
//   for (let i = 0; i < arr.length; i++) {
//     if (arr.lastIndexOf(arr[i]) === arr.indexOf(arr[i])) {
//       return arr[i];
//     }
//   }
// };

const singleElement = (arr) => {
  let res = 0;

  for (let num of arr) {
    res ^= num;
  }

  return res;
};

module.exports = singleElement;
