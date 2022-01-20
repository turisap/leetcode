const rle = (input) => {
  const res = []
  let stack = input[0]

  const encode = string => `${string[0]}${string.length > 1 ? string.length: ''}`

  for(let c = 1; c <= input.length; c++){
    const char = input[c]
    if(char === stack[0]){
      stack += char;
      continue
    } 

    res.push(encode(stack))
    stack = char
  }

  return res.join('')
}


// these two should be the same
console.log(rle('AVVVBBBVVXDHJFFFFDDDDDDHAAAAJJJDDSLSSSDDD'))
console.log('AV3B3V2XDHJF4D6HA4J3D2SLS3D3')
