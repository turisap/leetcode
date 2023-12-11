const oddSort = (arg) => {
  for (let i = 0; i < arg.length; i++){
    for(let c = 0; c < arg.length; c++){
      let curr = arg[c];
      let next = arg[c + 1];

      if(curr % 2 === 0 || next % 2 === 0) continue

      if(curr > next){
	const temp = curr;
	curr = next;
	next = temp;

	arg[c] = curr
	arg[c + 1] = next
      }
    }
  }

  return arg
}


const arr =  [1, 5,  3, 6, 7, 4, 12, 14, 16, 15, 9]
const arr2 = [...arr]
console.log(arr)
console.log(oddSort(arr2))
