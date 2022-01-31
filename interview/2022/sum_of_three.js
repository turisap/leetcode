const arr = [2, 5, 4, 1, 9, -3, 0,  8, 7]
const arr1 = [1, 5, 9, 3, 11, 2, 16]

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


const sumOfThree = (arr, target) => {
  let res = [];

  if(arr.length < 3){
    return res
  }

  for(let i = 0; i < arr.length; i++){
    const num = arr[i]
    const copy = [...arr]
    copy.splice(i, 1)
    const sumRest = sumOfTwo(copy, target - num)

    if(sumRest.length){
      res.push([num, ...sumRest])
    }
  }

  return res
}

console.log(sumOfThree(arr1, 18))
