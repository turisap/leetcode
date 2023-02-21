const majorityElementNaive = (arr) => {
  const numberMap = {};
  let max = 1;
  let major = arr[0];

  for (let num of arr) {
    numberMap[num] ? ++numberMap[num] : (numberMap[num] = 1);

    if (numberMap[num] > max) {
      max = numberMap[num];
      major = num;
    }
  }

  return major;
};

const majorityElementSorting = (arr) => {
  // sort an array
  arr.sort();

  // the majority element will be in the middle anyway
  return arr[Math.floor(arr.length / 2)];
};

const boyerMore = (arr) => {
  let count = 1;
  let candidate = arr[0];

  for (let i = 1; i < arr.length; i++) {
    arr[i] === candidate ? count++ : count--;

    if (count === 0) {
      candidate = arr[i];
      count = 1;
    }
  }

  return candidate;
};

module.exports = boyerMore;
