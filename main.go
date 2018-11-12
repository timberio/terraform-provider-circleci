package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/timberio/terraform-provider-circleci/circleci"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: circleci.Provider})
}
