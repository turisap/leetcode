const bitsNumber = (number) => {
  return number.toString(2).replace(/0/g, "").length;
};

console.log(bitsNumber(1234));
