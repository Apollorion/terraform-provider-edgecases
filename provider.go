// provider.go
package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"edgecases_die_during":     dieDuring(),
			"edgecases_kill_terraform": killTerraform(),
		},
	}
}
