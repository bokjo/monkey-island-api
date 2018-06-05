# monkey-island-api

###### I have no idea what I'm doing *[confused dog meme picture goes here...]* :)
___

# Monkey Island RESTFull API (written in Go)

### Requirement for running this api (NOT go gettable yet...)
- Correct GO instalation 
- Default POSTGRES installation (user: postgres, dbname: postgres)
- `Postman` or shell with `curl` (for testing)


### Server runs on `localhost` port `1234`
http://localhost:1234/api/

Download this repo to: $GOPATH/src/github.com/bokjo/

```
cd $GOPATH/src/github.com/bokjo/monkey-island-api
go build / or go install
psql -U postgres -f initDBTables.sql
./monkey-island-api

fingers crossed :)
```

## Endpoints
- [GET]
    - http://localhost:1234/api/cuddly_toys         - List of all toys
    - http://localhost:1234/api/cuddly_toys/monkeys - List of all monkeys
    - http://localhost:1234/api/cuddly_toys/dogs    - List of all dogs
    - http://localhost:1234/api/weapons             - List of all weapons
- [GET {ID}]   
    - http://localhost:1234/api/cuddly_toys/monkeys/{id} - Single monkey by ID
    - http://localhost:1234/api/cuddly_toys/dogs/{id}    - Single dog by ID
    - http://localhost:1234/api/weapons/{id}             - Single weapon by ID
- [POST] - JSON
    - http://localhost:1234/api/cuddly_toys/monkeys - Create new monkey
    - http://localhost:1234/api/cuddly_toys/dogs    - Create new dog
    ```json
    {
        "name": "Name",
        "energy_level": 1
    }
    ```
    - http://localhost:1234/api/weapons             - Create new weapon
     ```json
    {
        "name": "Name",
        "power_level": 1
    }
    ```
- [PUT] - JSON | ID  
    - http://localhost:1234/api/cuddly_toys/monkeys/{id}  - Update monkey with ID
    - http://localhost:1234/api/cuddly_toys/dogs/{id}     - Update dog with ID
    - http://localhost:1234/api/weapons/{id}              - Update weapon with ID
- [DELETE]   
    - http://localhost:1234/api/cuddly_toys/monkeys/{id}  - Delete monkey with ID
    - http://localhost:1234/api/cuddly_toys/dogs/{id}     - Delete dog with ID
    - http://localhost:1234/api/weapons/{id}              - Delete weapon with ID