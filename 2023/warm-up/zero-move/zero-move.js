const zeroMoveSort = (arr) => {
  return arr.sort((a, b) => {
    if (b === 0) return -1;

    return 0;
  });
};

const zeroMoveLoopAddMemory = (arr) => {
  let zeroC = 0;
  const answer = [];

  for (let num of arr) {
    if (num === 0) zeroC++;
    if (num !== 0) answer.push(num);
  }

  while (zeroC--) {
    answer.push(0);
  }

  return answer;
};

const zeroMoveTwoPointers = (arr) => {
  let nonZeroPointer = 0;

  for (let i = 0; i < arr.length; i++) {
    if (arr[i] !== 0) {
      arr[nonZeroPointer] = arr[i];
      nonZeroPointer++;
    }
  }

  for (let p = nonZeroPointer; p < arr.length; p++) {
    arr[p] = 0;
  }

  return arr;
};

const zeroMoveOptimal = (arr) => {
  for (let lastNonZero = 0, i = 0; i < arr.length; i++) {
    if (arr[i] !== 0) {
      [arr[i], arr[lastNonZero]] = [arr[lastNonZero], arr[i]];

      lastNonZero++;
    }
  }

  return arr;
};

module.exports = {
  zeroMoveSort,
  zeroMoveLoopAddMemory,
  zeroMoveTwoPointers,
  zeroMoveOptimal,
};
