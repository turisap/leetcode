const flat = (arr) => {
  const res = [];

  for(let element of arr){
    if(!Array.isArray(element)){
      res.push(element)
    } else {
      res.push(...flat(element))
    }
  }

  return res
}

console.log(flat([1,2,3, [4, 5, [6, 7]]]))
