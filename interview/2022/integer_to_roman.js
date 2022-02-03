const intToRoman = (n) => {
  const numMap =  {
    1: 'I',
    4: 'IV',
    5: 'V',
    9: 'IX',
    10: 'X',
    40: 'XL',
    50: 'L',
    90: 'XC',
    100: 'C',
    400: 'CD',
    500: 'D',
    900: 'CM',
    1000: 'M'
  }

  const integers = Object.keys(numMap).sort((a, b) => b - a)
  let res = ''

  for (let i = 0; i < integers.length; i++){
    const current = integers[i]

    if(n < current){
      continue
    } else {
      const times = Math.floor(n / current)
      
      res += numMap[current].repeat(times)
      n = n % current
    }
  }

  return res
}

console.log(intToRoman(1994))

