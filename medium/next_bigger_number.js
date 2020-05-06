const next_bigger_number = (number) => {
  const digits = number.toString().split("");
  let inaAsc = true;
  let possiblySorted = parseInt([...digits].sort().reverse().join(""));

  if (possiblySorted === number) return -1;

  if (digits.length === 2) {
    return parseInt(digits.reverse().join(""));
  }
  // for (let c = 0; c < digits.length - 2; c++) {
  //   if (parseInt(digits[c]) >= digits[c + 1]) {
  //     inaAsc = false;
  //   }
  // }

  // if (inaAsc) {
  //   const lastI = digits.length - 1;
  //   const temp = digits[lastI - 1];
  //   digits[lastI - 1] = digits[lastI];
  //   digits[lastI] = temp;

  //   return parseInt(digits.join(""));
  // }

  let nextBig;
  let stopIndex = 0;

  for (let i = digits.length - 2; i >= 0; i--) {
    if (digits[i] < digits[i + 1]) {
      stopIndex = i;
      break;
    }
  }

  let smallestIndex = digits.length - 1;
  let smallestToRight = digits[smallestIndex];

  for (let j = digits.length - 1; j > stopIndex; j--) {
    const current = parseInt(digits[j]);

    if (current < smallestToRight) {
      smallestToRight = current;
      smallestIndex = j;
    }
  }

  const temp = digits[smallestIndex];
  digits[smallestIndex] = digits[stopIndex];
  digits[stopIndex] = temp;

  const leftArr = digits.slice(0, stopIndex);
  const rightArr = digits.slice(stopIndex).sort();

  const res = parseInt(leftArr.concat(rightArr).join(""));

  return res;
};

console.log(next_bigger_number(12));
