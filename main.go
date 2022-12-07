package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/rca0/terraform-deadmanssnitch/deadmanssnitch"
)

var version string

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() terraform.ResourceProvider {
			return deadmanssnitch.Provider(version)
		},
	})
}
