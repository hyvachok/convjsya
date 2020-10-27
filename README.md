# convjsya

Convjsya is a simple application that converts json to yaml and back.

## Install

```
$ go get -u github.com/hyvachok/convjsya
```

### Docker

[Docker image](https://hub.docker.com/r/hyvachok/convjsya)

```
$ docker pull hyvachok/convjsya
```

## Usage

```
$ convjsya test.json

$ convjsya test.json > test.yaml

$ convjsya -i test.json -o test.yaml

$ cat test.json | convjsya > test.yaml
```

### Usage in dockerfile

```dockerfile
FROM hyvachok/convjsya

RUN convjsya test.json > test2.yaml
RUN convjsya -i test.json -o test1.yaml
RUN cat test.json | convjsya > test3.yam
```