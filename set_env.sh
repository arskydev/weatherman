#!/bin/sh
# set_env.sh

echo "Initiating files..."
cp Dockerfile_raw Dockerfile
cp config/db_config_raw.yaml config/db_config.yaml

echo "Set PostgreSQL db password:"
read pg_pass
sed -i "s/PG_PASS/$pg_pass/g" Dockerfile

echo "Set WEATHER_API_KEY:"
read weather_api_key
sed -i "s/WEATHER_KEY/$weather_api_key/g" Dockerfile

echo "Set IPGEO_API_KEY:"
read ipgeo_api_key
sed -i "s/IPGEO_KEY/$ipgeo_api_key/g" Dockerfile

echo "Set database IP address:"
read db_addr
sed -i "s/DB_ADDR/$db_addr/g" config/db_config.yaml

echo "Done!"
