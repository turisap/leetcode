
const graph = {
  a: ['b', 'c'],
  b: ['d'],
  d: ['c'],
  c: ['e'],
  e: ['f', 'o', 'g'],
  g: ['k', 'l']
}

function dfs(graph, v, t) {
	if(v === t) return true
	if(v.visited) return false

	v.visited = true

	if(graph[v]){
	      for(let neighbor of graph[v]) {
		if(!neighbor.visited) {
		    let reached = dfs(graph, neighbor, t)
		    if(reached) return true
	        }
	      }
	}

	return false
}

console.log(dfs(graph, 'a', 'e'))
