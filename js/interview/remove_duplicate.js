const remove_duplicate = (arr) => {
  const unique = new Set();

  arr.forEach((el) => unique.add(el));

  return Array.from(unique.values());
};

remove_duplicate([1, 1, 2, 5, 2, 7, 4, 458, 34, 3]);
