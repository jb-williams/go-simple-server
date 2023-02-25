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
curl -H "User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.109 Safari/537.36"
