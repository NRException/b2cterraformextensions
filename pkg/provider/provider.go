package provider

import (
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Primary resource provider definition.
func BCExtensions() *schema.Provider {
    return &schema.Provider{
            ResourcesMap: map[string] *schema.Resource {
                "b2cextensions_aadb2c_authpolicy": authPolicy(),
                "b2cextensions_aadb2c_policykey": policyKey(),
            },
            Schema: map[string] *schema.Schema {
                "client_id": { Type: schema.TypeString, Required: true }, 
                "client_secret": {Type: schema.TypeString, Required: true },
            },
    }
}
