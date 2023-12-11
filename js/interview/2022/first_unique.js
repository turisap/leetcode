const word1 = 'leetcode'
const word2 = 'mememe'
const word3 = 'neraarnek'

const firstUnique = (text) => {
  const map = new Map()
  let res = -1;

  for(let char of text){
    if(map.has(char)) {
      map.set(char, map.get(char) + 1)
    } else {
      map.set(char, 1)
    }
  }

  for(let i = 0; i < text.length; i++) {
    if(map.get(text[i]) === 1){
      return i
    }
  }

  return -1
}

console.log(firstUnique(word3))
