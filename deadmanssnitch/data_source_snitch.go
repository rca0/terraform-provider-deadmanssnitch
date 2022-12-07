package deadmanssnitch

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSnitch() *schema.Resource {
	fmt.Printf("Hello world, i'm dataSourceSnitch()")
	return &schema.Resource{}
}
