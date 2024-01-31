### Description

This is a simple app generator. It can be used to create a new app with the same structure as the one you already have.

### Install

```
go install github.com/ydssx/morphix/appgen@latest
```

### Usage

```
appgen [options] <appname> <protopath> <port>
```

### Options

```
-app  <appname>  The name of the app to create.
-proto <path>    The path of the proto file to use.
-port <port>     The port of the app to use.
```

### Examples

```
appgen -app myapp -proto /home/user/myapp.proto -port 8080
```