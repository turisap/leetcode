const anagram = (a, b) => {
  if(!a.length || !b.length){
    return false
  }

  if(a.length !== b.length){
    return false
  }

  let res = true;
  const charsMap = {}
  a = a.toLowerCase()
  b = b.toLowerCase()

  for(let char of a){
    if(charsMap[char]) {
      charsMap[char]++
    } else {
      charsMap[char] = 1
    }
  }

  for(let char of b){
    if(!charsMap[char] || charsMap[char] < 1){
      res = false;
      break;
    } else if(charsMap[char]){
      charsMap[char]--
    }
  }

  return res
}


console.log( anagram('friendd', 'fineder'))
console.log(anagram('b', 'B'))
