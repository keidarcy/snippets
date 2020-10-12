function Horse(name) {
  this.name = name;
  this.voice = function () {
    console.log('yoyo');
    return this;
  };
}

const myHorse = new Horse('yuyu');
myHorse.voice().voice();
