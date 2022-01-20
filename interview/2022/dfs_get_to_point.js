const graph = {
  a: ['b', 'c'],
  b: ['d'],
  d: [],
  c: ['e'],
  e: ['f', 'o', 'g'],
  g: ['k', 'l']
}

const dfs = (graph, v, t) => {
  if(v === t) return true
  if(v.visited) return false

  const queue = [v]
  v.visited = true

  while(queue.length > 0) {
    const node = queue.shift()       // if you replace shift()
				     // with pop you'll get DFS
    node.visited  = true;

    if(graph[node]){
      for(let child of graph[node]){
	console.log(child)
	if(!child.visited){
	  queue.push(child)
	  child.visited = true;
	}

	if(child === t) return true
      }
    }
  }

  return false;
}

console.log(dfs(graph, 'a', 'd'))
