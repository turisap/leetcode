const grid2 = [
  [1, 1, 0, 0, 0],
  [0, 1, 0, 0, 0],
  [1, 0, 0, 1, 1],
  [0, 0, 0, 1, 1]
]

const getIslands = (grid) => {
  let result = 0

  if(grid.length === 0) return 0

  for(let i = 0; i < grid.length; i++) {
    const row = grid[i]

    for(let c = 0; c < row.length; c++){
      const isLeftLand = row[c - 1] === 1;
      const isUpLand =   grid[i - 1]?.[c] === 1
      const isLand = row[c] === 1

      if(isLand && !isLeftLand && !isUpLand){
	result++
      }
    }
  }

  return result
}

console.log(getIslands(grid2))

