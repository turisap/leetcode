const next_bigger_number = (number) => {
  const digits = number.toString().split("");
  let swapped = false;

  for (let i = digits.length - 1; i >= 0; i--) {
    for (let j = i - 1; j >= 0; j--) {
      const current = parseInt(digits[i]);
      const next = parseInt(digits[j]);

      console.log(i, j);
      if (current > next) {
        const temp = current;

        digits[i] = next;
        digits[j] = temp;
        swapped = true;
        break;
      }
    }

    if (swapped) break;
  }

  return swapped ? parseInt(digits.join("")) : -1;
};

console.log(next_bigger_number(144));
// console.log(next_bigger_number(123456));
