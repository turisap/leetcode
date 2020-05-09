const next_bigger_number = (number) => {
  const digits = number.toString().split("");
  const lastIn = digits.length - 1;

  const desc = digits.every((d, i, arr) => {
    return i === 0 || d <= arr[i - 1];
  });

  if (desc) return -1;

  const asc = digits.every((d, i, arr) => {
    return i === 0 || d > arr[i - 1];
  });

  const swap = (aI, bI, arr) => {
    return arr;
  };

  if (asc) {
    let cp = [...digits];
    swap(cp.length - 1, cp.length - 2, cp);

    return parseInt(cp.join(""));
  }

  let stopIn = lastIn;

  for (let i = lastIn - 1; i >= 0; i--) {
    if (digits[i] < digits[i + 1]) {
      stopIn = i;
      break;
    }
  }

  const left = digits.slice(0, stopIn);
  const right = digits.slice(stopIn);
  let smallIdx = 0;
  let smallest = 9;

  for (let c = 1; c < right.length; c++) {
    if (right[c] > right[0] && right[c] < smallest) {
      smallIdx = c;
      smallest = right[c];
    }
  }

  swap(0, smallIdx, right);

  const sortedPart = right.slice(1).sort();

  const res = left.concat(right[0]).concat(sortedPart);

  return parseInt(res.join(""));
};

// console.log(next_bigger_number(32494));
// console.log(next_bigger_number(144));

module.exports = next_bigger_number;
