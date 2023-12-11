type SuccessCb = (val: any) => any;

type ErrorCb = (val: any) => any;

class PromisePolyfill<T> {
  successCb: Array<SuccessCb> = [];
  failureCb: ErrorCb | null = null;

  constructor(
    executor: (resolve: (val: T) => void, reject: (err: any) => void) => void
  ) {
    executor(this.resolve, this.reject);
  }

  private resolve = (val: T) => {
    let nextValue = val;

    this.successCb.forEach((cb) => {
      nextValue = cb(nextValue);
    });
  };

  private reject = (error: any) => {
    this.failureCb?.(error);
  };

  public catch = (cb: ErrorCb): PromisePolyfill<T> => {
    this.failureCb = cb;

    return this;
  };

  public then = (cb: SuccessCb): PromisePolyfill<T> => {
    this.successCb.push(cb);

    return this;
  };

  public static resolve(val: any) {
    return new PromisePolyfill((res) => {
      queueMicrotask(() => res(val));
    });
  }

  public static reject(error: any) {
    return new PromisePolyfill((res, rej) => {
      queueMicrotask(() => rej(error));
    });
  }
}

PromisePolyfill.resolve("kiskis").then((result) => {
  console.log(result);
});

PromisePolyfill.reject("oh no").catch(console.log);

const pr = new PromisePolyfill<number>((res, rej) => {
  setTimeout(() => res(45), 500);
});

pr.then((res) => {
  console.log(res);
  return 5;
})
  .then((res) => {
    console.log("second then");
    console.log(res);
  })
  .then((res) => {
    console.log("third res");
    console.log(res);
  })
  .catch((err) => {
    console.log("rejected");
    console.log(err);
  });
