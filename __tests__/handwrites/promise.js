// const PENDING = 'pending';
// const FULFILLED = 'fulfilled';
// const REJECTED = 'rejected';

// class Promise {
//   constructor(executor) {
//     this.status = PENDING;
//     this.value = null;
//     this.reason = null;
//     this.resolveCallbacks = [];
//     this.rejectCallbacks = [];
//     try {
//       executor(this.resolve.bind(this), this.reject.bind(this));
//     } catch (error) {
//       this.reject(error);
//     }
//   }

//   resolve(value) {
//     if (this.status === PENDING) {
//       this.value = value;
//       this.status = FULFILLED;
//       this.resolveCallbacks.forEach((onFulfilled) => {
//         onFulfilled();
//       });
//     }
//   }

//   reject(reason) {
//     if (this.status === PENDING) {
//       this.reason = reason;
//       this.status = REJECTED;
//     }
//     this.rejectCallbacks.forEach((onRejected) => {
//       onRejected();
//     });
//   }

//   then(onFulfilled, onRejected) {
//     onFulfilled = typeof onFulfilled === 'function' ? onFulfilled : (data) => data;
//     onRejected =
//       typeof onRejected === 'function'
//         ? onRejected
//         : (err) => {
//             throw err;
//           };

//     const controlledPromise = new Promise((resolve, reject) => {
//       if (this.status === PENDING) {
//         this.resolveCallbacks.push(() => {
//           setTimeout(() => {
//             try {
//               const x = onFulfilled(this.value);
//               this.resolvePromise(x, controlledPromise, resolve, reject);
//             } catch (error) {
//               reject(error);
//             }
//           });
//         });
//         this.rejectCallbacks.push(() => {
//           setTimeout(() => {
//             try {
//               const x = onRejected(this.reason);
//               this.resolvePromise(x, controlledPromise, resolve, reject);
//             } catch (error) {
//               reject(error);
//             }
//           });
//         });
//       }

//       if (this.status === FULFILLED) {
//         setTimeout(() => {
//           try {
//             const x = onFulfilled(this.value);
//             this.resolvePromise(x, controlledPromise, resolve, reject);
//           } catch (error) {
//             reject(error);
//           }
//         });
//       }

//       if (this.status === REJECTED) {
//         setTimeout(() => {
//           try {
//             const x = onRejected(this.reason);
//             this.resolvePromise(x, controlledPromise, resolve, reject);
//           } catch (error) {
//             reject(error);
//           }
//         });
//       }
//     });
//     return controlledPromise;
//   }

//   resolvePromise(x, controlledPromise, resolve, reject) {
//     if (x === controlledPromise) {
//       return reject(new TypeError('Chaining cycle detected for promise !'));
//     }
//     if (x && (typeof x === 'object' || typeof x === 'function')) {
//       let called;
//       try {
//         let then = x.then;
//         if (typeof then === 'function') {
//           then.call(
//             x,
//             (value) => {
//               if (called) return;
//               called = true;
//               resolvePromise(value, controlledPromise, resolve, reject);
//             },
//             (reason) => {
//               if (called) return;
//               called = true;
//               reject(reason);
//             }
//           );
//         } else {
//           resolve(x);
//         }
//       } catch (err) {
//         if (called) return;
//         called = true;
//         reject(err);
//       }
//     } else {
//       resolve(x);
//     }
//   }
// }

const states = {
  PENDING: 'pending',
  FULFILLED: 'fulfilled',
  REJECTED: 'rejected'
};

const isThenable = (maybePromise) =>
  maybePromise && typeof maybePromise.then === 'function';

class Promise {
  constructor(computation) {
    this._state = states.PENDING;

    this._value = undefined;
    this._reason = undefined;

    this._thenQueue = [];
    this._finallyQueue = [];

    if (typeof computation === 'function') {
      setTimeout(() => {
        try {
          computation(this._onFulfilled.bind(this), this._onRejected.bind(this));
        } catch (ex) {
          this._onRejected(ex);
        }
      });
    }
  }

  then(fulfilledFn, catchFn) {
    const controlledPromise = new Promise();
    this._thenQueue.push([controlledPromise, fulfilledFn, catchFn]);

    if (this._state === states.FULFILLED) {
      this._propagateFulfilled();
    } else if (this._state === states.REJECTED) {
      this._propagateRejected();
    }

    return controlledPromise;
  }

