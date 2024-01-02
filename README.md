# Golang Skeleton (CLI Based)

This is Golang skeleton CLI based, this project runs on native Go with library [cli](https://github.com/urfave/cli) to make go project run on CLI based

## Purpose 
This structure help us to serve this application as CLI based on its modules.
for example if you have 2 modules like moduleA and moduleB, you can serve each one separately

## How to run

```bash
go run main.go serve-api
```

### **Project Structure**

```
|   Dockerfile
|   go.mod
|   go.sum
|   main.go
|   README.md
|   
+---cmd
|       serve_api.go
|       
+---config
|       config.yaml
|       
+---core
|   |   database.go
|   |   redis.go
|   \---helpers
|           utils.go
|           
\---internals
    +---yourmodule
    |   |   handler.go
    |   |   router.go
    |   |   
    |   +---models
    |   |       entity.go
    |   |       
    |   \---usecase
    |           yourusecase.go
    |           yourusecase_test.go
    |           
    \---yourmodule2

```

### **Description of the structure**

### internals

```
\---internals
    +---yourmodule
    |   |   handler.go
    |   |   router.go
    |   |   
    |   +---models
    |   |       entity.go
    |   |       
    |   \---usecase
    |           yourusecase.go
    |           yourusecase_test.go
    |           
    \---yourmodule2
```

`internals` contains modules/domains that is serving the purpose of this project, your models, repository, usecase and handler should be in this folder

- keep each of the module independent
- each module will serve itself

### core

```
+---core
|   |   database.go
|   |   redis.go
|   \---helpers
|           utils.go
|
```

`core` contains core function/library that will be used by our modules, for example like redis, message broker, middlewares and also helpers like utility etc that will be used by `yourmodule`

### config

```
+---config
|       config.yaml
|
```

`config` contains config related files that will be used.

### cmd
```
+---cmd
|       serve_api.go
|   
```

`cmd` contains files that wrap our modules and serve it as `cli` based application. For example like `serve_api.go` that wrapping `yourmodule` as a command serve-api

### root files
```
|   Dockerfile
|   go.mod
|   go.sum
|   main.go
|   README.md
```

`main.go` works as our main function of the application

`go.mod` contains list of external libraries

`go.sum` this is auto generated files from command `go mod tidy`

`Dockerfile` here you can configure docker related commands

`README.md` Strongly suggest you write the readme file so we can know what is the application about

### resources
```
+---resources
|       migration.go
```

`resources` contains useful file that helps the project run smoothly for example like open api files, migrations etc. 

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.