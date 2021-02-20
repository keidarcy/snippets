const PENDING = 'pending';
const FULFILLED = 'fulfilled';
const REJECTED = 'rejected';

class Promise {
  constructor(executor) {
    this.status = PENDING;
    this.value = null;
    this.reason = null;
    this.resolveCallbacks = [];
    this.rejectCallbacks = [];
    try {
      executor(this.resolve.bind(this), this.reject.bind(this));
    } catch (error) {
      this.reject(error);
    }
  }

  resolve(value) {
    if (this.status === PENDING) {
      this.value = value;
      this.status = FULFILLED;
      this.resolveCallbacks.forEach((onFulfilled) => {
        onFulfilled();
      });
    }
  }

  reject(reason) {
    if (this.status === PENDING) {
      this.reason = reason;
      this.status = REJECTED;
    }
    this.rejectCallbacks.forEach((onRejected) => {
      onRejected();
    });
  }

  then(onFulfilled, onRejected) {
    onFulfilled = typeof onFulfilled === 'function' ? onFulfilled : (data) => data;
    onRejected =
      typeof onRejected === 'function'
        ? onRejected
        : (err) => {
            throw err;
          };

    const controlledPromise = new Promise((resolve, reject) => {
      if (this.status === PENDING) {
        this.resolveCallbacks.push(() => {
          setTimeout(() => {
            try {
              const x = onFulfilled(this.value);
              this.resolvePromise(x, controlledPromise, resolve, reject);
            } catch (error) {
              reject(error);
            }
          });
        });
        this.rejectCallbacks.push(() => {
          setTimeout(() => {
            try {
              const x = onRejected(this.reason);
              this.resolvePromise(x, controlledPromise, resolve, reject);
            } catch (error) {
              reject(error);
            }
          });
        });
      }

      if (this.status === FULFILLED) {
        setTimeout(() => {
          try {
            const x = onFulfilled(this.value);
            this.resolvePromise(x, controlledPromise, resolve, reject);
          } catch (error) {
            reject(error);
          }
        });
      }

      if (this.status === REJECTED) {
        setTimeout(() => {
          try {
            const x = onRejected(this.reason);
            this.resolvePromise(x, controlledPromise, resolve, reject);
          } catch (error) {
            reject(error);
          }
        });
      }
    });
    return controlledPromise;
  }

  resolvePromise(x, controlledPromise, resolve, reject) {
    if (x === controlledPromise) {
      return reject(new TypeError('Chaining cycle detected for promise !'));
    }
    if (x && (typeof x === 'object' || typeof x === 'function')) {
      let called;
      try {
        let then = x.then;
        if (typeof then === 'function') {
          then.call(
            x,
            (value) => {
              if (called) return;
              called = true;
              resolvePromise(value, controlledPromise, resolve, reject);
            },
            (reason) => {
              if (called) return;
              called = true;
              reject(reason);
            }
          );
        } else {
          resolve(x);
        }
      } catch (err) {
        if (called) return;
        called = true;
        reject(err);
      }
    } else {
      resolve(x);
    }
  }
}

// Promise.deferred = function () {
//   let deferred = {};
//   deferred.promise = new Promise((resolve, reject) => {
//     deferred.resolve = resolve;
//     deferred.reject = reject;
//   });
//   return deferred;
// };

// var promisesAplusTests = require('promises-aplus-tests');

// promisesAplusTests(Promise, function (err) {
//   console.log('errorororooror');
// });

module.exports = Promise;
