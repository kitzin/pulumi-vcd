// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNsxtNetworkDhcp(ctx *pulumi.Context, args *LookupNsxtNetworkDhcpArgs, opts ...pulumi.InvokeOption) (*LookupNsxtNetworkDhcpResult, error) {
	var rv LookupNsxtNetworkDhcpResult
	err := ctx.Invoke("vcd:index/getNsxtNetworkDhcp:getNsxtNetworkDhcp", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxtNetworkDhcp.
type LookupNsxtNetworkDhcpArgs struct {
	DnsServers   []string `pulumi:"dnsServers"`
	Org          *string  `pulumi:"org"`
	OrgNetworkId string   `pulumi:"orgNetworkId"`
	// Deprecated: Org network will be looked up based on 'org_network_id' field
	Vdc *string `pulumi:"vdc"`
}

// A collection of values returned by getNsxtNetworkDhcp.
type LookupNsxtNetworkDhcpResult struct {
	DnsServers []string `pulumi:"dnsServers"`
	// The provider-assigned unique ID for this managed resource.
	Id                string                   `pulumi:"id"`
	LeaseTime         int                      `pulumi:"leaseTime"`
	ListenerIpAddress string                   `pulumi:"listenerIpAddress"`
	Mode              string                   `pulumi:"mode"`
	Org               *string                  `pulumi:"org"`
	OrgNetworkId      string                   `pulumi:"orgNetworkId"`
	Pools             []GetNsxtNetworkDhcpPool `pulumi:"pools"`
	// Deprecated: Org network will be looked up based on 'org_network_id' field
	Vdc *string `pulumi:"vdc"`
}

func LookupNsxtNetworkDhcpOutput(ctx *pulumi.Context, args LookupNsxtNetworkDhcpOutputArgs, opts ...pulumi.InvokeOption) LookupNsxtNetworkDhcpResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNsxtNetworkDhcpResult, error) {
			args := v.(LookupNsxtNetworkDhcpArgs)
			r, err := LookupNsxtNetworkDhcp(ctx, &args, opts...)
			var s LookupNsxtNetworkDhcpResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNsxtNetworkDhcpResultOutput)
}

// A collection of arguments for invoking getNsxtNetworkDhcp.
type LookupNsxtNetworkDhcpOutputArgs struct {
	DnsServers   pulumi.StringArrayInput `pulumi:"dnsServers"`
	Org          pulumi.StringPtrInput   `pulumi:"org"`
	OrgNetworkId pulumi.StringInput      `pulumi:"orgNetworkId"`
	// Deprecated: Org network will be looked up based on 'org_network_id' field
	Vdc pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupNsxtNetworkDhcpOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtNetworkDhcpArgs)(nil)).Elem()
}

// A collection of values returned by getNsxtNetworkDhcp.
type LookupNsxtNetworkDhcpResultOutput struct{ *pulumi.OutputState }

func (LookupNsxtNetworkDhcpResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtNetworkDhcpResult)(nil)).Elem()
}

func (o LookupNsxtNetworkDhcpResultOutput) ToLookupNsxtNetworkDhcpResultOutput() LookupNsxtNetworkDhcpResultOutput {
	return o
}

func (o LookupNsxtNetworkDhcpResultOutput) ToLookupNsxtNetworkDhcpResultOutputWithContext(ctx context.Context) LookupNsxtNetworkDhcpResultOutput {
	return o
}

func (o LookupNsxtNetworkDhcpResultOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNsxtNetworkDhcpResult) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNsxtNetworkDhcpResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNetworkDhcpResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNsxtNetworkDhcpResultOutput) LeaseTime() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNsxtNetworkDhcpResult) int { return v.LeaseTime }).(pulumi.IntOutput)
}

func (o LookupNsxtNetworkDhcpResultOutput) ListenerIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNetworkDhcpResult) string { return v.ListenerIpAddress }).(pulumi.StringOutput)
}

func (o LookupNsxtNetworkDhcpResultOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNetworkDhcpResult) string { return v.Mode }).(pulumi.StringOutput)
}

func (o LookupNsxtNetworkDhcpResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtNetworkDhcpResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupNsxtNetworkDhcpResultOutput) OrgNetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNetworkDhcpResult) string { return v.OrgNetworkId }).(pulumi.StringOutput)
}

func (o LookupNsxtNetworkDhcpResultOutput) Pools() GetNsxtNetworkDhcpPoolArrayOutput {
	return o.ApplyT(func(v LookupNsxtNetworkDhcpResult) []GetNsxtNetworkDhcpPool { return v.Pools }).(GetNsxtNetworkDhcpPoolArrayOutput)
}

// Deprecated: Org network will be looked up based on 'org_network_id' field
func (o LookupNsxtNetworkDhcpResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtNetworkDhcpResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNsxtNetworkDhcpResultOutput{})
}
