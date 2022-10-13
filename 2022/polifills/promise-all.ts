// @TODO next all settled
type SuccessCB<T> = (val: T) => void;
type FailureCB = (err: any) => void;

class PromiseAllPolyfill<T extends unknown[] | []> {
  private resolveCounter = 0;
  private successCB: Array<SuccessCB<T>> = [];
  private failureCB: Array<FailureCB> = [];
  private result: T;
  private threshold: number = 0;

  // @TODO non-promise values
  constructor(values: Array<any>) {
    this.threshold = values.length;
    this.result = [] as T;

    this._init(values);

    return this;
  }

  private _checkCounter() {
    this.resolveCounter++;
    if (this.resolveCounter === this.threshold) {
      this.successCB.forEach((cb) => cb(this.result));
    }
  }

  private _rejectAll(err: any) {
    if (this.failureCB.length) {
      this.failureCB.forEach((cb) => cb(err));

      return;
    }

    throw new Error("Unhandled rejection in promise");
  }

  private _init(values: any[]) {
    values.forEach((value, idx) => {
      if (typeof value.then !== "undefined") {
        value
          .then((res: any) => {
            this.result[idx] = res;
            this._checkCounter();
          })
          .catch((e: any) => {
            this._rejectAll(e);
          });
      } else {
        this.result[idx] = value;
        this.resolveCounter++;
      }
    });
  }

  then(cb: SuccessCB<T>) {
    this.successCB.push(cb);

    return this;
  }

  catch(cb: FailureCB) {
    this.failureCB.push(cb);

    return this;
  }
}

const pr1 = new Promise((res) => {
  setTimeout(res, 3000, "trololo");
});

const pr2 = new Promise((res) => {
  setTimeout(res, 1000, "kiskis");
});

const pr3 = new Promise((res, rej) => {
  setTimeout(res, 1500, "all good");
});

const res = new PromiseAllPolyfill<string[]>([pr1, pr2, 5, pr3])
  .then(console.log)
  .catch((err) => {
    console.log(err);
  });
