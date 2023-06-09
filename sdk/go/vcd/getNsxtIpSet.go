// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNsxtIpSet(ctx *pulumi.Context, args *LookupNsxtIpSetArgs, opts ...pulumi.InvokeOption) (*LookupNsxtIpSetResult, error) {
	var rv LookupNsxtIpSetResult
	err := ctx.Invoke("vcd:index/getNsxtIpSet:getNsxtIpSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxtIpSet.
type LookupNsxtIpSetArgs struct {
	EdgeGatewayId string  `pulumi:"edgeGatewayId"`
	Name          string  `pulumi:"name"`
	Org           *string `pulumi:"org"`
	// Deprecated: Deprecated in favor of `edge_gateway_id`. IP Set will inherit VDC from parent Edge Gateway.
	Vdc *string `pulumi:"vdc"`
}

// A collection of values returned by getNsxtIpSet.
type LookupNsxtIpSetResult struct {
	Description   string `pulumi:"description"`
	EdgeGatewayId string `pulumi:"edgeGatewayId"`
	// The provider-assigned unique ID for this managed resource.
	Id          string   `pulumi:"id"`
	IpAddresses []string `pulumi:"ipAddresses"`
	Name        string   `pulumi:"name"`
	Org         *string  `pulumi:"org"`
	OwnerId     string   `pulumi:"ownerId"`
	// Deprecated: Deprecated in favor of `edge_gateway_id`. IP Set will inherit VDC from parent Edge Gateway.
	Vdc *string `pulumi:"vdc"`
}

func LookupNsxtIpSetOutput(ctx *pulumi.Context, args LookupNsxtIpSetOutputArgs, opts ...pulumi.InvokeOption) LookupNsxtIpSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNsxtIpSetResult, error) {
			args := v.(LookupNsxtIpSetArgs)
			r, err := LookupNsxtIpSet(ctx, &args, opts...)
			var s LookupNsxtIpSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNsxtIpSetResultOutput)
}

// A collection of arguments for invoking getNsxtIpSet.
type LookupNsxtIpSetOutputArgs struct {
	EdgeGatewayId pulumi.StringInput    `pulumi:"edgeGatewayId"`
	Name          pulumi.StringInput    `pulumi:"name"`
	Org           pulumi.StringPtrInput `pulumi:"org"`
	// Deprecated: Deprecated in favor of `edge_gateway_id`. IP Set will inherit VDC from parent Edge Gateway.
	Vdc pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupNsxtIpSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtIpSetArgs)(nil)).Elem()
}

// A collection of values returned by getNsxtIpSet.
type LookupNsxtIpSetResultOutput struct{ *pulumi.OutputState }

func (LookupNsxtIpSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtIpSetResult)(nil)).Elem()
}

func (o LookupNsxtIpSetResultOutput) ToLookupNsxtIpSetResultOutput() LookupNsxtIpSetResultOutput {
	return o
}

func (o LookupNsxtIpSetResultOutput) ToLookupNsxtIpSetResultOutputWithContext(ctx context.Context) LookupNsxtIpSetResultOutput {
	return o
}

func (o LookupNsxtIpSetResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtIpSetResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupNsxtIpSetResultOutput) EdgeGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtIpSetResult) string { return v.EdgeGatewayId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNsxtIpSetResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtIpSetResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNsxtIpSetResultOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNsxtIpSetResult) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

func (o LookupNsxtIpSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtIpSetResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNsxtIpSetResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtIpSetResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupNsxtIpSetResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtIpSetResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

// Deprecated: Deprecated in favor of `edge_gateway_id`. IP Set will inherit VDC from parent Edge Gateway.
func (o LookupNsxtIpSetResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtIpSetResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNsxtIpSetResultOutput{})
}
