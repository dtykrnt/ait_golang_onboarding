# Golang Skeleton (CLI Based)

This is Golang skeleton CLI based, this project runs on native Go with library [cli](https://github.com/urfave/cli) to make go project run on CLI based

## Purpose 
This structure help us to serve this application as CLI based on its modules.
for example if you have 2 modules like moduleA and moduleB, you can serve each one separately

## How to run

```bash
go run main.go serve-api
```

## Project Structure

`cmd` contains files that will be serving the service from `internals`.

`config` contains config files that will be used. For example like redis.go or db.go etc. 

`core` contains library, middlewares and also helpers like utility etc

`internals` contains modules/domains that is serving the purpose of this project, your models, repository, usecase and handler should be in this folder
- keep each of the module independent
- each module will serve itself

`resources` contains useful file that helps the project run smoothly for example like open api files, migrations etc. 

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.