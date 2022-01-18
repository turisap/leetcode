const inputs = [
  '({[]})',
  '(){[]}()',
  '(){}[',
  '({[})'
]

const map = {
  '{': '}',
  '[': ']',
  '(': ')'
}

const validator = (brackets) => {
  if(brackets.length % 2 === 1){
    return false
  }

  const stack = []
  let result = true

  for(let char of brackets){
    if(char === '{' || char === '[' || char === '(') {
	stack.push(char)
    } else {
      const openBracket = stack.pop();
      if(char !== map[openBracket]){
	result = false
	break;
      }
    }
  }

  return result
}

inputs.forEach(input => console.log(validator(input)))
