package deadmanssnitch

import (
	"fmt"
	"log"

	"github.com/84codes/go-api/api"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

var version string

func Provider(v string) *schema.Provider {
	version = v
	log.Printf("terraform-provider-deadmanssnitch version: %s", version)
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"apikey": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("DEADMANSSNITCH_APIKEY", nil),
				Description: "The API Key used to authentication to the Deadmanssnitch API",
			},
			"baseurl": {
				Type:        schema.TypeString,
				DefaultFunc: schema.EnvDefaultFunc("DEADMANSSNITCH_URL", "https://api.deadmanssnitch.com/v1/snitches"),
				Optional:    true,
				Description: "Base URL to Deadmanssnitch website",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"deadmanssnitch_snitch": dataSourceSnitch(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"deadmanssnitch_snitch": resourceSnitch(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	useragent := fmt.Sprintf("terraform-provider-deadmanssnitch_v%s", version)
	log.Printf("[DEBUG] deadmanssnitch::provider::configure useragent: %v", useragent)
	// TO-DO: get the api content and try to handle some error (if necessary)
	return api.New(d.Get("baseurl").(string), d.Get("apikey").(string), useragent), nil
}
