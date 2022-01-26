const input = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

const shuffleFast = (arr) => {
  const res = [...arr] // THIS IS NOT TO MUTATE THE INPUT

  for(let i = res.length - 1; i >= 0; i--){
    const temp = res[i]
    const rdx = Math.floor(Math.random() * res.length)

    res[i] = res[rdx]
    res[rdx] = temp
  }

  return res
}

let i = 10
while(i){
  console.log(shuffleFast(input).join(''))
  console.log(input)
  console.log('=====================')
  i--
}
