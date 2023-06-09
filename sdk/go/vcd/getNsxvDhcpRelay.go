// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNsxvDhcpRelay(ctx *pulumi.Context, args *LookupNsxvDhcpRelayArgs, opts ...pulumi.InvokeOption) (*LookupNsxvDhcpRelayResult, error) {
	var rv LookupNsxvDhcpRelayResult
	err := ctx.Invoke("vcd:index/getNsxvDhcpRelay:getNsxvDhcpRelay", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxvDhcpRelay.
type LookupNsxvDhcpRelayArgs struct {
	EdgeGateway string  `pulumi:"edgeGateway"`
	Org         *string `pulumi:"org"`
	Vdc         *string `pulumi:"vdc"`
}

// A collection of values returned by getNsxvDhcpRelay.
type LookupNsxvDhcpRelayResult struct {
	DomainNames []string `pulumi:"domainNames"`
	EdgeGateway string   `pulumi:"edgeGateway"`
	// The provider-assigned unique ID for this managed resource.
	Id          string                       `pulumi:"id"`
	IpAddresses []string                     `pulumi:"ipAddresses"`
	IpSets      []string                     `pulumi:"ipSets"`
	Org         *string                      `pulumi:"org"`
	RelayAgents []GetNsxvDhcpRelayRelayAgent `pulumi:"relayAgents"`
	Vdc         *string                      `pulumi:"vdc"`
}

func LookupNsxvDhcpRelayOutput(ctx *pulumi.Context, args LookupNsxvDhcpRelayOutputArgs, opts ...pulumi.InvokeOption) LookupNsxvDhcpRelayResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNsxvDhcpRelayResult, error) {
			args := v.(LookupNsxvDhcpRelayArgs)
			r, err := LookupNsxvDhcpRelay(ctx, &args, opts...)
			var s LookupNsxvDhcpRelayResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNsxvDhcpRelayResultOutput)
}

// A collection of arguments for invoking getNsxvDhcpRelay.
type LookupNsxvDhcpRelayOutputArgs struct {
	EdgeGateway pulumi.StringInput    `pulumi:"edgeGateway"`
	Org         pulumi.StringPtrInput `pulumi:"org"`
	Vdc         pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupNsxvDhcpRelayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxvDhcpRelayArgs)(nil)).Elem()
}

// A collection of values returned by getNsxvDhcpRelay.
type LookupNsxvDhcpRelayResultOutput struct{ *pulumi.OutputState }

func (LookupNsxvDhcpRelayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxvDhcpRelayResult)(nil)).Elem()
}

func (o LookupNsxvDhcpRelayResultOutput) ToLookupNsxvDhcpRelayResultOutput() LookupNsxvDhcpRelayResultOutput {
	return o
}

func (o LookupNsxvDhcpRelayResultOutput) ToLookupNsxvDhcpRelayResultOutputWithContext(ctx context.Context) LookupNsxvDhcpRelayResultOutput {
	return o
}

func (o LookupNsxvDhcpRelayResultOutput) DomainNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNsxvDhcpRelayResult) []string { return v.DomainNames }).(pulumi.StringArrayOutput)
}

func (o LookupNsxvDhcpRelayResultOutput) EdgeGateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvDhcpRelayResult) string { return v.EdgeGateway }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNsxvDhcpRelayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvDhcpRelayResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNsxvDhcpRelayResultOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNsxvDhcpRelayResult) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

func (o LookupNsxvDhcpRelayResultOutput) IpSets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNsxvDhcpRelayResult) []string { return v.IpSets }).(pulumi.StringArrayOutput)
}

func (o LookupNsxvDhcpRelayResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxvDhcpRelayResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupNsxvDhcpRelayResultOutput) RelayAgents() GetNsxvDhcpRelayRelayAgentArrayOutput {
	return o.ApplyT(func(v LookupNsxvDhcpRelayResult) []GetNsxvDhcpRelayRelayAgent { return v.RelayAgents }).(GetNsxvDhcpRelayRelayAgentArrayOutput)
}

func (o LookupNsxvDhcpRelayResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxvDhcpRelayResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNsxvDhcpRelayResultOutput{})
}
