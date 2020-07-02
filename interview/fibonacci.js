const fibonacci = (n) => {
  if (n === 1) return 0;
  if (n === 2) return 1;

  return fibonacci(n - 1) + fibonacci(n - 2);
};

const fibonacci_linear = (n) => {
  const arr = [0, 1];

  for (let i = 2; i < n; i++) {
    arr.push(arr[i - 1] + arr[i - 2]);
  }

  return arr[arr.length - 1];
};

const stress_tester = (n) => {
  for (let c = 3; c < n; c++) {
    const linear = fibonacci_linear(c);
    const recursive = fibonacci(c);

    if (linear !== recursive) {
      console.log(`ERROR!!!! Recursive: ${recursive}, linear: ${linear}`);
      break;
    }

    console.log(`Iteration at ${c}, all good`);
  }
};

stress_tester(42);
