const interval3 = [[11,15], [3,5], [8,9], [1, 4], [9, 10]]

const merge = (args) => {
  if(args.length < 2) {
    return args
  }

  const sorted = args.sort((a, b) => a[0] - b[0])
  const [firstStart, firstEnd] = sorted[0]
  let currentStart = firstStart;
  let currentEnd = firstEnd;
  const result = []

  for(let c = 1; c < sorted.length; c++){
    const [start, end] = sorted[c]

    if(start <= currentEnd){
	currentEnd = end;
	continue;
    }

    if(start > currentStart){
      result.push([currentStart, currentEnd]);

      currentStart = start;
      currentEnd = end;
    }
    
  }

  result.push([currentStart, currentEnd])

  return result
}

console.log(merge(interval3)) // [ [ 1, 5 ], [ 8, 10 ], [ 11, 15 ] ]
