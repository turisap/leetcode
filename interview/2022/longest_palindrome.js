const str1 = 'babad'
const str2 = 'mississippi'
const str3 = 'cbbd'
const str4 = 'sb'

const longestPolyndrome = (string) => {
  let longest  = ''
  
  for(let i = 1; i < string.length; i++){
    let leftPointer = i - 1;
    let rightPointer = i + 1;
    let res = string[0]

    while(leftPointer >= 0 || rightPointer < string.length){
      if(string[leftPointer] === string[rightPointer]) {
	res = string.substring(leftPointer, rightPointer + 1)
      } else if(string[leftPointer + 1] === string[rightPointer]){
	res = string.substring(leftPointer +1, rightPointer + 1)
      } else {
	break;
      } 

      leftPointer--
      rightPointer++
    }

    if(longest.length < res.length) longest = res

  }

  return longest
}

console.log(longestPolyndrome(str1))
console.log(longestPolyndrome(str2))
console.log(longestPolyndrome(str3))
console.log(longestPolyndrome(str4))
