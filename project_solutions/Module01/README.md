# Creating the Basic Service

### Overview

In this lab you will be creating your first basic http service

at the end of this module , your project should look like this:

```
├── README.md

├── cmd

│  └── main.go

├── go.mod

└── internal

  └── routes

​    └── routes.go
```

## Getting started

### Creating a module

1. create a module for your go project - please note that you should choose the username for repo that you will commit to. From the root directory execute:

```bash
$ go mod init github.com/<someuser>/gameserver
```

2. Optional - you can use an existing github repo or create one (if you want to persist your code to GH); to do so, easiest way to accomplish is using VS Code Git integration options (initialize repo and publish)

### Adding Routes

1. add a file named routes.go following the directory structure outlined above 
2. create a Handler() func that contains the following routes 

```bash
func Handlers() {
	http.HandleFunc("/welcome", welcome)
	http.Handle("/greet", http.HandlerFunc(greet))
}
```

3. implement welcome and greet functions in a way that accessing those endpoints should result in "welcome to go!!!" and "hello from go!!!" respectively

### Adding main

1. add a file named main.go following the directory structure outlined above 
2. call the Handlers func from main
3. start the http server listening on port :5000

### Test 

make sure that your endpoints return the correct results

in your browser go to http://localhost:5000/welcome , and  http://localhost:5000/greet