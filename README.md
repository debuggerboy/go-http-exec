# go-http-exec

Execute local commands over HTTP interface

## Build Executable:

Create a module

```
go mod init debuggerboy/go-http-exec
```

Get the dependencies:

```
go get .
```

Build `httpexec` executable binary.

```
go build httpexec.go
```

## Invocation

We can run `httpexec` as below:

```
./httpexec
```

## To-Do

- specify PORT

# WARNING

This is a UNFINISHED and BUGGY application, use at your own risk. Authors claim no guarantee that the application work as expected.
