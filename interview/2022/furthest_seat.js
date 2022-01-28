const input1 = [1, 0, 0, 0, 1, 0, 1];
const input2 = [1, 0, 0, 0]
const input3 = [0, 0, 0, 1]
const input4 = [1, 0, 0, 0, 1, 0]

const maxDistToClosest = seats => {
  let max = 0
  let currMax = 0
  let i = 0;

  if(seats[0] === 0){
    while(seats[i] === 0) {
      i++
      currMax++
    }

    max = currMax;
    currMax = 0;
  }

  for(; i < seats.length; i++) {
    const current  = seats[i]

    if(i === seats.length - 1 && current === 0){
      currMax++
      max = Math.max(max, currMax)
      break;
    }

    if(current === 0){
      currMax++
    } else if (current === 1){
      max = Math.max(max, Math.ceil(currMax / 2))
      currMax = 0;
    }
  }

  return max;
}

console.log(maxDistToClosest(input1))
console.log(maxDistToClosest(input2))
console.log(maxDistToClosest(input3))
console.log(maxDistToClosest(input4))
