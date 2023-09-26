package main

import (
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func resourceGroup(ctx *pulumi.Context, rgname, location string) (*resources.ResourceGroup, error) {
	rg, err := resources.NewResourceGroup(ctx, "resourceGroup", &resources.ResourceGroupArgs{
		ResourceGroupName: pulumi.String(rgname),
		Location:          pulumi.String(location),
	})
	if err != nil {
		panic(err)
	}
	return rg, err
}
