package website

import (
	"os"

	"github.com/gofiber/template/html"
	"github.com/websublime/foundation/contracts"
	"github.com/websublime/foundation/kernel"
	"github.com/websublime/platform/modules/website/http"
)

func RegisterWebsiteModule() *contracts.Module {
	dir, _ := os.Getwd()

	home := http.Home{}

	return kernel.RegisterModule(
		contracts.ModuleConfig{
			Path:        "/",
			Name:        "website",
			Active:      true,
			Views:       html.New(dir, ".html"),
			ViewsLayout: "views/layouts/main",
		},
		contracts.Base{
			Controllers: contracts.Controllers{
				kernel.CreateControler(contracts.ControllerConfig{
					Path:    "",
					Name:    "home",
					Handler: home.Hello,
					Method:  "GET",
				}),
			},
		},
	)
}
