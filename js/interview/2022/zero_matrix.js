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

const input2 = [
  [1, 1, 3],
  [0, 5, 6],
  [3, 5, 9]
]

const fillMatrix = matrix => {
  const rowMap = {}
  const columnMap = {}

  for(let R = 0; R < matrix.length; R++) {
    const row = matrix[R]

    for (let C = 0; C < row.length; C++){ 
      if(row[C] === 0) {
	rowMap[R] = true
	columnMap[C] = true
      }
    }
  }

  for(let R = 0; R < matrix.length; R++) {
    const row = matrix[R]

    for (let C = 0; C < row.length; C++){ 
      if(columnMap[C] || rowMap[R]) {
	matrix[R][C] = 0
      }
    }
  }

  return matrix
}

// console.log(fillMatrix(input))
console.log(fillMatrix(input2))
