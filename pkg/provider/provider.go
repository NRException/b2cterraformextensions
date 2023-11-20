package provider

import (
    "github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func BCExtensions() *schema.Provider {
    return &schema.Provider{
            ResourcesMap: map[string]*schema.Resource{
                    "b2cextensions_aadb2c_authpolicy": authPolicy(),
                    "b2cextensions_aadb2c_policykey": policyKey(),
            },
    }
}
