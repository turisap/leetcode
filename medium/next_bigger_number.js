const next_bigger_number = (number) => {
  const digits = number.toString().split("");
  let nextBig = Number.MAX_SAFE_INTEGER;

  // idea is to loop over each one and compare with any other one
  for (let i = digits.length - 1; i >= 0; i--) {
    for (let j = digits.length - 1; j >= 0; j--) {
      const copy = [...digits];
      const current = parseInt(digits[i]);
      const next = parseInt(digits[j]);
      const temp = current;

      copy[i] = next;
      copy[j] = temp;

      const n = parseInt(copy.join(""));

      if (n < nextBig && n > number) {
        nextBig = n;
      }
    }
  }

  return nextBig;
};

// console.log(next_bigger_number(144));
// console.log(next_bigger_number(123456));
// console.log(next_bigger_number(67980));
console.log(next_bigger_number(116398)); // expect 116839
