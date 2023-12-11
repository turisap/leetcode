const rotateArraySlow = (arr, k) => {
  while (k--) {
    arr.unshift(arr.pop());
  }

  return arr;
};

const rotateArrayEasy = (arr, k) => {
  const temp = [];

  while (k--) {
    if (arr.length) {
      temp.unshift(arr.pop());
    } else {
      temp.unshift(temp.pop());
    }
  }

  arr.unshift(...temp);

  return arr;
};

module.exports = rotateArrayEasy;

// 1 2 3 4 5
// 5 4 1 2 3
