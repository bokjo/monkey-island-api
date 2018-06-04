# monkey-island-api

###### I have no idea what I'm doing *[confused dog meme picture goes here...]* :)
___

# Monkey Island RESTFull API (written in Go)

### Requirement for running this api (NOT go gettable yet...)
- Correct GO instalation 
- Default POSTGRES installation (user: postgres, dbname: postgres)
- `Postman` or shell with `curl` (for testing)

Download this repo to: $GOPATH/src/github.com/bokjo/

```
cd $GOPATH/src/github.com/bokjo/monkey-island-api
go build / or go install
psql -U postgres -f initDBTables.sql
./monkey-island-api

fingers crossed :)
```

