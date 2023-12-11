const myAtoi = string => {
  const trimed = string.trimStart()
  const negative = /^-/.test(trimed)
  const validNumbers = trimed
    .replace(/^[+,-]/, '')
    .replace(/^0*/, '')
    .match(/^\d*/)[0]

  if(!validNumbers) return 0

  // const int  = parseInt(validNumbers, 10)
  
  let int = 0;

  for (let n of validNumbers) {
    int += n
  }

  int = Number(int)
  
  if(negative){
    return Math.max(-Math.pow(2, 31), -int)
  } 

  return Math.min((Math.pow(2, 31) - 1), int)
}

console.log(myAtoi('  -00123'))
console.log(myAtoi('  -00123  dfjdk '))
console.log(myAtoi('  -009999999999999999999999999999999999999999999  dfjdk '))
console.log(myAtoi('words and 987'))
console.log(myAtoi('00000-42a1234'))
console.log(myAtoi('123-'))


