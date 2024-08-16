@echo off

:: Set GOOS and GOARCH for Windows
set "GOOS=windows"
set "GOARCH=amd64"

:: Print the detected OS and architecture
echo Detected OS: %GOOS%
echo Detected ARCH: %GOARCH%

:: Run Docker Compose with the specified environment variables
:: Pass all necessary environment variables to Docker Compose
set "BUILD_LOC=%BUILD_LOC:="./cmd/telemetry"%"
docker-compose up --build --env GOOS=%GOOS% --env GOARCH=%GOARCH% --env BUILD_LOC=%BUILD_LOC%
