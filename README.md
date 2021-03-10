# secrets
A simple package to handle secrets providers

## Test

To test locally I used [localstack](https://github.com/localstack/localstack) to mimic some aws services, and you have to rename the file `.env.example` to `.env` and fill the fields 

```
localstack start
```

and create a aws secret locally 

```
aws secretsmanager create-secret --name local-secret --secret-string '{"username":"bob","password":"abc123xyz456"}' --endpoint http://localhost:4566

```

after this just run the this command

```
go test ./... -v --cover
```

## TODO
- AWS Ok
- Azuke -
- GPC - 
- Vault -