const input = `An old silent pond...
A frog jumps into the pond,
splash! Silence again.`;
//
// const input = `Autumn moonlight -
// a worm digs silently
// into the chestnut`;

const is_haiku = (str) => {
  const syllabilRegex = /[^aeiouy]*[aeiouy]+(?:[^aeiouy]*$|[^aeiouy](?=[^aeiouy]))?/gi;

  const arr = str.split("\n").map((line) => {
    line
      .replace(/-/g, "")
      .split(" ")
      .filter((w) => /\w+/.test(w))
      .map((w) => {
        console.log(w);
        return /[aeiou][bcdfghjklmnpqrstvwxys]+e$/.test(w)
          ? w.match(syllabilRegex).length - 1
          : w.match(syllabilRegex).length;
      })
      .reduce((tail, curr) => (tail += curr), 0);
  });

  return arr;
};

console.log(is_haiku(input));
