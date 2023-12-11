const rgb_2_hex = (...args) => {
  // exclude illigal inputs
  args = args.map((color) => {
    if (color < 0) return 0;

    if (color > 255) return 255;

    return color;
  });

  // 0, 0, 0 input edge case
  if (args.join("") === "000") return "000000";

  return args.reduce((tail, current) => {
    let hex = current.toString(16).toUpperCase();

    hex = hex.length === 1 ? "0" + hex : hex;

    return tail + hex;
  }, "");
};

rgb_2_hex(0, 155, 216);

module.exports = rgb_2_hex;
