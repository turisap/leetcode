const input = '1 2 + 4 * 3 +'

const polist_notation = (arg) => {
  const operandsStack = [];
  
  for(let char of arg){
    if(char === ' '){
      continue;
    }

    const entry = parseInt(char, 10)
    if(!isNaN(entry)) {
      operandsStack.push(parseInt(entry))
    } else {
      const result = eval(`${operandsStack.pop()} ${char} ${operandsStack.pop()}`)

      operandsStack.push(result)
    }
  }

  return operandsStack
}

console.log(polist_notation(input))
