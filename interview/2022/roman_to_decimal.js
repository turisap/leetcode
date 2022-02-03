const roman2Decimal = (s) => {
    s = s.toUpperCase()

    const numMap = {
	I: 1,
	V: 5,
	X: 10,
	L: 50,
	C: 100,
	D: 500,
	M: 1000
    }
    
    const doublesMap = {
	IV: 4,
	IX: 9,
	XL: 40,
	XC: 90,
	CD: 400,
	CM: 900
    }

    
    let i = 0;
    let counter = 0;
    
    while(i < s.length){
      const curr = s[i]
      const next = s[i + 1]
      if(doublesMap[curr + next]) {
	counter += doublesMap[curr + next]
	i += 2
      } else {
	counter += numMap[curr]
	i++
      }
    }

  return counter
}


console.log(roman2Decimal(''))
console.log(roman2Decimal('III'))
console.log(roman2Decimal('IV'))
console.log(roman2Decimal('VI'))
console.log(roman2Decimal('XIV'))
console.log(roman2Decimal('IX'))
console.log(roman2Decimal('L'))
console.log(roman2Decimal('XL'))
console.log(roman2Decimal('CM'))
