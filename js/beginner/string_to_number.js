const input =
  " _  _  _  _  _  _  _  _  _ \n" +
  "| | _| _|| ||_ |_   ||_||_|\n" +
  "|_||_  _||_| _||_|  ||_| _|\n";

const parseStringToInteger = (bankAccount) => {
  const digits = (bankAccount.length - 3) / 9;
  const stringGroups = Array(digits).fill("");
  let currentRow = 0;

  for (let i = 0; i < bankAccount.length; i++) {
    if (bankAccount[i] === "\n") {
      currentRow++;
      continue;
    }

    const currentGroupIn = Math.floor((i - currentRow * (digits * 3 + 1)) / 3);

    stringGroups[currentGroupIn] += bankAccount[i];
  }

  console.log(stringGroups);

  const patterns = {
    0: " _ | ||_|",
    1: "     |  |",
    2: " _  _||_ ",
    3: " _  _| _|",
    4: "   |_|  |",
    5: " _ |_  _|",
    6: " _ |_ |_|",
    7: " _   |  |",
    8: " _ |_||_|",
    9: " _ |_| _|",
  };

  return parseInt(
    stringGroups
      .map((group) => {
        for (let key in patterns) {
          if (group === patterns[key]) return key;
        }
      })
      .join("")
  );
};

parseStringToInteger(input);
