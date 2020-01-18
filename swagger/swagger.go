// Package swagger provides swagger UI resource files for swagger API service.
//
// Should be used with gf cli tool:
// gf pack public boot/data-packed.go -n boot -p=swagger -y
//
package swagger

import (
	_ "github.com/gogf/gf-swagger/boot"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
)

// Swagger is the struct for swagger feature management.
type Swagger struct {
	Title          string   // Title of the swagger API.
	Version        string   // Version of the swagger API.
	Schemes        []string // Supported schemes of the swagger API like "http", "https".
	Host           string   // The host of the swagger APi like "127.0.0.1", "www.mydomain.com"
	BasicPath      string   // The URI for the swagger API like "/", "v1", "v2".
	TermsOfService string   // As the attribute name.
	Description    string   // Detail description of the swagger API.
	BasicAuthUser  string   `c:"user"` // HTTP basic authentication username.
	BasicAuthPass  string   `c:"pass"` // HTTP basic authentication password.
}

// Install installs the swagger to server.
func (swagger *Swagger) Install(s *ghttp.Server) error {
	// Retrieve the configuration map and assign it to swagger object.
	m := g.Cfg().GetMap("swagger")
	if m != nil {
		gconv.Struct(m, swagger)
	}
	// The swagger resource files are served as static file service.
	s.AddStaticPath("/swagger", "swagger")
	// It here uses HOOK feature handling basic auth authentication and swagger.json modification.
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
				if swagger.Host != "" {
					j.Set("host", swagger.Host)
				} else if !j.Contains("host") || gstr.Contains(j.GetString("host"), "{") {
					j.Set("host", r.GetHost())
				}
				if swagger.BasicPath != "" {
					j.Set("basePath", swagger.BasicPath)
				} else if !j.Contains("basePath") || gstr.Contains(j.GetString("basePath"), "{") {
					j.Set("basePath", "/")
				}
				if len(swagger.Schemes) > 0 {
					j.Set("schemes", swagger.Schemes)
				}
				if swagger.Title != "" {
					j.Set("info.title", swagger.Title)
				}
				if swagger.Version != "" {
					j.Set("info.version", swagger.Version)
				}
				if swagger.TermsOfService != "" {
					j.Set("info.termsOfService", swagger.TermsOfService)
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
