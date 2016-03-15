#! /bin/bash

set -e

docker run --name gorma-cellar-db -p 5432:5432 -e POSTGRES_USER=gorma -e POSTGRES_DB=gorma -e POSTGRES_PASSWORD=gorma -d postgres

