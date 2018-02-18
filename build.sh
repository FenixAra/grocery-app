#!/usr/bin/env bash


export APP_NAME="Sessions"
export DMS_URL=""
export PORT=3000
export DATABASE_URL='postgres://postgres@localhost:5432/grocery_app?sslmode=disable'
export DATABASE_URL_TEST='postgres://postgres@localhost:5432/grocery_app_test?sslmode=disable'
export MIGRATION_FILE_PATH=file://$GOPATH/src/github.com/FenixAra/grocery-app/migrations
export MAX_DB_CONNECTIONS=15
export SERVICE_TOKEN="6D7U6Xh2uRRUxcVRPeSGNfxmVT4tepxuMVC5vLkL2LtLBrDBstrz2qbKuHtZVsPM"
export REDIS_URL="redis://localhost:6379"
export REDIS_CONN_POOL_IDLE_TIMEOUT_MINS="5"
export REDIS_MAX_ACTIVE_CONNECTIONS="20"
export EXPIRY_TIME="10"

go install
if [ $? != 0 ]; then
  echo "## Build Failed ##"
  exit
fi


echo "Doing some cleaning ..."
go clean
echo "Done."

echo "Running goimport ..."
goimports -w=true .
echo "Done."

echo "Running go vet ..."
go vet ./internal/...
if [ $? != 0 ]; then
  exit
fi
echo "Done."

echo "Running go generate ..."
go generate ./internal/...
echo "Done."

echo "Running go format ..."
gofmt -w .
echo "Done."

echo "Running go build ..."
go build -race
if [ $? != 0 ]; then
  echo "## Build Failed ##"
  exit
fi
echo "Done."

echo "Running unit test ..."
go test -p=1 ./internal/... ./utils/...
if [ $? == 0 ]; then
    echo "Done."
	echo "## Starting service ##"
    ./grocery-app
fi
