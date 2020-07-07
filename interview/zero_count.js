const zero_count = (threshold) => {
  let str = "";

  for (let i = 0; i <= threshold; i += 10) {
    str += i.toString();
  }

  return str.match(/0/g).length;
};

console.log(zero_count(10));
console.log(zero_count(100));
