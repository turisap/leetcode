const input = [1, 1, 1, 2, 2]
const input2 = [1, 2, 3]
const input3 = [1, 2, 2, 3, 3, 3, 3, 3, 3, 4, 4, 5]

const removeDuplicates = (arr) => {
  let i = 0;
  let j = 0;

  while(i < arr.length){
    if(arr[i] !== arr[i - 1]) {
      arr[j] = arr[i]
      i++
      j++
    } else {
      arr[j] = arr[i]
      i++
    }
  }


  arr.length = j

  return arr.length
}

console.log(removeDuplicates(input), input)
console.log(removeDuplicates(input2), input2)
console.log(removeDuplicates(input3), input3)
