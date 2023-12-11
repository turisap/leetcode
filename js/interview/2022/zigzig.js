const zigziag = (str, rows) => {
  const matrix = []
  
  if(rows < 2) return str

  for(let i = 0; i < rows; i++){
    const row = new Array(Math.ceil(str.length)).fill(' ')
    matrix.push(row)
  }

  let row = 0;
  let column = 0;

  for(let j = 0; j < str.length; j++){
    const isDiagonal = Math.floor(j / (rows - 1)) % 2 !== 0 

    matrix[row][column] = str[j]
    if(isDiagonal){
      row--
      column++
    } else {
      row++
    }
    // console.log(isDiagonal, j, row, column)
  }

  const res = matrix.map(r => r.join('')).join('').replace(/\s/g, '')

  return res
}

// console.log( zigziag('PAYPALISHIRING', 3) === 'PAHNAPLSIIGYIR')
// console.log( zigziag('PAYPALISHIRING', 4) === 'PINALSIGYAHRPI')
console.log( zigziag('AB', 1))
