// resource_server.go
package main

import (
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dieDuring() *schema.Resource {
	return &schema.Resource{
		Create: dieDuringCreate,
		Read:   dieDuringRead,
		Update: dieDuringUpdate,
		Delete: dieDuringDelete,

		Schema: map[string]*schema.Schema{
			"during_create": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"during_read": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"during_update": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"during_delete": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func dieDuringCreate(d *schema.ResourceData, m interface{}) error {
	d.SetId(uuid.New().String())

	if d.Get("during_create").(bool) {
		panic("die during create...")
	}

	return nil
}

func dieDuringRead(d *schema.ResourceData, m interface{}) error {
	if d.Get("during_read").(bool) {
		panic("die during read...")
	}

	return nil
}

func dieDuringUpdate(d *schema.ResourceData, m interface{}) error {
	if d.Get("during_update").(bool) {
		panic("die during update...")
	}
	return nil
}

func dieDuringDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	if d.Get("during_delete").(bool) {
		panic("die during delete...")
	}

	return nil
}
