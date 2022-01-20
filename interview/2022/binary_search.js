const arr = [-1, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

const binarySearch = (input, subject) => {
  let left = 0;
  let right = input.length - 1;
  let mid = null

  while(left <= right){
    mid = Math.round((left + right) / 2)

    if(input[mid] === subject) return mid
    
    if(input[mid] < subject) {
      left = mid + 1;
    }

    if(input[mid] > subject) {
      right = mid - 1
    }
  }
}

console.log(binarySearch(arr, 2))
