const rotateImage = (matrix) => {
  console.log(matrix.reduce((acc, curr) => acc + curr.join(" ") + "\n", ""));

  // transpose
  const n = matrix.length;
  for (let i = 0; i < n; i++) {
    for (let j = i + 1; j < n; j++) {
      const tmp = matrix[j][i];
      matrix[j][i] = matrix[i][j];
      matrix[i][j] = tmp;
    }
  }

  // reflect
  for (let i = 0; i < n; i++) {
    for (let j = 0; j < n / 2; j++) {
      const tmp = matrix[i][j];
      matrix[i][j] = matrix[i][n - j - 1];
      matrix[i][n - j - 1] = tmp;
    }
  }

  console.log(matrix.reduce((acc, curr) => acc + curr.join(" ") + "\n", ""));
  return matrix;
};

module.exports = rotateImage;
