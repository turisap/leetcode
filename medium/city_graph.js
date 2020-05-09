function city_graph(city) {
  var visited = [];
  var count = 0;
  Object.keys(city).forEach((v) => {
    if (!visited.includes(v)) {
      dfs(v, visited, count, city);
      count++;
    }
  });

  return count;
}

function dfs(v, visited, count, city) {
  visited.push(v);
  city[v].forEach((n) => {
    if (!visited.includes(n)) {
      dfs(n, visited, count, city);
    }
  });
}

console.log(city_graph({ p0: ["p1", "p2"], p1: ["p0"], p2: ["p0"], p3: [] }));
