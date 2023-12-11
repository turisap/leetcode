const reverse_words = (sentence) =>
  sentence
    .split(" ")
    .map((word) => word.split("").reverse().join(""))
    .join(" ");

console.log(reverse_words("I am travelling down the river"));
