// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVappVm(ctx *pulumi.Context, args *LookupVappVmArgs, opts ...pulumi.InvokeOption) (*LookupVappVmResult, error) {
	var rv LookupVappVmResult
	err := ctx.Invoke("vcd:index/getVappVm:getVappVm", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVappVm.
type LookupVappVmArgs struct {
	Name                   string  `pulumi:"name"`
	NetworkDhcpWaitSeconds *int    `pulumi:"networkDhcpWaitSeconds"`
	Org                    *string `pulumi:"org"`
	PlacementPolicyId      *string `pulumi:"placementPolicyId"`
	SizingPolicyId         *string `pulumi:"sizingPolicyId"`
	VappName               string  `pulumi:"vappName"`
	Vdc                    *string `pulumi:"vdc"`
}

// A collection of values returned by getVappVm.
type LookupVappVmResult struct {
	ComputerName                 string                   `pulumi:"computerName"`
	CpuCores                     int                      `pulumi:"cpuCores"`
	CpuHotAddEnabled             bool                     `pulumi:"cpuHotAddEnabled"`
	CpuLimit                     int                      `pulumi:"cpuLimit"`
	CpuPriority                  string                   `pulumi:"cpuPriority"`
	CpuReservation               int                      `pulumi:"cpuReservation"`
	CpuShares                    int                      `pulumi:"cpuShares"`
	Cpus                         int                      `pulumi:"cpus"`
	Customizations               []GetVappVmCustomization `pulumi:"customizations"`
	Description                  string                   `pulumi:"description"`
	Disks                        []GetVappVmDisk          `pulumi:"disks"`
	ExposeHardwareVirtualization bool                     `pulumi:"exposeHardwareVirtualization"`
	GuestProperties              map[string]interface{}   `pulumi:"guestProperties"`
	HardwareVersion              string                   `pulumi:"hardwareVersion"`
	Href                         string                   `pulumi:"href"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string                  `pulumi:"id"`
	InternalDisks       []GetVappVmInternalDisk `pulumi:"internalDisks"`
	Memory              int                     `pulumi:"memory"`
	MemoryHotAddEnabled bool                    `pulumi:"memoryHotAddEnabled"`
	MemoryLimit         int                     `pulumi:"memoryLimit"`
	MemoryPriority      string                  `pulumi:"memoryPriority"`
	MemoryReservation   int                     `pulumi:"memoryReservation"`
	MemoryShares        int                     `pulumi:"memoryShares"`
	// Deprecated: Use metadata_entry instead
	Metadata               map[string]interface{}   `pulumi:"metadata"`
	MetadataEntries        []GetVappVmMetadataEntry `pulumi:"metadataEntries"`
	Name                   string                   `pulumi:"name"`
	NetworkDhcpWaitSeconds *int                     `pulumi:"networkDhcpWaitSeconds"`
	Networks               []GetVappVmNetwork       `pulumi:"networks"`
	Org                    *string                  `pulumi:"org"`
	OsType                 string                   `pulumi:"osType"`
	PlacementPolicyId      string                   `pulumi:"placementPolicyId"`
	SizingPolicyId         string                   `pulumi:"sizingPolicyId"`
	Status                 int                      `pulumi:"status"`
	StatusText             string                   `pulumi:"statusText"`
	StorageProfile         string                   `pulumi:"storageProfile"`
	VappName               string                   `pulumi:"vappName"`
	Vdc                    *string                  `pulumi:"vdc"`
	VmType                 string                   `pulumi:"vmType"`
}

func LookupVappVmOutput(ctx *pulumi.Context, args LookupVappVmOutputArgs, opts ...pulumi.InvokeOption) LookupVappVmResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVappVmResult, error) {
			args := v.(LookupVappVmArgs)
			r, err := LookupVappVm(ctx, &args, opts...)
			var s LookupVappVmResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVappVmResultOutput)
}

// A collection of arguments for invoking getVappVm.
type LookupVappVmOutputArgs struct {
	Name                   pulumi.StringInput    `pulumi:"name"`
	NetworkDhcpWaitSeconds pulumi.IntPtrInput    `pulumi:"networkDhcpWaitSeconds"`
	Org                    pulumi.StringPtrInput `pulumi:"org"`
	PlacementPolicyId      pulumi.StringPtrInput `pulumi:"placementPolicyId"`
	SizingPolicyId         pulumi.StringPtrInput `pulumi:"sizingPolicyId"`
	VappName               pulumi.StringInput    `pulumi:"vappName"`
	Vdc                    pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupVappVmOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVappVmArgs)(nil)).Elem()
}

// A collection of values returned by getVappVm.
type LookupVappVmResultOutput struct{ *pulumi.OutputState }

func (LookupVappVmResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVappVmResult)(nil)).Elem()
}

func (o LookupVappVmResultOutput) ToLookupVappVmResultOutput() LookupVappVmResultOutput {
	return o
}

func (o LookupVappVmResultOutput) ToLookupVappVmResultOutputWithContext(ctx context.Context) LookupVappVmResultOutput {
	return o
}

func (o LookupVappVmResultOutput) ComputerName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappVmResult) string { return v.ComputerName }).(pulumi.StringOutput)
}

func (o LookupVappVmResultOutput) CpuCores() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVappVmResult) int { return v.CpuCores }).(pulumi.IntOutput)
}

func (o LookupVappVmResultOutput) CpuHotAddEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVappVmResult) bool { return v.CpuHotAddEnabled }).(pulumi.BoolOutput)
}

func (o LookupVappVmResultOutput) CpuLimit() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVappVmResult) int { return v.CpuLimit }).(pulumi.IntOutput)
}

func (o LookupVappVmResultOutput) CpuPriority() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappVmResult) string { return v.CpuPriority }).(pulumi.StringOutput)
}

func (o LookupVappVmResultOutput) CpuReservation() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVappVmResult) int { return v.CpuReservation }).(pulumi.IntOutput)
}

func (o LookupVappVmResultOutput) CpuShares() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVappVmResult) int { return v.CpuShares }).(pulumi.IntOutput)
}

func (o LookupVappVmResultOutput) Cpus() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVappVmResult) int { return v.Cpus }).(pulumi.IntOutput)
}

func (o LookupVappVmResultOutput) Customizations() GetVappVmCustomizationArrayOutput {
	return o.ApplyT(func(v LookupVappVmResult) []GetVappVmCustomization { return v.Customizations }).(GetVappVmCustomizationArrayOutput)
}

func (o LookupVappVmResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappVmResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupVappVmResultOutput) Disks() GetVappVmDiskArrayOutput {
	return o.ApplyT(func(v LookupVappVmResult) []GetVappVmDisk { return v.Disks }).(GetVappVmDiskArrayOutput)
}

func (o LookupVappVmResultOutput) ExposeHardwareVirtualization() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVappVmResult) bool { return v.ExposeHardwareVirtualization }).(pulumi.BoolOutput)
}

func (o LookupVappVmResultOutput) GuestProperties() pulumi.MapOutput {
	return o.ApplyT(func(v LookupVappVmResult) map[string]interface{} { return v.GuestProperties }).(pulumi.MapOutput)
}

func (o LookupVappVmResultOutput) HardwareVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappVmResult) string { return v.HardwareVersion }).(pulumi.StringOutput)
}

func (o LookupVappVmResultOutput) Href() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappVmResult) string { return v.Href }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVappVmResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappVmResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVappVmResultOutput) InternalDisks() GetVappVmInternalDiskArrayOutput {
	return o.ApplyT(func(v LookupVappVmResult) []GetVappVmInternalDisk { return v.InternalDisks }).(GetVappVmInternalDiskArrayOutput)
}

func (o LookupVappVmResultOutput) Memory() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVappVmResult) int { return v.Memory }).(pulumi.IntOutput)
}

func (o LookupVappVmResultOutput) MemoryHotAddEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVappVmResult) bool { return v.MemoryHotAddEnabled }).(pulumi.BoolOutput)
}

func (o LookupVappVmResultOutput) MemoryLimit() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVappVmResult) int { return v.MemoryLimit }).(pulumi.IntOutput)
}

func (o LookupVappVmResultOutput) MemoryPriority() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappVmResult) string { return v.MemoryPriority }).(pulumi.StringOutput)
}

func (o LookupVappVmResultOutput) MemoryReservation() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVappVmResult) int { return v.MemoryReservation }).(pulumi.IntOutput)
}

func (o LookupVappVmResultOutput) MemoryShares() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVappVmResult) int { return v.MemoryShares }).(pulumi.IntOutput)
}

// Deprecated: Use metadata_entry instead
func (o LookupVappVmResultOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v LookupVappVmResult) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

func (o LookupVappVmResultOutput) MetadataEntries() GetVappVmMetadataEntryArrayOutput {
	return o.ApplyT(func(v LookupVappVmResult) []GetVappVmMetadataEntry { return v.MetadataEntries }).(GetVappVmMetadataEntryArrayOutput)
}

func (o LookupVappVmResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappVmResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVappVmResultOutput) NetworkDhcpWaitSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVappVmResult) *int { return v.NetworkDhcpWaitSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupVappVmResultOutput) Networks() GetVappVmNetworkArrayOutput {
	return o.ApplyT(func(v LookupVappVmResult) []GetVappVmNetwork { return v.Networks }).(GetVappVmNetworkArrayOutput)
}

func (o LookupVappVmResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVappVmResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupVappVmResultOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappVmResult) string { return v.OsType }).(pulumi.StringOutput)
}

func (o LookupVappVmResultOutput) PlacementPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappVmResult) string { return v.PlacementPolicyId }).(pulumi.StringOutput)
}

func (o LookupVappVmResultOutput) SizingPolicyId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappVmResult) string { return v.SizingPolicyId }).(pulumi.StringOutput)
}

func (o LookupVappVmResultOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVappVmResult) int { return v.Status }).(pulumi.IntOutput)
}

func (o LookupVappVmResultOutput) StatusText() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappVmResult) string { return v.StatusText }).(pulumi.StringOutput)
}

func (o LookupVappVmResultOutput) StorageProfile() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappVmResult) string { return v.StorageProfile }).(pulumi.StringOutput)
}

func (o LookupVappVmResultOutput) VappName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappVmResult) string { return v.VappName }).(pulumi.StringOutput)
}

func (o LookupVappVmResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVappVmResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func (o LookupVappVmResultOutput) VmType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVappVmResult) string { return v.VmType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVappVmResultOutput{})
}
