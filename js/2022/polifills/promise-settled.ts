type SuccessCB<T> = (val: T) => void;

class PromiseAllSettled {
  private result: Array<any> = [];
  private successCB: Array<SuccessCB<any>> = [];
  private resolveCounter = 0;
  private lengthInitial = 0;

  constructor(values: Array<any>) {
    this.lengthInitial = values.length;

    values.forEach((value, idx) => {
      const isThenable =
        value !== null &&
        typeof value === "object" &&
        typeof value.then === "function";

      if (isThenable) {
        value
          .then((res: any) => {
            this.result[idx] = res;
          })
          .catch((err: any) => {
            this.result[idx] = err;
          })
          .finally(this._checkResolveStatus);
      } else {
        this.result[idx] = value;
        this._checkResolveStatus();
      }
    });
  }

  then(cb: SuccessCB<any>) {
    this.successCB.push(cb);
  }

  private _checkResolveStatus = () => {
    this.resolveCounter++;

    if (this.lengthInitial === this.resolveCounter) {
      this.successCB.forEach((cb) => cb(this.result));
    }
  };
}

const pr0 = new Promise((res) => {
  setTimeout(res, 1500, "very  good");
});

const resG = new Promise((res) => {
  res("good");
});

const rejN = new Promise((res, rej) => {
  setTimeout(rej, 2000, "oh no");
});

const prA = new PromiseAllSettled([pr0, resG, rejN, null]);

prA.then(console.log);
