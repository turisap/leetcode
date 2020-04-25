const array_reverse = (array) => {
  for (let i = 0, j = array.length - 1; i < array.length; i++, j--) {
    if (i >= j) break;

    const temp = array[j];

    array[j] = array[i];
    array[i] = temp;
  }

  return array;
};

