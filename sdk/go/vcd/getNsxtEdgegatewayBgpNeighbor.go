// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNsxtEdgegatewayBgpNeighbor(ctx *pulumi.Context, args *LookupNsxtEdgegatewayBgpNeighborArgs, opts ...pulumi.InvokeOption) (*LookupNsxtEdgegatewayBgpNeighborResult, error) {
	var rv LookupNsxtEdgegatewayBgpNeighborResult
	err := ctx.Invoke("vcd:index/getNsxtEdgegatewayBgpNeighbor:getNsxtEdgegatewayBgpNeighbor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxtEdgegatewayBgpNeighbor.
type LookupNsxtEdgegatewayBgpNeighborArgs struct {
	EdgeGatewayId string  `pulumi:"edgeGatewayId"`
	IpAddress     string  `pulumi:"ipAddress"`
	Org           *string `pulumi:"org"`
}

// A collection of values returned by getNsxtEdgegatewayBgpNeighbor.
type LookupNsxtEdgegatewayBgpNeighborResult struct {
	AllowAsIn           bool   `pulumi:"allowAsIn"`
	BfdDeadMultiple     int    `pulumi:"bfdDeadMultiple"`
	BfdEnabled          bool   `pulumi:"bfdEnabled"`
	BfdInterval         int    `pulumi:"bfdInterval"`
	EdgeGatewayId       string `pulumi:"edgeGatewayId"`
	GracefulRestartMode string `pulumi:"gracefulRestartMode"`
	HoldDownTimer       int    `pulumi:"holdDownTimer"`
	// The provider-assigned unique ID for this managed resource.
	Id                      string  `pulumi:"id"`
	InFilterIpPrefixListId  string  `pulumi:"inFilterIpPrefixListId"`
	IpAddress               string  `pulumi:"ipAddress"`
	KeepAliveTimer          int     `pulumi:"keepAliveTimer"`
	Org                     *string `pulumi:"org"`
	OutFilterIpPrefixListId string  `pulumi:"outFilterIpPrefixListId"`
	Password                string  `pulumi:"password"`
	RemoteAsNumber          string  `pulumi:"remoteAsNumber"`
	RouteFiltering          string  `pulumi:"routeFiltering"`
}

func LookupNsxtEdgegatewayBgpNeighborOutput(ctx *pulumi.Context, args LookupNsxtEdgegatewayBgpNeighborOutputArgs, opts ...pulumi.InvokeOption) LookupNsxtEdgegatewayBgpNeighborResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNsxtEdgegatewayBgpNeighborResult, error) {
			args := v.(LookupNsxtEdgegatewayBgpNeighborArgs)
			r, err := LookupNsxtEdgegatewayBgpNeighbor(ctx, &args, opts...)
			var s LookupNsxtEdgegatewayBgpNeighborResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNsxtEdgegatewayBgpNeighborResultOutput)
}

// A collection of arguments for invoking getNsxtEdgegatewayBgpNeighbor.
type LookupNsxtEdgegatewayBgpNeighborOutputArgs struct {
	EdgeGatewayId pulumi.StringInput    `pulumi:"edgeGatewayId"`
	IpAddress     pulumi.StringInput    `pulumi:"ipAddress"`
	Org           pulumi.StringPtrInput `pulumi:"org"`
}

func (LookupNsxtEdgegatewayBgpNeighborOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtEdgegatewayBgpNeighborArgs)(nil)).Elem()
}

// A collection of values returned by getNsxtEdgegatewayBgpNeighbor.
type LookupNsxtEdgegatewayBgpNeighborResultOutput struct{ *pulumi.OutputState }

func (LookupNsxtEdgegatewayBgpNeighborResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtEdgegatewayBgpNeighborResult)(nil)).Elem()
}

func (o LookupNsxtEdgegatewayBgpNeighborResultOutput) ToLookupNsxtEdgegatewayBgpNeighborResultOutput() LookupNsxtEdgegatewayBgpNeighborResultOutput {
	return o
}

func (o LookupNsxtEdgegatewayBgpNeighborResultOutput) ToLookupNsxtEdgegatewayBgpNeighborResultOutputWithContext(ctx context.Context) LookupNsxtEdgegatewayBgpNeighborResultOutput {
	return o
}

func (o LookupNsxtEdgegatewayBgpNeighborResultOutput) AllowAsIn() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayBgpNeighborResult) bool { return v.AllowAsIn }).(pulumi.BoolOutput)
}

func (o LookupNsxtEdgegatewayBgpNeighborResultOutput) BfdDeadMultiple() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayBgpNeighborResult) int { return v.BfdDeadMultiple }).(pulumi.IntOutput)
}

func (o LookupNsxtEdgegatewayBgpNeighborResultOutput) BfdEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayBgpNeighborResult) bool { return v.BfdEnabled }).(pulumi.BoolOutput)
}

func (o LookupNsxtEdgegatewayBgpNeighborResultOutput) BfdInterval() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayBgpNeighborResult) int { return v.BfdInterval }).(pulumi.IntOutput)
}

func (o LookupNsxtEdgegatewayBgpNeighborResultOutput) EdgeGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayBgpNeighborResult) string { return v.EdgeGatewayId }).(pulumi.StringOutput)
}

func (o LookupNsxtEdgegatewayBgpNeighborResultOutput) GracefulRestartMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayBgpNeighborResult) string { return v.GracefulRestartMode }).(pulumi.StringOutput)
}

func (o LookupNsxtEdgegatewayBgpNeighborResultOutput) HoldDownTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayBgpNeighborResult) int { return v.HoldDownTimer }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNsxtEdgegatewayBgpNeighborResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayBgpNeighborResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNsxtEdgegatewayBgpNeighborResultOutput) InFilterIpPrefixListId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayBgpNeighborResult) string { return v.InFilterIpPrefixListId }).(pulumi.StringOutput)
}

func (o LookupNsxtEdgegatewayBgpNeighborResultOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayBgpNeighborResult) string { return v.IpAddress }).(pulumi.StringOutput)
}

func (o LookupNsxtEdgegatewayBgpNeighborResultOutput) KeepAliveTimer() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayBgpNeighborResult) int { return v.KeepAliveTimer }).(pulumi.IntOutput)
}

func (o LookupNsxtEdgegatewayBgpNeighborResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayBgpNeighborResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupNsxtEdgegatewayBgpNeighborResultOutput) OutFilterIpPrefixListId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayBgpNeighborResult) string { return v.OutFilterIpPrefixListId }).(pulumi.StringOutput)
}

func (o LookupNsxtEdgegatewayBgpNeighborResultOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayBgpNeighborResult) string { return v.Password }).(pulumi.StringOutput)
}

func (o LookupNsxtEdgegatewayBgpNeighborResultOutput) RemoteAsNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayBgpNeighborResult) string { return v.RemoteAsNumber }).(pulumi.StringOutput)
}

func (o LookupNsxtEdgegatewayBgpNeighborResultOutput) RouteFiltering() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayBgpNeighborResult) string { return v.RouteFiltering }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNsxtEdgegatewayBgpNeighborResultOutput{})
}