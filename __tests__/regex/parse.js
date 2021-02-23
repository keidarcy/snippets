const last = (stack) => stack[stack.length - 1];
function parse(re) {
  const stack = [[]];
  let i = 0;
  while (i < re.length) {
    const char = re[i];
    switch (char) {
      case '.': {
        last(stack).push({
          type: 'wildcard',
          quantifier: 'exactlyOne'
        });
        i++;
        continue;
      }
      case '?': {
        const lastElement = last(last(stack));
        if (!lastElement || lastElement.quantifier !== 'exactlyOne') {
          throw new Error('Quantifer must follow an unquantified element or group');
        }
        lastElement.quantifier = 'zeroOrOne';
        i++;
        continue;
      }
      case '*': {
        const lastElement = last(last(stack));
        if (!lastElement || lastElement.quantifier !== 'exactlyOne') {
          throw new Error('Quantifer must follow an unquantified element or group');
        }
        lastElement.quantifier = 'zeroOrMore';
        i++;
        continue;
      }
      case '+': {
        const lastElement = last(last(stack));
        if (!lastElement || lastElement.quantifier !== 'exactlyOne') {
          throw new Error('Quantifer must follow an unquantified element or group');
        }
        // lastElement.quantifier = 'oneOrMore';
        const zeroOrMoreCopy = { ...lastElement, quantifier: 'zeroOrMore' };
        last(stack).push(zeroOrMoreCopy);
        i++;
        continue;
      }

      case '(': {
        stack.push([]);
        i++;
        continue;
      }
      case ')': {
        if (stack.length <= 1) {
          throw new Error(`Not group to close at index ${i}`);
        }
        const states = stack.pop();
        last(stack).push({
          type: 'groupElement',
          states,
          quantifier: 'exactlyOne'
        });
        i++;
        continue;
      }
      case '\\': {
        if (i + 1 >= re.length) {
          throw new Error(`Bad escape character at index ${i}`);
        }
        last(stack).push({
          type: 'element',
          value: re[i + 1],
          quantifier: 'exactlyOne'
        });
        i += 2;
        continue;
      }
      default: {
        last(stack).push({
          type: 'element',
          value: char,
          quantifier: 'exactlyOne'
        });
        i++;
        continue;
      }
    }
  }
  if (stack.length !== 1) {
    throw new Error('Unmatched groups in regualr expression');
  }

  return stack[0];
}

// const { inspect } = require('util');

// const regex = 'a?c+(b.*c)+d';
// // console.log(inspect(parse(regex), false, Infinity));
// console.log(inspect(parse(regex), false, 10));

module.exports = parse;
