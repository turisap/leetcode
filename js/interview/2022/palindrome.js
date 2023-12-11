const input = 'А роза упала на лапу Азора'

const isPalindromeDifficult = (input) => {
  const chars = input.split('') 
  let leftPointer = 0;
  let rightPointer = chars.length - 1;
  let result = true

  while (leftPointer <= rightPointer){
    const leftChar = chars[leftPointer].toLowerCase()
    const rightChar = chars[rightPointer].toLowerCase()

    if(leftChar === ' '){
      leftPointer++;
    } else if(rightChar === ' ') {
      rightPointer--
    } else if( rightChar !== leftChar){
      result = false  
      break
    } else {
      leftPointer++
      rightPointer--
    }
  }

  return result
}

console.log(isPalindromeDifficult(input)) // true


