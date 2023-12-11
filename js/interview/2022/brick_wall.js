const input = [[4, 2, 3], [2, 2, 5], [2, 4, 3], [3, 1, 5]]
const input2 = [[2, 3, 2, 4], [5, 3, 3], [2, 5, 4], [4,2 ,4,1], [3, 6, 2]]

const getWall = (wall) => {
  const wallHeight = wall.length;
  const jointsMap = new Map()
  const wallLength = wall[0].reduce((acc, curr) => acc + curr, 0)
  
  for(let row of wall){
    let acc = 0
    for(let brick of row){
      acc += brick
      if(acc === wallLength) continue;

      if(jointsMap.has(acc)){
	jointsMap.set(acc, jointsMap.get(acc) + 1)
      } else {
	jointsMap.set(acc, 1)
      }
    }
  }

  const joints = jointsMap.values()

  return wallHeight - Math.max(...joints)
}

console.log(getWall(input2))
