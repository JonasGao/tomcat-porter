tomcat-porter
=============

A little tool to read tomcat server.xml and print all using port.

## Usage

### 0. Print version

```shell
tomcat-porter version
```

### 1. Auto parse conf/server.xml and print ports.

```shell
cd tomcat
tomcat-porter
```

`tomcat-porter` auto read `conf/server.xml` in current dir. And print ports like:

```text
Parsing ./conf/server.xml
Server port   : 8005                            
  Service Catalina                              
    Connector : port = 8080, redirectPort = 8443
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
tomcat-porter /usr/local/ -q
```

## Build

Example

```shell
go build -ldflags "-X main.Version=1 \
  -X main.BuildTime=`TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ'` \
  -X main.GitHash=`git rev-parse HEAD`"
```

On windows, run `Build.ps1` build for linux.