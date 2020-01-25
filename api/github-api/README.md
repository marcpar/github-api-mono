## github api

you have notice we have 2 dockerfile one is for development and one is for production build 

## GO testing multistage build 

this ensures that your code is tested in the stages of the application

checkout dockerfile

## AIR ensures 

with air package you can quickly code and autorefesh your upon saving 


```bash
db_1          | 2020-01-25 00:30:37.447 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
db_1          | 2020-01-25 00:30:37.448 UTC [1] LOG:  listening on IPv6 address "::", port 5432
db_1          | 2020-01-25 00:30:37.452 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
github-api_1  | 
github-api_1  |   __    _   ___  
github-api_1  |  / /\  | | | |_) 
github-api_1  | /_/--\ |_| |_| \_ // live reload for Go apps [v1.12.0]
github-api_1  | 
db_1          | 2020-01-25 00:30:37.470 UTC [23] LOG:  database system was shut down at 2020-01-24 18:09:56 UTC
db_1          | 2020-01-25 00:30:37.476 UTC [1] LOG:  database system is ready to accept connections
github-api_1  | watching .
github-api_1  | watching manifest
github-api_1  | watching manifest/charts
github-api_1  | watching manifest/templates
github-api_1  | watching manifest/templates/tests
github-api_1  | !exclude tmp
github-api_1  | building...
github-api_1  | go: finding github.com/lib/pq v1.3.0
github-api_1  | go: finding github.com/gorilla/mux v1.7.3
github-api_1  | go: downloading github.com/lib/pq v1.3.0
github-api_1  | go: downloading github.com/gorilla/mux v1.7.3
github-api_1  | go: extracting github.com/gorilla/mux v1.7.3
github-api_1  | go: extracting github.com/lib/pq v1.3.0
github-api_1  | # github-api
github-api_1  | ./model.go:4:16: syntax error: unexpected json, expecting semicolon or newline or }
github-api_1  | ./model.go:11:23: syntax error: unexpected [ after top level declaration
github-api_1  | failed to build, error: exit status 2

```
