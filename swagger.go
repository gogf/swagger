// Package swagger provides swagger UI resource files for swagger API service.
//
// Should be used with gf cli tool:
// gf pack public ./public-packed.go -p=swagger -y
//
package swagger

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"net/http"
	"time"
)

// Swagger is the struct for swagger feature management.
type Swagger struct {
	Info          SwaggerInfo // Swagger information.
	Schemes       []string    // Supported schemes of the swagger API like "http", "https".
	Host          string      // The host of the swagger APi like "127.0.0.1", "www.mydomain.com"
	BasicPath     string      // The URI for the swagger API like "/", "v1", "v2".
	BasicAuthUser string      `c:"user"` // HTTP basic authentication username.
	BasicAuthPass string      `c:"pass"` // HTTP basic authentication password.
}

// SwaggerInfo is the information field for swagger.
type SwaggerInfo struct {
	Title          string // Title of the swagger API.
	Version        string // Version of the swagger API.
	TermsOfService string // As the attribute name.
	Description    string // Detail description of the swagger API.
}

const (
	Name               = "gf-swagger"
	Author             = "john@goframe.org"
	Version            = "v1.2.0"
	Description        = "gf-swagger provides swagger API document feature for GoFrame project. https://github.com/gogf/gf-swagger"
	MaxAuthAttempts    = 10          // Max authentication count for failure try.
	AuthFailedInterval = time.Minute // Authentication retry interval after last failed.
)

// Name returns the name of the plugin.
func (swagger *Swagger) Name() string {
	return Name
}

// Author returns the author of the plugin.
func (swagger *Swagger) Author() string {
	return Author
}

// Version returns the version of the plugin.
func (swagger *Swagger) Version() string {
	return Version
}

// Description returns the description of the plugin.
func (swagger *Swagger) Description() string {
	return Description
}

// Install installs the swagger to server as a plugin.
// It implements the interface ghttp.Plugin.
func (swagger *Swagger) Install(s *ghttp.Server) error {
	// Retrieve the configuration map and assign it to swagger object.
	m := g.Cfg().GetMap("swagger")
	if m != nil {
		if err := gconv.Struct(m, swagger); err != nil {
			s.Logger().Fatal(err)
		}
	}
	// The swagger resource files are served as static file service.
	s.AddStaticPath("/swagger", "swagger")
	// It here uses HOOK feature handling basic auth authentication and swagger.json modification.
	s.Group("/swagger", func(group *ghttp.RouterGroup) {
		group.Hook("/*", ghttp.HookBeforeServe, func(r *ghttp.Request) {
			if swagger.BasicAuthUser != "" {
				// Authentication security checks.
				var (
					authCacheKey = fmt.Sprintf(`swagger_auth_failed_%s`, r.GetClientIp())
					v, _         = gcache.GetVar(authCacheKey)
					authCount    = v.Int()
				)
				if authCount > MaxAuthAttempts {
					r.Response.WriteStatus(
						http.StatusForbidden,
						"max authentication count exceeds, please try again in one minute!",
					)
					r.ExitAll()
				}
				// Basic authentication.
				if !r.BasicAuth(swagger.BasicAuthUser, swagger.BasicAuthPass) {
					gcache.Set(authCacheKey, authCount+1, AuthFailedInterval)
					r.ExitAll()
				}
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
					j.Set("host", r.Host)
				}
				if swagger.BasicPath != "" {
					j.Set("basePath", swagger.BasicPath)
				} else if !j.Contains("basePath") || gstr.Contains(j.GetString("basePath"), "{") {
					j.Set("basePath", "/")
				}
				if len(swagger.Schemes) > 0 {
					j.Set("schemes", swagger.Schemes)
				}
				if swagger.Info.Title != "" {
					j.Set("info.title", swagger.Info.Title)
				}
				if swagger.Info.Version != "" {
					j.Set("info.version", swagger.Info.Version)
				}
				if swagger.Info.TermsOfService != "" {
					j.Set("info.termsOfService", swagger.Info.TermsOfService)
				}
				if swagger.Info.Description != "" {
					j.Set("info.description", swagger.Info.Description)
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
