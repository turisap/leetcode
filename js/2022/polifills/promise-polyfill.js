// constructable
// resolve / reject
// then / catch

// @TODO next promise.all, race, bind

function PromisePolyfill(executor) {
  executor(this._resolve.bind(this), this._reject.bind(this));
}

PromisePolyfill.prototype.then = function (cb) {
  this.successCb.push(cb);

  return this;
};

PromisePolyfill.prototype.catch = function (cb) {
  this.failureCb = cb;

  return this;
};

PromisePolyfill.prototype._resolve = function (res) {
  let processingValue = res;

  this.successCb.forEach((cb) => {
    processingValue = cb(processingValue);
  });
};

PromisePolyfill.prototype._reject = function (error) {
  this.failureCb(error);
};

const pr = new PromisePolyfill((res, rej) => {
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
