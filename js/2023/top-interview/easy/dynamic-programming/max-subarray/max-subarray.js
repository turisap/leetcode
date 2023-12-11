const maxSubArraySlow = (arr) => {
  if (arr.length === 1) return arr[0];

  let maxSubArray = -Infinity;

  for (let i = 0; i < arr.length; i++) {
    let currSubArray = 0;

    for (let j = i; j < arr.length; j++) {
      currSubArray += arr[j];

      maxSubArray = Math.max(currSubArray, maxSubArray);
    }
  }

  return maxSubArray;
};

const maxSubArray = (arr) => {
  let maxSubarray = -Infinity;
  let currSum = 0;

  for (let i = 0; i < arr.length; i++) {
    currSum += arr[i];

    maxSubarray = Math.max(maxSubarray, currSum);

    if (currSum < 0) {
      currSum = 0;
    }
  }

  return maxSubarray;
};

module.exports = maxSubArray;
