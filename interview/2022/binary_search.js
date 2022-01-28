const arr = [-1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

const binarySearch = (input, target) => {
  if(arr.length < 2) return arr;

  let left = 0;
  let right = arr.length - 1;

  while (left <= right){
    const mid = Math.floor( (left + right) / 2 )

    if(input[mid] === target) return mid

    if(target > input[mid]) {
      left = mid + 1;
    } else {
      right = mid - 1;
    }
  }

  return -1;
}

console.log(binarySearch(arr, 6))
