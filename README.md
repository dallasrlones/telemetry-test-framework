# Telemetry Test Framework

This test framework runs in Go and requires Docker to build. In the future it will have custom OS and ARCH to build based on a specific environment.

This framework runs a process specified in the scripts folder as well as creates/modifies/delete's files, and finally makes an http get request to a configurable endpoint.

It then saves the expected results inside of telemetry-logs folder in the root of where the app was called from, inside of this folder is another folder with a timestamp of when the opperation happened, and then telemetry logs in JSON format to check and see if the output results match your expected results.

Examples of how to run the app are below.

## Building and Running

To build the app we can look inside of the /scripts folder. This is a collection of scripts to build the app as well as run a docker environment to test locally.

There are three different ways to run the framework, the first is dev in Docker where we can quickly run testing while itterating through changes without building the app.

The second is building the app based on the current system's OS and generated using Docker and then testing it wby running a script passing in the correct ENV variables.

And finally the third is running the app in a cron job every minute so we can automate this process.

## Run Without Building

To run the app in dev without building we can run the /scripts/start_local_docker.sh script, this will pass in the correct ENV variables as well as run a nodejs server to test the http call.

```bash
./scripts/start_local_docker.sh
```

Note that this is where you can modify any of the parameters you want to change being passed in via ENV variables. When you run in docker your output files will be inside of ./telemetry-app/telemetry-logs

## Build App

To build the app based on your current OS and ARCH we can run the /scripts/build_local_docker.sh script, this will build based on your current system's OS and ARCH as well as pass in the required ENV variables to build.

```bash
./scripts/build_local_docker.sh
```

When it's built it will go into /telemetry-app/cmd/telemetry/telemetry-test-framework as compiled code.

## Run Built App

To run the built app we can use the /scripts/run_local_build.sh script, this will pass in all the ENV variables we need.

```bash
./scripts/run_local_build.sh
```

You can modify this script's exports to change the file location, process to run, and url/port to hit with the http get request. The logs in this case will be in /whatfolderyoucalledthescriptfrom/telemetry-logs.

# Notes

Once the app is run take a look inside of the /telemetry-logs folder and open the current timestamp that you just ran (folder named by timestamp) this will contain multiple files, JSON and a .log file, the .log file contains an entire history of what happend like STD out and the JSON files contain the telemetry data to compare on.

When running inside of docker without building the telemetry data won't be accurate, docker is only for building and local quick testing. Because the docker env is set to linux alpine, it won't be accurate to what OS you're on. To get accuracy build the app with the script, and then run the build app with the script.