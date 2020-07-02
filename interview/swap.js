const swap = (a, b) => {
  [a, b] = [b, a];

  return `${a} ${b}`;
};

console.log(swap(1, 5));
