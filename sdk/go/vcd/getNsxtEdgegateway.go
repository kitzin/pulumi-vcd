// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNsxtEdgegateway(ctx *pulumi.Context, args *LookupNsxtEdgegatewayArgs, opts ...pulumi.InvokeOption) (*LookupNsxtEdgegatewayResult, error) {
	var rv LookupNsxtEdgegatewayResult
	err := ctx.Invoke("vcd:index/getNsxtEdgegateway:getNsxtEdgegateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxtEdgegateway.
type LookupNsxtEdgegatewayArgs struct {
	EdgeClusterId *string `pulumi:"edgeClusterId"`
	Name          string  `pulumi:"name"`
	Org           *string `pulumi:"org"`
	OwnerId       *string `pulumi:"ownerId"`
	// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
	Vdc *string `pulumi:"vdc"`
}

// A collection of values returned by getNsxtEdgegateway.
type LookupNsxtEdgegatewayResult struct {
	DedicateExternalNetwork bool    `pulumi:"dedicateExternalNetwork"`
	Description             string  `pulumi:"description"`
	EdgeClusterId           *string `pulumi:"edgeClusterId"`
	ExternalNetworkId       string  `pulumi:"externalNetworkId"`
	// The provider-assigned unique ID for this managed resource.
	Id        string                     `pulumi:"id"`
	Name      string                     `pulumi:"name"`
	Org       *string                    `pulumi:"org"`
	OwnerId   string                     `pulumi:"ownerId"`
	PrimaryIp string                     `pulumi:"primaryIp"`
	Subnets   []GetNsxtEdgegatewaySubnet `pulumi:"subnets"`
	// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
	Vdc string `pulumi:"vdc"`
}

func LookupNsxtEdgegatewayOutput(ctx *pulumi.Context, args LookupNsxtEdgegatewayOutputArgs, opts ...pulumi.InvokeOption) LookupNsxtEdgegatewayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNsxtEdgegatewayResult, error) {
			args := v.(LookupNsxtEdgegatewayArgs)
			r, err := LookupNsxtEdgegateway(ctx, &args, opts...)
			var s LookupNsxtEdgegatewayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNsxtEdgegatewayResultOutput)
}

// A collection of arguments for invoking getNsxtEdgegateway.
type LookupNsxtEdgegatewayOutputArgs struct {
	EdgeClusterId pulumi.StringPtrInput `pulumi:"edgeClusterId"`
	Name          pulumi.StringInput    `pulumi:"name"`
	Org           pulumi.StringPtrInput `pulumi:"org"`
	OwnerId       pulumi.StringPtrInput `pulumi:"ownerId"`
	// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
	Vdc pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupNsxtEdgegatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtEdgegatewayArgs)(nil)).Elem()
}

// A collection of values returned by getNsxtEdgegateway.
type LookupNsxtEdgegatewayResultOutput struct{ *pulumi.OutputState }

func (LookupNsxtEdgegatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtEdgegatewayResult)(nil)).Elem()
}

func (o LookupNsxtEdgegatewayResultOutput) ToLookupNsxtEdgegatewayResultOutput() LookupNsxtEdgegatewayResultOutput {
	return o
}

func (o LookupNsxtEdgegatewayResultOutput) ToLookupNsxtEdgegatewayResultOutputWithContext(ctx context.Context) LookupNsxtEdgegatewayResultOutput {
	return o
}

func (o LookupNsxtEdgegatewayResultOutput) DedicateExternalNetwork() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayResult) bool { return v.DedicateExternalNetwork }).(pulumi.BoolOutput)
}

func (o LookupNsxtEdgegatewayResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupNsxtEdgegatewayResultOutput) EdgeClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayResult) *string { return v.EdgeClusterId }).(pulumi.StringPtrOutput)
}

func (o LookupNsxtEdgegatewayResultOutput) ExternalNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayResult) string { return v.ExternalNetworkId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNsxtEdgegatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNsxtEdgegatewayResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNsxtEdgegatewayResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupNsxtEdgegatewayResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

func (o LookupNsxtEdgegatewayResultOutput) PrimaryIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayResult) string { return v.PrimaryIp }).(pulumi.StringOutput)
}

func (o LookupNsxtEdgegatewayResultOutput) Subnets() GetNsxtEdgegatewaySubnetArrayOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayResult) []GetNsxtEdgegatewaySubnet { return v.Subnets }).(GetNsxtEdgegatewaySubnetArrayOutput)
}

// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
func (o LookupNsxtEdgegatewayResultOutput) Vdc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayResult) string { return v.Vdc }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNsxtEdgegatewayResultOutput{})
}
