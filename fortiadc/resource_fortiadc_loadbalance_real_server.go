package fortiadc

import (
	"errors"

	"github.com/Ouest-France/gofortiadc"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceFortiadcLoadbalanceRealServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceFortiadcLoadbalanceRealServerCreate,
		Read:   resourceFortiadcLoadbalanceRealServerRead,
		Update: resourceFortiadcLoadbalanceRealServerUpdate,
		Delete: resourceFortiadcLoadbalanceRealServerDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"address6": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "::",
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "enable",
			},
		},
	}
}

func resourceFortiadcLoadbalanceRealServerCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*gofortiadc.Client)

	req := gofortiadc.LoadbalanceRealServer{
		Mkey:     d.Get("name").(string),
		Status:   d.Get("status").(string),
		Address:  d.Get("address").(string),
		Address6: d.Get("address6").(string),
	}

	err := client.LoadbalanceCreateRealServer(req)
	if err != nil {
		return err
	}

	d.SetId(req.Mkey)

	return resourceFortiadcLoadbalanceRealServerRead(d, m)
}

func resourceFortiadcLoadbalanceRealServerRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*gofortiadc.Client)

	rs, err := client.LoadbalanceGetRealServer(d.Id())
	if errors.Is(err, gofortiadc.ErrNotFound) {
		// If not found, remove from state
		d.SetId("")
		return nil
	}
	if err != nil {
		return err
	}

	arguments := map[string]interface{}{
		"name":     rs.Mkey,
		"address":  rs.Address,
		"address6": rs.Address6,
		"status":   rs.Status,
	}

	for arg, value := range arguments {
		err = d.Set(arg, value)
		if err != nil {
			return err
		}
	}

	return nil
}

func resourceFortiadcLoadbalanceRealServerUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*gofortiadc.Client)

	req := gofortiadc.LoadbalanceRealServer{
		Mkey:     d.Get("name").(string),
		Status:   d.Get("status").(string),
		Address:  d.Get("address").(string),
		Address6: d.Get("address6").(string),
	}

	err := client.LoadbalanceUpdateRealServer(req)
	if err != nil {
		return err
	}

	return resourceFortiadcLoadbalanceRealServerRead(d, m)
}

func resourceFortiadcLoadbalanceRealServerDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*gofortiadc.Client)

	return client.LoadbalanceDeleteRealServer(d.Get("name").(string))
}
