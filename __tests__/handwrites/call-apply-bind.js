const nameObj = {
  name: 'joe'
};

const getNameArrow = () => {
  console.log(this.name);
};

function getName(a, b, c, d) {
  console.log(this.name);
  console.log({ a, b, c, d });
}

// getNameArrow.call(nameObj, undefined);
// getName.call(nameObj, undefined);
// getName.apply(nameObj, []);
// const binded = getName.bind(nameObj);
// binded();

// Function.prototype.myCall = function (obj) {
//   const newObj = obj || global;
//   newObj.p = this;
//   let newArguments = [];
//   // console.log(arguments);
//   for (let i = 1; i < arguments.length; i++) {
//     newArguments.push('arguments[' + i + ']');
//   }
//   const result = eval('newObj.p(' + newArguments + ')');
//   delete newObj.p;
//   return result;
// };
Function.prototype.myCall = function (obj, ...arr) {
  const newObj = obj || global;
  newObj.p = this;
  const result = newObj.p(...arr);
  delete newObj.p;
  return result;
};

// Function.prototype.myApply = function (obj, arr) {
//   const newObj = obj || global;
//   let result;
//   newObj.p = this;

//   if (!arr) {
//     result = newObj.p();
//   } else {
//     let newArguments = [];
//     for (let i = 0; i < arr.length; i++) {
//       newArguments.push('arr[' + i + ']');
//     }
//     result = eval('newObj.p(' + newArguments + ')');
//   }
//   delete newObj.p;
//   return result;
// };

// Function.prototype.myBind = function (obj) {
//   const _this = this;
//   const arr = Array.prototype.slice.call(arguments, 1);
//   console.log({ arr });
//   return function () {
//     _this.apply(obj);
//   };
// };

getName.myCall(nameObj, 'a', 'b', 'c', 'd');
// getName.myApply(nameObj, ['a', 'b', 'c', 'd']);
// getName.myBind(nameObj, 'a')();

// const newGetName = getName.myBind(nameObj);
// newGetName();
// const newnewGetName = getName.myBind({ name: 'jojo' });
// newnewGetName();
