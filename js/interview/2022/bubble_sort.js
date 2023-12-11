const input = [2, 7, 5, 3, 5, 8, 0, 11];

const bubbleSort = (arr) => {
  for (let i = 0; i < arr.length; i++) {
    for (let j = 0; j < arr.length - i; j++) {
      let current = arr[j]
      let next = arr[j + 1]

        if (current > next) {
	  [current, next] = [next, current]
	  arr[j] = current;
	  arr[j + 1] = next
        }
      
    }
  }
    return arr
}

console.log(bubbleSort(input))
