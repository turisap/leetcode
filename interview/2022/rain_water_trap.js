// new trap on going down
//
//
const input1 = [0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]
const input2 = [10, 0, 8, 0, 5]
const input3 = [0, 1, 2, 3, 2, 3, 2, 1, 0]
const input4 = [5, 2, 5, 1, 1]

const getAmount = (layout) => {
  let total = 0
  let leftPointer = 0
  let rightPointer = layout.length - 1
  let leftMax = layout[0]
  let rightMax = layout[layout.length - 1]

  while(leftPointer <= rightPointer){
    leftMax = Math.max(leftMax, layout[leftPointer])
    rightMax = Math.max(rightMax, layout[rightPointer])
    const leftBox = layout[leftPointer]
    const rightBox = layout[rightPointer]

    if(leftBox < rightBox){
      total += Math.max(leftMax - leftBox, 0)
      leftPointer++
    } else  {
      total += Math.max(rightMax - rightBox, 0)
      rightPointer--
    }
  }

  return total
}

console.log(getAmount(input1))
console.log(getAmount(input2))
console.log(getAmount(input3))
console.log(getAmount(input4))

