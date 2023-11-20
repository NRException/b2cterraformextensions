package main

import (
    "github.com/nrexception/b2cterraformextensions/pkg/provider"
    "github.com/hashicorp/terraform-plugin-sdk/plugin"
    "github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Props to Infracloud on this one for the boilerplate and explaination...
// https://www.infracloud.io/blogs/developing-terraform-custom-provider/

func main() {
    plugin.Serve(&plugin.ServeOpts{
            ProviderFunc: func() terraform.ResourceProvider {
                return provider.BCExtensions()
            },
    })
}