  catch(catchFn) {
    return this.then(undefined, catchFn);
  }

  finally(sideEffectFn) {
    if (this._state !== states.PENDING) {
      sideEffectFn();

      return this._state === states.FULFILLED
        ? Promise.resolve(this._value)
        : Promise.reject(this._reason);
    }

    const controlledPromise = new Promise();
    this._finallyQueue.push([controlledPromise, sideEffectFn]);

    return controlledPromise;
  }

  _propagateFulfilled() {
    this._thenQueue.forEach(([controlledPromise, fulfilledFn]) => {
      if (typeof fulfilledFn === 'function') {
        const valueOrPromise = fulfilledFn(this._value);

        if (isThenable(valueOrPromise)) {
          valueOrPromise.then(
            (value) => controlledPromise._onFulfilled(value),
            (reason) => controlledPromise._onRejected(reason)
          );
        } else {
          controlledPromise._onFulfilled(valueOrPromise);
        }
      } else {
        return controlledPromise._onFulfilled(this._value);
      }
    });

    this._finallyQueue.forEach(([controlledPromise, sideEffectFn]) => {
      sideEffectFn();
      controlledPromise._onFulfilled(this._value);
    });

    this._thenQueue = [];
    this._finallyQueue = [];
  }

  _propagateRejected() {
    this._thenQueue.forEach(([controlledPromise, _, catchFn]) => {
      if (typeof catchFn === 'function') {
        const valueOrPromise = catchFn(this._reason);

        if (isThenable(valueOrPromise)) {
          valueOrPromise.then(
            (value) => controlledPromise._onFulfilled(value),
            (reason) => controlledPromise._onRejected(reason)
          );
        } else {
          controlledPromise._onFulfilled(valueOrPromise);
        }
      } else {
        return controlledPromise._onRejected(this._reason);
      }
    });

    this._finallyQueue.forEach(([controlledPromise, sideEffectFn]) => {
      sideEffectFn();
      controlledPromise._onRejected(this._value);
    });

    this._thenQueue = [];
    this._finallyQueue = [];
  }

  _onFulfilled(value) {
    if (this._state === states.PENDING) {
      this._state = states.FULFILLED;
      this._value = value;
      this._propagateFulfilled();
    }
  }

  _onRejected(reason) {
    if (this._state === states.PENDING) {
      this._state = states.REJECTED;
      this._reason = reason;
      this._propagateRejected();
    }
  }
}

Promise.resolve = (value) => new Promise((resolve) => resolve(value));
Promise.reject = (value) => new Promise((_, reject) => reject(value));

const fs = require('fs');
const path = require('path');

const readFile = (filename, encoding) =>
  new Promise((resolve, reject) => {
    fs.readFile(filename, encoding, (err, value) => {
      if (err) {
        return reject(err);
      }
      resolve(value);
    });
  });

const delay = (timeInMs, value) =>
  new Promise((resolve) => {
    setTimeout(() => {
      resolve(value);
    }, timeInMs);
  });

// readFile(path.join(__dirname, 'index.js'), 'utf8')
//   .then((text) => {
//     console.log({ text });
//   })
//   .then((newText) => {
//     console.log({ newText });
//   })
//   .catch((err) => {
//     console.log({ err });
//   })
//   .finally(() => {
//     console.log('--------DONE--------');
//   });

const doAsyncStuff = async () => {
  try {
    const text = await readFile(path.join(__dirname, 'indexxxx.js'), 'utf8');
    console.log({ text });

    const newText = await delay(2000, text.replace(/abcde/g, ''));
    console.log({ newText });
  } catch (error) {
    console.log({ error });
  }
  console.log('------DONE-------');
};

doAsyncStuff();

Promise.deferred = function () {
  let deferred = {};
  deferred.promise = new Promise((resolve, reject) => {
    deferred.resolve = resolve;
    deferred.reject = reject;
  });
  return deferred;
};

var promisesAplusTests = require('promises-aplus-tests');

promisesAplusTests(Promise, function (err) {
  console.log('errorororooror', err);
});

module.exports = Promise;
