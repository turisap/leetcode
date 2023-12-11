const input = {
  a: ['b', 'c'],
  b: ['d'],
  d: [],
  c: ['e'],
  e: ['f', 'o', 'g'],
  g: ['k', 'l']
}


const dfs = (graph, start, target, visited = new Set()) => {
  console.log(start)

  if(start === target) return true
  if(visited.has(target)) return false

  visited.add(start)

    if(graph[start]){
	  for(let neighbor of graph[start]) {
	    if(!neighbor.visited) {
		let reached = dfs(graph, neighbor, target, visited)
		if(reached) return true
	    }
	  }
    }

    return false
}

console.log(dfs(input, 'a', 'z'))
