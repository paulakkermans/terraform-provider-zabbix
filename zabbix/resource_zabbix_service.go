package zabbix

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/paulakkermans/go-zabbix-api"
)

func resourceZabbixService() *schema.Resource {
	return &schema.Resource{
		Create: resourceZabbixServiceCreate,
		Read:   resourceZabbixServiceRead,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Visible name of the service.",
			},
			"showsla": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"goodsla": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"algorithm": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"sortoder": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
			"triggerid": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
		},
	}
}

func createServiceObject(d *schema.ResourceData) *zabbix.Service {

	service := zabbix.Service{
		Algorithm: d.Get("algorithm").(int),
		Name:      d.Get("name").(string),
		Showsla:   d.Get("showsla").(int),
		Sortorder: zabbix.ValueType(d.Get("value_type").(int)),
		Goodsla:   zabbix.DataType(d.Get("data_type").(int)),
		Triggerid: d.Get("Triggerid").(int),
	}

	return &service
}

func resourceZabbixServiceCreate(d *schema.ResourceData, meta interface{}) error {
	service := createServiceObject(d)

	return createRetry(d, meta, createService, *service, resourceZabbixServiceRead)
}

func resourceZabbixServiceRead(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*zabbix.API)

	service, err := api.ServiceGetByID(d.Id())
	if err != nil {
		return err
	}

	d.Set("algorithm", item.Algorithm)
	d.Set("name", item.Name)
	d.Set("showsla", item.Showsla)
	d.Set("sortorder", item.Sortorder)
	d.Set("goodsla", item.Goodsla)
	d.Set("triggerid", item.Triggerid)

	log.Printf("[DEBUG] Item name is %s\n", item.Name)
	return nil
}
