package deadmanssnitch

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSnitch() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSnitchRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Snitch name",
			},
		},
	}
}

func dataSourceSnitchRead(d *schema.ResourceData, meta interface{}) error {
	// TO-DO
	// create an API written in GO to call here
}
