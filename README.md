postformvalue
curl -v -X POST -F 'name=helloworld' 'http://localhost:3333/hello'
formvalue
curl -v -X POST 'http://localhost:3333/hello?name=helloworld'

curl -v -X POST 'http://localhost:3333/hello'

postformvalue
curl -v -X POST -F 'name=world&address=world' 'http://localhost:3333/form'
formvalue
curl -v -X POST 'http://localhost:3333/form?name=hello&address=world'
curl -v -X POST 'http://localhost:3333/form'
