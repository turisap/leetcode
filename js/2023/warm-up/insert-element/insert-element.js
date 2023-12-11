// naive solution
// const insertElement = (arr, target) => {
//   let p = 0;
//
//   while (p < arr.length) {
//     const curr = arr[p];
//
//     if (curr < target) {
//       p++;
//       continue;
//     }
//
//     return p;
//   }
//
//   return p;
// };

const insertElement = (arr, target) => {
  let left = 0;
  let right = arr.length - 1;
  let middle;

  if (target < arr[left]) return 0;

  if (target > arr[right]) return arr.length;

  while (left <= right) {
    middle = Math.floor((left + right) / 2);

    if (target === arr[middle]) {
      return middle;
    } else if (target > arr[middle]) {
      left = middle + 1;
    } else if (target < arr[middle]) {
      right = middle - 1;
    }
  }

  return left;
};

module.exports = insertElement;
