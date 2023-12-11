const input = [-1, 0, 1, 2, -1, -4, -1];

const sumOfThree = (numbers, target) => {
  const result = []
  const sorted = numbers.sort()

  if(numbers.length < 3) return result


  for(let i = 0; i < sorted.length; i++){

    if(sorted[i] > target){
      break;
    }

    if(sorted[i] === sorted[i - 1]){
      continue
    }
    
    let j = i + 1;
    let k = sorted.length - 1;

    while(j < k){
      const summ = sorted[i] + sorted[j] + sorted[k];

      if(summ === target) {
	result.push([sorted[i], sorted[j], sorted[k]])

	while(sorted[j] === sorted[j + 1])j++
	while(sorted[k] === sorted[k - 1])k--

	j++
	k--

	continue
      }

      if(summ < target){
	j++
	continue
      }

      if(summ > target){
	k--
	continue
      }
    }
  }

  return result
}

console.log(sumOfThree(input, 0))
