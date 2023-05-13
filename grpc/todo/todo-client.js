const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const packageDef = protoLoader.loadSync('todo.proto', {});
const grpcObject = grpc.loadPackageDefinition(packageDef);
const todoPackage = grpcObject.todoPackage;

const client = new todoPackage.Todo('localhost:40000', grpc.credentials.createInsecure());
client.createTodo(
  {
    text: process.argv[2],
  },
  (err, response) => {
    console.log('Client response createTodo: ', response);
  },
);

client.readTodos({}, (err, response) => {
  console.log('Client response readTodos: ', response);
});

const call = client.readTodosStream();
call.on('data', (item) => {
  console.log('Client response readTodosStream: ', item);
});

call.on('end', (e) => {
  console.log('Client response readTodosStream end: ', e);
});
