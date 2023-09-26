package main

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var (
	resourceGroupName     = "production"
	resourceGroupLocation = "westeurope"
	snet                  = "subnetprod1"
	snet2                 = "subnetprod2"
	VNETName              = "vnetprodwesteu"
)

func main() {

	pulumi.Run(func(ctx *pulumi.Context) error {
		rg, err := resourceGroup(ctx, resourceGroupName, resourceGroupLocation)
		if err != nil {
			fmt.Errorf("resource creation failed %s", err)
		}

		subnets := []SubnetConfig{
			{Name: snet, AddressRange: "10.0.1.0/24"},
			{Name: snet2, AddressRange: "10.0.2.0/24"},
		}

		opts := VNETOptions{
			Name:             VNETName,
			Location:         resourceGroupLocation,
			ResourceGroup:    rg,
			ResourceLocation: pulumi.String(resourceGroupLocation),
			Subnets:          subnets,
			AddressSpaces:    []string{"10.0.0.0/16"},
			Tags: pulumi.StringMap{
				"environment": pulumi.String("Production"),
			},
		}
		_, err = CreateVNET(ctx, opts)
		if err != nil {
			return err
		}

		return nil
	})

}
