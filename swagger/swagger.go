// Package swagger provides swagger UI resource files for swagger API service.
//
// Should be used with gf cli tool:
// gf pack public boot/data-packed.go -n boot -p=swagger -y
//
package swagger

import (
	_ "github.com/gogf/gf-swagger/boot"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gstr"
)

// Swagger is the struct for swagger feature management.
type Swagger struct {
	Schemes       []string
	Description   string
	BasicAuthUser string
	BasicAuthPass string
}

// Install installs the swagger to server.
func (swagger *Swagger) Install(s *ghttp.Server) error {
	s.AddStaticPath("/swagger", "swagger")
	s.Group("/swagger", func(group *ghttp.RouterGroup) {
		group.Hook("/*", ghttp.HOOK_BEFORE_SERVE, func(r *ghttp.Request) {
			if !r.BasicAuth(swagger.BasicAuthUser, swagger.BasicAuthPass) {
				r.ExitAll()
			}
			// Modify the swagger.json.
			if r.StaticFile != nil && gfile.Basename(r.URL.Path) == "swagger.json" {
				var content []byte
				if r.StaticFile.File != nil {
					content = r.StaticFile.File.Content()
				} else {
					content = gfile.GetBytes(r.StaticFile.Path)
				}
				j, _ := gjson.LoadContent(content)
				if !j.Contains("host") || gstr.Contains(j.GetString("host"), "{") {
					j.Set("host", r.GetHost())
				}
				if !j.Contains("basePath") || gstr.Contains(j.GetString("basePath"), "{") {
					j.Set("basePath", "/")
				}
				if swagger.Schemes != nil {
					j.Set("schemes", swagger.Schemes)
				}
				if swagger.Description != "" {
					j.Set("info.description", swagger.Description)
				}
				r.Response.WriteJson(j.MustToJson())
				r.ExitAll()
			}
		})
	})
	return nil
}

// Remove uninstalls swagger feature from server.
func (swagger *Swagger) Remove() error {
	return nil
}
