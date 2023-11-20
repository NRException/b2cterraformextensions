package provider 

import (
        "log"
        "github.com/google/uuid"
        "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func authPolicy() *schema.Resource {
        return &schema.Resource{
                Create: authPolicyCreate,
                Read:   authPolicyRead,
                Update: authPolicyUpdate,
                Delete: authPolicyDelete,

                Schema: map[string]*schema.Schema{
                        "uuid": &schema.Schema{
                                Type:     schema.TypeString,
                                Required: true,
                        },
                },
        }
}

func authPolicyCreate(d *schema.ResourceData, m interface{}) error {
    // Set the state file object uuid.
    uid := uuid.New().String()
    log.Println("Setting object Id " + uid)
    d.SetId(uuid.New().String())

    // Finally return the read resource...
    return authPolicyRead(d, m)
}

func authPolicyRead(d *schema.ResourceData, m interface{}) error {
        return nil
}

func authPolicyUpdate(d *schema.ResourceData, m interface{}) error {
        return authPolicyRead(d, m)
}

func authPolicyDelete(d *schema.ResourceData, m interface{}) error {
        d.SetId("")
        return nil
}
