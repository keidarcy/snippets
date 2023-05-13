const Schema = require('./employees_pb');

const john = new Schema.Employee();
john.setId(1001);
// john.setName('John');
// john.setSalary(50000);
console.log(john);

const mary = new Schema.Employee();
mary.setId(1002);
mary.setName('Mary');
mary.setSalary(60000);
console.log(mary);

const mike = new Schema.Employee();
mike.setId(1003);
mike.setName('Mike');
// mike.setSalary(70000);
console.log(mike);

const employees = new Schema.Employees();
employees.addEmployees(john);
employees.addEmployees(mary);
employees.addEmployees(mike);

const bytes = employees.serializeBinary();
require('fs').writeFileSync('binarydata', bytes);

const employees2 = Schema.Employees.deserializeBinary(bytes);
// console.log(employees2.toString());
