const maize = [
  [1, 1, 1, 0, 0, 1],
  [0, 0, 0, 0, 1, 1],
  [0, 1, 1, 0, 1, 1],
  [0, 0, 0, 0, 1, 1],
  [1, 1, 1, 0, 1, 1],
  [1, 0, 0, 0, 0, 1]
]

function checkPath(start, end, maize) {
  const connectMap = {}
  let result = true

  for(let i = 0; i < maize.length - 1; i++){
    const currentRow = maize[i]
    const nextRow = maize[i + 1]

    for(let j = 0; j < currentRow.length; j++){
      const isRowConnected = currentRow[j] === 0 && nextRow[j] === 0

      if(isRowConnected){
	connectMap[i] = connectMap[i] ? connectMap[i].concat(j) : [j]
      }
    }
  }

  let rowStart = start.y
  for(let i = 0; i < maize.length - 1; i++){
    const currentRow = maize[i]
    let startY = Math.min(rowStart, connectMap[i])
    let finishY = Math.max(rowStart, connectMap[i])

    while (startY <= finishY){
      if(currentRow[startY] !== 0){
	result = false
      }
      startY++
    }

    rowStart = connectMap[i];

    if(!rowStart) {
      result = false;
      break;
    }
  }

  console.log(result)
}

checkPath({x: 0, y: 4}, {x: 5, y: 5}, maize)
