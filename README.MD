# gf-swagger

Provides swagger API document feature for GoFrame project.

## Usage

### 1. Install gf cli tool

Download and install gf cli tool: https://github.com/gogf/gf-cli


### 2. Install gf-swagger

```sh
go get -u github.com/gogf/gf-swagger
```

### 3. Produce and pack swagger.json

Use command `gf swagger --pack` parses and produces the `/swagger/swagger.json` file and packs it to `/data/data-swagger.go`.

> You can also use `gf run main.go --swagger -y` running your codes with hot-compiled-like feature and automatically packs the swagger files into your project.

#### 4. Use the swagger plugin

Install the swagger using as plugin:

```go
package main

import (
	"github.com/gogf/gf-swagger/swagger"
	"github.com/gogf/gf/frame/g"
)

func main() {
    s := g.Server()
    s.Plugin(&swagger.Swagger{})
    s.SetPort(8199)
    s.Run()
}
```

Start your server and visit: http://127.0.0.1:8199/swagger

Enjoy!

## Comment and configuration

As the swagger files are parsed and produced by third party tool, 
you can get details about comment at: https://github.com/swaggo/swag .

GF swagger feature also supports configuration from file with `swagger` node.
The configuration items of the node please refer to the attributes of Swagger struct: 


