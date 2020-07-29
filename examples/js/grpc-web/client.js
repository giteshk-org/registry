const {ListApisRequest, ListApisResponse} = require('./registry_service_pb.js');
const {RegistryClient} = require('./registry_service_grpc_web_pb.js');

var client = new RegistryClient('http://localhost:9999');

var request = new ListApisRequest();
request.setParent('projects/google');
request.setPageSize(3);

console.log(request);

client.listApis(request, {}, (err, response) => {
  console.log(err);
  console.log(response);
});
