# API Company Monorepo

API monorepo is a api monorepo in go to access github api.

## Development

Use docker-compose [docker-compose](https://docs.docker.com/compose/) to run the code.

devmode consist of auto refresh when you save your go code

apply environment variable

don't forget to add GITHUB_ACCESS_TOKEN in docker-compose

```bash
cd api
export GITHUB_ACCESS_TOKEN=<your-token>
export $(egrep -v '^#' ./github-api/.env | xargs)
docker-compose build
docker-compose up -d
```
### PG Admin 

Can be access in localhost:5050

### Credential can be seen in 

/api/github-api/.env

```echo
# Used by pgadmin service 
PGADMIN_DEFAULT_EMAIL=test@test
PGADMIN_DEFAULT_PASSWORD=password

# Postgres Test
DB_HOST=postgres_test                      
DB_HOST=db                     # when running the app without docker 
DB_USER=marc
DB_PASSWORD=password
DB_NAME=github
DB_PORT=5432
```


to see the code running execute these commands
```bash
cd 
docker-compose logs -f
```

## Testing in Go

Use this test your go 

```GO
    cd api/github-api
    go test 
```
## Documentation of a package

```go
   cd api/github-api
   go docs
```

## Usage

```GO

```

## Deployment 

for the deployment for the api I use helm script to template the deployment

for each api folder contains manifest folder where your content for deployment are and use CI/CD tools to deploy your application

```bash
cd api/github-api
helm install manifest 
```


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)