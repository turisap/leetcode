const input = [3, 4, 5, 6, 7, 8, 0, 1, 2]


const binaryShift = (arr, target) => {
  let leftPointer = 0;
  let rightPointer = arr.length - 1;

  while(leftPointer <= rightPointer){
    let mid = Math.floor((leftPointer + rightPointer) / 2);

    if(arr[mid] === target) return mid

    if(arr[leftPointer] <= arr[mid]) { // left part is sorted search in this part
      if(arr[leftPointer] <= target && target <= arr[mid]) {
	rightPointer = mid - 1;
      } else {
	leftPointer = mid + 1
      }
    } else { // then the right part is sorted search in this part
      if(arr[mid] <= target && target <= arr[rightPointer]) {
	leftPointer = mid + 1
      } else {
	rightPointer = mid - 1
      }
    }
  }

  return -1
}

console.log(binaryShift(input, 12))
