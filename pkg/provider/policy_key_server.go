package provider 

import (
        "log"
        "github.com/google/uuid"
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func policyKey() *schema.Resource {
        return &schema.Resource{
                Create: policyKeyCreate,
                Read:   policyKeyRead,
                Update: policyKeyUpdate,
                Delete: policyKeyDelete,

                Schema: map[string]*schema.Schema{
                        "uuid": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                },
        }
}

func policyKeyCreate(d *schema.ResourceData, m interface{}) error {
    // Set the state file object uuid.
    uid := uuid.New().String()
    log.Println("Setting object Id " + uid)
    d.SetId(uuid.New().String())
    
    // Finally return the read resource...
    return policyKeyRead(d, m)
}

func policyKeyRead(d *schema.ResourceData, m interface{}) error {
        return nil
}

func policyKeyUpdate(d *schema.ResourceData, m interface{}) error {
        return policyKeyRead(d, m)
}

func policyKeyDelete(d *schema.ResourceData, m interface{}) error {
        d.SetId("")
        return nil
}
