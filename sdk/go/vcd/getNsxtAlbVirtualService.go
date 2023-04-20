// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNsxtAlbVirtualService(ctx *pulumi.Context, args *LookupNsxtAlbVirtualServiceArgs, opts ...pulumi.InvokeOption) (*LookupNsxtAlbVirtualServiceResult, error) {
	var rv LookupNsxtAlbVirtualServiceResult
	err := ctx.Invoke("vcd:index/getNsxtAlbVirtualService:getNsxtAlbVirtualService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxtAlbVirtualService.
type LookupNsxtAlbVirtualServiceArgs struct {
	EdgeGatewayId string  `pulumi:"edgeGatewayId"`
	Name          string  `pulumi:"name"`
	Org           *string `pulumi:"org"`
	// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
	Vdc *string `pulumi:"vdc"`
}

// A collection of values returned by getNsxtAlbVirtualService.
type LookupNsxtAlbVirtualServiceResult struct {
	ApplicationProfileType string `pulumi:"applicationProfileType"`
	CaCertificateId        string `pulumi:"caCertificateId"`
	Description            string `pulumi:"description"`
	EdgeGatewayId          string `pulumi:"edgeGatewayId"`
	Enabled                bool   `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id                   string                                `pulumi:"id"`
	Name                 string                                `pulumi:"name"`
	Org                  *string                               `pulumi:"org"`
	PoolId               string                                `pulumi:"poolId"`
	ServiceEngineGroupId string                                `pulumi:"serviceEngineGroupId"`
	ServicePorts         []GetNsxtAlbVirtualServiceServicePort `pulumi:"servicePorts"`
	// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
	Vdc              string `pulumi:"vdc"`
	VirtualIpAddress string `pulumi:"virtualIpAddress"`
}

func LookupNsxtAlbVirtualServiceOutput(ctx *pulumi.Context, args LookupNsxtAlbVirtualServiceOutputArgs, opts ...pulumi.InvokeOption) LookupNsxtAlbVirtualServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNsxtAlbVirtualServiceResult, error) {
			args := v.(LookupNsxtAlbVirtualServiceArgs)
			r, err := LookupNsxtAlbVirtualService(ctx, &args, opts...)
			var s LookupNsxtAlbVirtualServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNsxtAlbVirtualServiceResultOutput)
}

// A collection of arguments for invoking getNsxtAlbVirtualService.
type LookupNsxtAlbVirtualServiceOutputArgs struct {
	EdgeGatewayId pulumi.StringInput    `pulumi:"edgeGatewayId"`
	Name          pulumi.StringInput    `pulumi:"name"`
	Org           pulumi.StringPtrInput `pulumi:"org"`
	// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
	Vdc pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupNsxtAlbVirtualServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtAlbVirtualServiceArgs)(nil)).Elem()
}

// A collection of values returned by getNsxtAlbVirtualService.
type LookupNsxtAlbVirtualServiceResultOutput struct{ *pulumi.OutputState }

func (LookupNsxtAlbVirtualServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtAlbVirtualServiceResult)(nil)).Elem()
}

func (o LookupNsxtAlbVirtualServiceResultOutput) ToLookupNsxtAlbVirtualServiceResultOutput() LookupNsxtAlbVirtualServiceResultOutput {
	return o
}

func (o LookupNsxtAlbVirtualServiceResultOutput) ToLookupNsxtAlbVirtualServiceResultOutputWithContext(ctx context.Context) LookupNsxtAlbVirtualServiceResultOutput {
	return o
}

func (o LookupNsxtAlbVirtualServiceResultOutput) ApplicationProfileType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbVirtualServiceResult) string { return v.ApplicationProfileType }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbVirtualServiceResultOutput) CaCertificateId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbVirtualServiceResult) string { return v.CaCertificateId }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbVirtualServiceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbVirtualServiceResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbVirtualServiceResultOutput) EdgeGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbVirtualServiceResult) string { return v.EdgeGatewayId }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbVirtualServiceResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNsxtAlbVirtualServiceResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNsxtAlbVirtualServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbVirtualServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbVirtualServiceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbVirtualServiceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbVirtualServiceResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtAlbVirtualServiceResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupNsxtAlbVirtualServiceResultOutput) PoolId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbVirtualServiceResult) string { return v.PoolId }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbVirtualServiceResultOutput) ServiceEngineGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbVirtualServiceResult) string { return v.ServiceEngineGroupId }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbVirtualServiceResultOutput) ServicePorts() GetNsxtAlbVirtualServiceServicePortArrayOutput {
	return o.ApplyT(func(v LookupNsxtAlbVirtualServiceResult) []GetNsxtAlbVirtualServiceServicePort { return v.ServicePorts }).(GetNsxtAlbVirtualServiceServicePortArrayOutput)
}

// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
func (o LookupNsxtAlbVirtualServiceResultOutput) Vdc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbVirtualServiceResult) string { return v.Vdc }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbVirtualServiceResultOutput) VirtualIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbVirtualServiceResult) string { return v.VirtualIpAddress }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNsxtAlbVirtualServiceResultOutput{})
}