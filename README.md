# timber
Fast and lightweight logging server built with go

## Installation
Builds for windows, linux and mac available in the builds folder
- [Windows](https://github.com/mohitudupa/timber/blob/main/builds/windows/timber.exe)
- [linux](https://github.com/mohitudupa/timber/blob/main/builds/linux/timber)
- [Mac](https://github.com/mohitudupa/timber/blob/main/builds/mac/timber)

## Running the server
Run the binary/executable
``` 
./timber
```
By default, the server will be running on port 36036.
Log files will be stored in ./timber-logs and a default log file called ```default.log``` will be created.

Any ```POST``` requests made to ```/logs/<log-file-name>/``` will be logged in ```<log-directory>/<log-file-name>```


## Configuration
By default, the log server will be running with the following config
```
data: ./timber-logs
port: 36036
logs: [default.log]
```

To add your own configs, create a file ```timberconf.json``` in the the working directory with the following fields
```
{
    "data": "./timber-logs", // Log file directory
    "port": 36036, // HTTP Port
    "logs": ["default.log"] // List of log files
}
```

## Docker
Building docker image
```
docker build -f Dockerfile -t timber .
```
Optionally the image is also available at dockerhub with
```
docker pull timber:latest
```

Running the image
```
docker container run -d -p 36036:36036 -v <your-volume>:/application/timber-logs timber:latest
```