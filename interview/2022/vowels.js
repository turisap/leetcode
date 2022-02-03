const input1 = 'hello'
const input2 = 'hhh'

const findVowels = string => {
  const matches  = string.match(/[aeiou]/g);
  return matches?.length || 0
}

console.log(findVowels(input1))
console.log(findVowels(input2))
