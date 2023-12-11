const resort = (str) => {
  if (!str) return [];

  return str
    .split("")
    .reduce((tail, current, index) => {
      return (tail += index % 2 === 1 ? `${current},` : current);
    }, "")
    .replace(/,$/, "")
    .split(",")
    .map((couple) => (couple.length === 2 ? couple : couple + "_"));
};

// this is the best solution from codewars.com
function solution(str) {
  return (str.length % 2 ? str + "_" : str).match(/../g);
}

