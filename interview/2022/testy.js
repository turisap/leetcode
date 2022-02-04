const summ = (a, b) => {
  // 
  const res = []

  while(a.length || b.length){
    const summ = Math.floor(a + b)

    res.push(summ)
  }
}


const summ = (a, b) => {
  const res = ''

  while(a.length || b.length){
    const summ = Math.floor(a + b)

    // 
    res = `${summ}${res}`
  }
}


const summ = (a, b) => {
  // 
  const res = new Array(Math.max(a.length, b.length))

  while(a.length || b.length){
    const summ = Math.floor(a + b)

    res.push(summ)
  }
}
