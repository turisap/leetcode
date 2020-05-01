const increment_string = (str) => {
  if (/\D$/.test(str)) return `${str}1`;
  if (!str) return "1";

  let number = str.match(/\d+$/)[0].split("");
  let rest = str.match(/^\D+/);
  const text = rest ? rest[0] : "";

  for (let i = number.length - 1; i > -1; i--) {
    let d = parseInt(number[i]);
    if (d < 9) {
      number[i] = d + 1;
      break;
    } else if (d === 9 && i === 0) {
      number[i] = 0;
      number.unshift(1);
      break;
    }
    number[i] = 0;
  }

  return `${text}${number.join("")}`;
};
