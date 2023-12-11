const numbers = [1, 5, 2, 6, 7, 4, 3, 0, 11, 15, 8]


const quickSort  = (arr) => {
  if(arr.length < 2) return arr

  const left = []
  const right = []
  const random = Math.floor(Math.random() * arr.length)
  const pivot = arr[random]

  for(let entry of arr) {
    if(entry < pivot) {
      left.push(entry)
    } else if (entry > pivot) {
      right.push(entry)
    }
  }

   return quickSort(left).concat(pivot, quickSort(right))
}
