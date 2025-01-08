# goomu
A lightweight and efficient file server written in Go

## Description
Goomu is a simple yet powerful file server implementation in Go, designed to serve files efficiently over HTTP. It provides an easy way to share and access files through a web interface.

## Features
- Simple HTTP file serving
- Easy to set up and configure
- Lightweight and efficient
- Cross-platform compatibility

## Docker Image
``` 
docker run -d -p 8080:8080 goomu
```

## Installation
```bash
git clone https://github.com/osamikoyo/goomu.git && cd goomu && go install cmd/goomu/main.go
```

## Usage

1. Start the server:

```bash
goomu 
```

2.Acces your files thtough the web interface at ```http://localhost:8080``` (default port)

## Configuration

goomu will find config.toml file in:
```bash
    $HOME/goomu/config.toml
    $CONFIG_DIR/goomu/config.toml
```

