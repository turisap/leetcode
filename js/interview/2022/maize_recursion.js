const input = [
  [1, 1, 1, 0, 0, 1],
  [0, 0, 0, 0, 1, 1],
  [0, 1, 1, 0, 1, 1],
  [0, 0, 0, 0, 1, 1],
  [1, 1, 1, 0, 1, 1],
  [1, 0, 0, 0, 0, 1]
]

const checkPath = (maize, start, finish) => {
  maize[start.y][start.x] = 5

  const siblings = getValidSiblings(maize, start)

  if(siblings.length > 0){
    for(let sibling of siblings){
      const isReached = sibling.x === finish.x && sibling.y === finish.y
      const isVisited = maize[sibling.y][sibling.x] === 5

      if(isReached){
	return true
      } else if(!isVisited){
	return checkPath(maize, sibling, finish)
      }
  }

  return false;
}


function getValidSiblings  (maize, node)  {
  const {x, y} = node
  const siblings = []

  if(typeof maize[y - 1] !== 'undefined') {
    siblings.push({x, y: y - 1, value: maize[y - 1][x]})
  }

  if(typeof maize[y + 1] !== 'undefined') {
    siblings.push({x : x, y: y + 1, value: maize[y + 1][x]})
  }

  if(typeof maize[y][x + 1] !== 'undefined') {
    siblings.push({x : x + 1, y: y, value: maize[y][x + 1]})
  }

  if(typeof maize[y][x - 1] !== 'undefined') {
    siblings.push({x : x - 1, y: y, value: maize[y][x - 1]})
  }

  return siblings.filter(el => el.value === 0)
}
}

console.log(checkPath(input, {y: 0, x: 4}, {y: 5, x: 4}))
console.log(input)
