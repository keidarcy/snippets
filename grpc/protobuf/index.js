const fs = require('fs');
const employees = [];
employees.push({ name: 'John', salary: 50000, id: 1001 });

employees.push({ name: 'Mary', salary: 60000, id: 1002 });

employees.push({ name: 'Mike', salary: 70000, id: 1003 });

fs.writeFileSync('employees.json', JSON.stringify(employees));
