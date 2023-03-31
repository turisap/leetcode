var MinStack = function () {
  this._stack = [];
  this._minTracker = [];
};

/**
 * @param {number} val
 * @return {void}
 */
MinStack.prototype.push = function (val) {
  this._stack.push(val);

  if (!this._minTracker.length) {
    this._minTracker.push(val);
    return;
  }

  val <= this._minTracker.at(-1) ? this._minTracker.push(val) : null;
};

/**
 * @return {void}
 */
MinStack.prototype.pop = function () {
  const removed = this._stack.pop();

  if (this._minTracker.at(-1) === removed) this._minTracker.pop();
};

/**
 * @return {number}
 */
MinStack.prototype.top = function () {
  return this._stack.at(-1);
};

/**
 * @return {number}
 */
MinStack.prototype.getMin = function () {
  return this._minTracker.at(-1);
};

module.exports = MinStack;
