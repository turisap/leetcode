const missing_number = (arr) => {
  const expectedSum = ((arr.length + 1) * (arr[0] + arr[arr.length - 1])) / 2;

  const sum = arr.reduce((tail, current) => (tail += current), 0);

  return expectedSum - sum;
};

console.log(missing_number([1, 2, 3, 4, 5, 7, 8, 9, 10]));
console.log(missing_number([13, 14, 15, 16, 18, 19]));
