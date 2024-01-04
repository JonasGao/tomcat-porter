tomcat-porter
=============

A little tool to read tomcat server.xml and print all using port.

## Usage

```shell
cd tomcat
tomcat-porter conf/server.xml
```

## Build

```shell
go build -ldflags "-X main.Version=1 \
  -X main.BuildTime=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'` \
  -X main.GitHash=`git rev-parse HEAD`"
```
