# API Company Monorepo

API monorepo is a api monorepo in go to access github api.

## Development

Use docker-compose [docker-compose](https://docs.docker.com/compose/) to run the code.

devmode consist of auto refresh when you save your go code

```bash
cd api
docker-compose build
docker-compose up -d
```

to see the code running
```bash
docker-compose logs -f
```

go test 


## Usage

```GO

```

## Deployment 

for the deployment for the api I use helm script to template the deployment

for each api folder contains manifest folder where your content for deployment are
```bash
docker-compose logs -f
```


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)