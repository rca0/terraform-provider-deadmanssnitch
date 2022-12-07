package deadmanssnitch

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceSnitch() *schema.Resource {
	fmt.Printf("Hello world, i'm resourceSnitch()")
	return &schema.Resource{}
}
