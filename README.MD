# gf-swagger

Provides swagger API document feature for GoFrame project.

## Usage

### 1. Install gf cli tool

Download and install gf cli tool: https://github.com/gogf/gf-cli


### 2. Install gf-swagger

```sh
go get -u github.com/gogf/gf-swagger
```

### 3. Produce swagger.json

Use command `gf swagger --pack` parses and produces the `/swagger/swagger.json` file and packs it to `/data/data-swagger.go`.

> You can also use `gf run main.go --swagger` running your codes with hot-compiled-like feature and automatically packs the swagger files into your project.

#### 4. Use the swagger

Install the swagger using as plugin:

```go
package boot

import (
	"github.com/gogf/gf-swagger/swagger"
	"github.com/gogf/gf/frame/g"
)

func init() {
	s := g.Server()
	s.Plugin(&swagger.Swagger{
		Description:   g.Cfg().GetString("swagger.description"),
		BasicAuthUser: g.Cfg().GetString("swagger.user"),
		BasicAuthPass: g.Cfg().GetString("swagger.pass"),
	})
}

```

Start your server and enjoy!

