const remove_all_duplicates = (string) => {
  const unique = new Set();
  let res = "";

  for (let char of string) {
    if (!unique.has(char)) res += char;

    unique.add(char);
  }

  return res;
};

remove_all_duplicates("heey me i love that lass");
