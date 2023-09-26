package main

import (
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure/resources"
	"github.com/pulumi/pulumi-azure/sdk/v5/go/azure/network"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VNETOptions struct {
	Name             string
	Location         string
	ResourceGroup    *resources.ResourceGroup
	ResourceLocation pulumi.String
	Subnets          []SubnetConfig
	AddressSpaces    []string
	Tags             pulumi.StringMap
}

type SubnetConfig struct {
	Name         string
	AddressRange string
}

func CreateVNET(ctx *pulumi.Context, opts VNETOptions) (*network.VirtualNetwork, error) {
	subnetArgs := make(network.VirtualNetworkSubnetArray, len(opts.Subnets))

	for i, subnet := range opts.Subnets {
		subnetArgs[i] = &network.VirtualNetworkSubnetArgs{
			Name:          pulumi.String(subnet.Name),
			AddressPrefix: pulumi.String(subnet.AddressRange),
		}
	}

	addressSpaces := pulumi.StringArray{}
	for _, addressSpace := range opts.AddressSpaces {
		addressSpaces = append(addressSpaces, pulumi.String(addressSpace))
	}

	vnet, err := network.NewVirtualNetwork(ctx, opts.Name, &network.VirtualNetworkArgs{
		Location:          opts.ResourceLocation,
		ResourceGroupName: opts.ResourceGroup.Name,
		AddressSpaces:     addressSpaces,
		Subnets:           subnetArgs,
		Tags:              opts.Tags,
	})

	if err != nil {
		return nil, err
	}

	return vnet, nil
}
