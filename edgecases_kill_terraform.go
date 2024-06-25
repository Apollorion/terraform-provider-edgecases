// resource_server.go
package main

import (
	"bufio"
	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"os"
	"os/exec"
	"strings"
)

func killTerraform() *schema.Resource {
	return &schema.Resource{
		Create: killTerraformCreate,
		Read:   killTerraformRead,
		Update: killTerraformUpdate,
		Delete: killTerraformDelete,

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

func killTerraformProcess() error {
	f, err := os.OpenFile("tflog.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	// find the terraform or opentofu process
	cmd := exec.Command("ps", "aux")
	out, err := cmd.Output()
	if err != nil {
		return err
	}

	var line string
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		scanText := scanner.Text()
		if strings.Contains(scanText, "terraform") || strings.Contains(scanText, "tofu") {
			line = scanText
			break
		}
	}

	processID := strings.Fields(line)[1]

	// kill the process
	cmd = exec.Command("kill", "-9", processID)
	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func killTerraformCreate(d *schema.ResourceData, meta interface{}) error {
	d.SetId(uuid.New().String())

	if d.Get("during_create").(bool) {
		err := killTerraformProcess()
		if err != nil {
			return err
		}
	}

	return nil
}

func killTerraformRead(d *schema.ResourceData, m interface{}) error {
	if d.Get("during_read").(bool) {
		err := killTerraformProcess()
		if err != nil {
			return err
		}
	}

	return nil
}

func killTerraformUpdate(d *schema.ResourceData, m interface{}) error {
	if d.Get("during_update").(bool) {
		err := killTerraformProcess()
		if err != nil {
			return err
		}
	}
	return nil
}

func killTerraformDelete(d *schema.ResourceData, m interface{}) error {
	d.SetId("")

	if d.Get("during_delete").(bool) {
		err := killTerraformProcess()
		if err != nil {
			return err
		}
	}

	return nil
}
