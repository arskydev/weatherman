# _Weatherman (Golang pet project)_

## _Requirements:_

### Weatherman app:
- API key for api.openweathermap.org - Required to get current weather for location. Set as IPGEO_API_KEY environment variable.

### API:
- App config need to be set in config/app_config.yaml . Values for CONF_PATH, PASS_ENV_NAME, APP_PORT are set for default.
- WEATHER_API_KEY environment variable. API key for api.ipgeolocation.io - Required to get current client location with ip.
- PostgreSQL - DB to store user data for authentication and authorization. Set password for DB user as DB_PASS (by default) environment variable. Config need to be set in config/db_config.yaml (by default)

> You can preset all environment variables for Docker containers with `set_env.sh` script

## _Build app_
### Build CLI:
`go build ./cmd/cli/main.go`

### API:
`go build ./cmd/api/main.go`

### Docker
`docker build -t "weatherman" .`

`docker build -t "weatherman_pg" ./db/`

## _Run app_

### CLI & API
`./main`

### Docker
`docker run -p 5432:5432 -e POSTGRES_PASSWORD=<YOUR-POSTGRES-PASSWORD> -d weatherman_pg`

`docker run -p 8888:8888 weatherman`

## _API endpoints_
### Auth:
|VERB|PATH|USED FOR|
| ------ | ------ | ------ |
|`POST` |`/auth/sign-up`| register new user|
|`POST` | `/auth/sign-in` | get JWT token for existing user. Required for using API|

### API:
|VERB|PATH|USED FOR|
| ------ | ------ | ------ |
|`GET`|`/api/get-mock`|get mock of API response|
|`GET`|`/api/get-weather`|get weather for your location|

## _Response demo_
```
{
    "location":{
        "city":"London",
        "country":"GB",
        "flag":"🇬🇧"
    },
    "suncycle":{
        "sunrise":"20 Jul 22 05:26 +0400",
        "sunset":"20 Jul 22 20:38 +0400"
    },
    "weather":{
        "temperature":24,
        "unit":"℃",
        "weathersymbol":"☁️",
        "weathertype":"Clouds"
    }
}
