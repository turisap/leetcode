const input1 = [1, 8 ,6, 2, 5, 4, 8, 3, 7]

const maxArea = (lines) => {
  let leftPointer = 0
  let rightPointer = lines.length - 1
  let maxRes = 0

  while(leftPointer < rightPointer){
    const leftHeight = lines[leftPointer];
    const rightHeight = lines[rightPointer];
    const minLine = Math.min(leftHeight, rightHeight);
    const currentArea =  minLine * (rightPointer - leftPointer);

    if(currentArea > maxRes) {
      maxRes = currentArea;
    }

    if(leftHeight <= rightHeight){
      leftPointer += 1;
    } else if(leftHeight > rightHeight){
      rightPointer -= 1;
    }
  } 

  return maxRes;
}

console.log(maxArea(input1)) // logs 49

