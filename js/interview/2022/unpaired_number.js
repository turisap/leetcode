const input = [1, 2, 1, 3, 2, 0, 0]

const unpairedNumber = (arr) => {
  const map = {}

  for(let number of arr){
    if(typeof map[number] !== 'undefined') {
      delete map[number]
    } else {
      map[number] = number
    }
  }

  return Number(JSON.stringify(map).match(/\d/)?.[0])
}

console.log(unpairedNumber(input))
