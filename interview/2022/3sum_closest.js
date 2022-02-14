const sum3 = (numbers, target) => {
  const sorted = numbers.sort((a, b) => a - b);
  let diff = Infinity
  let answer = null

  for(let i = 0; i < sorted.length; i++){
    let k = i + 1
    let j = sorted.length - 1

    while(k < j) {
      const currSumm = sorted[i] + sorted[k] + sorted[j]
      
      if(currSumm === target) {
        return target
      }

      if(Math.abs(target - currSumm) < diff){
         answer = currSumm
         diff = Math.abs(target - currSumm)
      }

      if(currSumm > target){
      	j--
      } else {
      	k++
      }
    }

  }

  return answer
}

console.log(sum3([-1, 2, 1, -4],1 ))
console.log(sum3([1, 2, 3, 4, 5],1 ))
console.log(sum3([0,0,0], 1))
console.log(sum3([0,2,1,-3], 1))
console.log(sum3([1,2,4,8,16,32,64,128], 82))
