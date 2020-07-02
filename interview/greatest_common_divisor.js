const gcd = (a, b) => {
  if (a === 0) return b;
  if (b === 0) return a;

  const min = Math.min(a, b);
  const max = Math.max(a, b);

  return gcd(min, max % min);
};

console.log(gcd(270, 192));
