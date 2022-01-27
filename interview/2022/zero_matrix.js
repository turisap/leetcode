const input = [
  [1, 1, 1],
  [1, 0, 1],
  [1, 1, 1]
]

const input1 = [
  [0, 1, 2, 0],
  [3, 4, 5, 2],
  [1, 3, 1, 5]
]

const fillMatrix = matrix => {
  for(let R = 0; R < matrix.length; R++) {
    const row = matrix[R]

    for (let C = 0; C < row.length; C++){
      if(row[C] === 'X') {
	row[C] = 0;
	continue;
      }

      if(row[C] === 0){
	const prevRow = R - 1 >= 0 ? R - 1 : null
	const nextRow = R + 1 < matrix.length ? R + 1 : null
	const prevItem = C - 1 >= 0 ? C - 1 : null
	const nextItem = C + 1 < row.length ? C + 1 : null

	if(prevRow !== null) matrix[prevRow][C] = 0;
	if(nextRow !== null) matrix[nextRow][C] = 'X';
	if(prevItem !== null) matrix[R][prevItem] = 0;
	if(nextItem !== null) matrix[R][nextItem] = 'X';
      }
    }
  }

  return matrix
}

// console.log(fillMatrix(input))
console.log(fillMatrix(input1))
