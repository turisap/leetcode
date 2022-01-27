const input1 = 12521;
const input2 = 1221;
const input3 = 1889881;

const isNumberPolyndrome = number => {
  if(number < 0) return false
  if(number < 10) return true
  if(number % 10 === 0) return false


  let reverse = 0

  while (reverse < number) {
    reverse *= 10
    reverse += number % 10

    number = Math.trunc(number / 10)
  }

  return reverse === number || Math.trunc(reverse / 10) === number
}

console.log(isNumberPolyndrome(input3))
