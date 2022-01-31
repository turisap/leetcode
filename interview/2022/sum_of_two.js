const input = [2, 7, 11, 15]
const input2 = [4, 13, 7, 8, 19, 7, 2, 1, 5]
const input3 = [1, 5, 12, 3]

const sumOfTwo = (arr, target) => {
  const diffMap = {}

  for(let num of arr){
    diffMap[target - num] = num
  }

  for(let num of arr){
    if(diffMap[num]){
      return [diffMap[num], num]
    }
  }

  return []
}

console.log(sumOfTwo(input, 9))
console.log(sumOfTwo(input2, 18))
console.log(sumOfTwo(input3, 9))
