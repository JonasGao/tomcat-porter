tomcat-porter
=============

A little tool to read tomcat server.xml and print all using port.

## Usage

```text
NAME:
   tomcat-porter - Parse tomcat server.xml, and print all ports.
USAGE:
   tomcat-porter [global options] [path to server.xml or dir]
VERSION:
   1.0.2
COMMANDS:
   help, h  Shows a list of commands or help for one command
GLOBAL OPTIONS:
   --quite, -q             Ignore filesystem error. (default: false)
   --mode value, -m value  Print ports in list/simple/table mode. (default: "lis
t")
   --help, -h              show help
   --version, -v           print the version

```

### 1. Auto parse conf/server.xml and print ports.

```shell
cd tomcat
tomcat-porter
```

### 2. Specify server.xml

```shell
tomcat-porter /usr/local/apache-tomcat-xxx/conf/server.xml
```

### 3. Walk specified dir, search server.xml

```shell
tomcat-porter /usr/local/
```

**Notice: Path must end with '/'.**

Ignore filesystem error with `-q`:

```shell
tomcat-porter -q /usr/local/
```

Print table or other with `-m`

```shell
tomcat-porter -m table /usr/local/
tomcat-porter -m simple /usr/local/
```

## Build

Example

```shell
go build -ldflags "-X main.Version=1 \
  -X main.BuildTime=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'` \
  -X main.GitHash=`git rev-parse HEAD`"
```

On windows, run `Build.ps1` build for linux.