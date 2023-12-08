// CHECKLIST
// Write a few test examples. This ensures that I understood the task

/*
- Think out loud the algorithm and the way (sliding window, two pointers) you want to implement it.
    - declare all variables you need
    - body of the algorithm
    - returning result
*/

// Think about bad input (at least ask about this)

// [1, 3, 5, 6], 6
const searchInsert = (arr, target) => {
  let left = 0;
  let right = arr.length - 1;

  while (left <= right) {
    let mid = Math.floor((left + right) / 2);

    if (target === arr[mid]) {
      return mid;
    } else if (target < arr[mid]) {
      right = mid - 1;
    } else if (target > arr[mid]) {
      left = mid + 1;
    }
  }

  return left;
};

module.exports = searchInsert;
