# secrets
A simple package to handle secrets providers

## How To Use

```go

package main

import (
    "fmt"
    "encoding/json"
    "github.com/krismorte/secrets/aws"
    )

type mySecret struct {
	Username            string
	Password            string
	Engine              string
	Host                string
	Port                int
	Dbname              string
	DbClusterIdentifier string
}

func main(){
    var secret mySecret
    rawSecret = aws.GetSecret(secretID)

    b, _ := json.Marshal(rawSecret)

    json.Unmarshal(b, &secret)

    fmt.Println(secret.Host)
}

```

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