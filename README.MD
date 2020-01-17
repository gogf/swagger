# gf-swagger

Provides swagger UI resource files for Swagger 2.0 API service.

## Usage

### 1. Install gf cli tool

Download and install gf cli tool: https://github.com/gogf/gf-cli


### 2. Install gf-swagger

```sh
go get -u github.com/gogf/gf-swagger
```

### 3. Produce swagger.json

Use command `gf swagger` parses and produces the `/swagger/swaager.json` file.

> You can also use `gf swagger -s 8080` automatically parsing and producing `swagger.json` file when codes change in development.

#### 4. Pack swagger.json

Use gf cli tool command `gf pack swagger boot/swagger.go -n boot -p=swagger` packing the swagger.json file 
to binary in package `boot`.

#### 5. Use the swagger

```go
package main

import (
	_ "github.com/gogf/gf-swagger/swagger"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().AddStaticPath("/swagger", "swagger")
	g.Server().Run()
}
```

