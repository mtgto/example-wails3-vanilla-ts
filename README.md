# example-vanilla-ts

1. Initialize the project
2. Add the method `GetTime` to the `GreetService` in `greetservice.go`
3. Generate the bindings

### Initialize the project

```console
$ wails3 init -t vanilla-ts -n example-vanilla-ts

 Wails (v3.0.0-alpha.41)  Init project
Creating project
----------------

Project Name      | example-vanilla-ts
Project Directory | /Users/user/work/wails/example-vanilla-ts
Template          | Vanilla + Vite (Typescript)
Template Source   | https://wails.io
Template Version  | v0.0.1


Project 'example-vanilla-ts' created successfully.
```

### Add the method `GetTime` to the `GreetService` in `greetservice.go`

```go
func (g *GreetService) GetTime() time.Time {
	return time.Now()
}
```

### Generate the bindings

```console
$ wails3 dev
```
