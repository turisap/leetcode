const if_two_passed = (...args) => {
  args.forEach((arg) => {
    if (arg === 2) console.log("yeah");
  });
};

if_two_passed("sdkfj", ",jd", true, 2);
