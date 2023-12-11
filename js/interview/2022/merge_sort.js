const numbers = [0, 4, 2, 8, 9, 10, 3, 1,  6, 5, 7]

const merge = (left, right) => {
  const merged = [];

  while(left.length && right.length) {
    if(left[0] < right[0]){
      merged.push(left.shift())
    } else {
      merged.push(right.shift())
    }
  }

  return [...merged, ...left, ...right]
}

const mergeSort = (array) => {
  if(array.length < 2) {
    return array
  }

  const mid = Math.floor(array.length / 2)
  const left = array.splice(0, mid)

  return merge(mergeSort(left), mergeSort(array))
}

console.log(mergeSort(numbers))
