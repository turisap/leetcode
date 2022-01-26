const input = [1, 5, 9, 13, 16, 8, 3]
const input2 = [0, 0, 0, 1, 0, 0]
const input3 = [0, 1, 2, 3, 4, 5]
const input4 = [5, 3, 2, 1, 0]


const getRangeTop = (attitudes) => {
  let leftPointer = 0;
  let rightPointer = attitudes.length - 1;
  
  while (leftPointer <= rightPointer) {
    const middle = Math.floor((leftPointer + rightPointer) / 2) 

    // don't forget to check if left and rights are in the input length
    if((attitudes[middle - 1] || 0) <= attitudes[middle] && attitudes[middle] >= (attitudes[middle + 1]  || 0)) {
      return attitudes[middle]
    }

    if(attitudes[middle - 1] <= attitudes[middle]) {
      leftPointer = middle + 1
    } else {
      rightPointer = middle - 1
    }
  }

  return -1
}

console.log(getRangeTop(input))
console.log(getRangeTop(input2))
console.log(getRangeTop(input3))
console.log(getRangeTop(input4))
