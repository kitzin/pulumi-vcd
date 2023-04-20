// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetVm
    {
        public static Task<GetVmResult> InvokeAsync(GetVmArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVmResult>("vcd:index/getVm:getVm", args ?? new GetVmArgs(), options.WithDefaults());

        public static Output<GetVmResult> Invoke(GetVmInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetVmResult>("vcd:index/getVm:getVm", args ?? new GetVmInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVmArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("networkDhcpWaitSeconds")]
        public int? NetworkDhcpWaitSeconds { get; set; }

        [Input("org")]
        public string? Org { get; set; }

        [Input("placementPolicyId")]
        public string? PlacementPolicyId { get; set; }

        [Input("sizingPolicyId")]
        public string? SizingPolicyId { get; set; }

        [Input("vappName")]
        public string? VappName { get; set; }

        [Input("vdc")]
        public string? Vdc { get; set; }

        public GetVmArgs()
        {
        }
        public static new GetVmArgs Empty => new GetVmArgs();
    }

    public sealed class GetVmInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("networkDhcpWaitSeconds")]
        public Input<int>? NetworkDhcpWaitSeconds { get; set; }

        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("placementPolicyId")]
        public Input<string>? PlacementPolicyId { get; set; }

        [Input("sizingPolicyId")]
        public Input<string>? SizingPolicyId { get; set; }

        [Input("vappName")]
        public Input<string>? VappName { get; set; }

        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public GetVmInvokeArgs()
        {
        }
        public static new GetVmInvokeArgs Empty => new GetVmInvokeArgs();
    }


    [OutputType]
    public sealed class GetVmResult
    {
        public readonly string ComputerName;
        public readonly int CpuCores;
        public readonly bool CpuHotAddEnabled;
        public readonly int CpuLimit;
        public readonly string CpuPriority;
        public readonly int CpuReservation;
        public readonly int CpuShares;
        public readonly int Cpus;
        public readonly ImmutableArray<Outputs.GetVmCustomizationResult> Customizations;
        public readonly string Description;
        public readonly ImmutableArray<Outputs.GetVmDiskResult> Disks;
        public readonly bool ExposeHardwareVirtualization;
        public readonly ImmutableDictionary<string, object> GuestProperties;
        public readonly string HardwareVersion;
        public readonly string Href;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetVmInternalDiskResult> InternalDisks;
        public readonly int Memory;
        public readonly bool MemoryHotAddEnabled;
        public readonly int MemoryLimit;
        public readonly string MemoryPriority;
        public readonly int MemoryReservation;
        public readonly int MemoryShares;
        public readonly ImmutableDictionary<string, object> Metadata;
        public readonly ImmutableArray<Outputs.GetVmMetadataEntryResult> MetadataEntries;
        public readonly string Name;
        public readonly int? NetworkDhcpWaitSeconds;
        public readonly ImmutableArray<Outputs.GetVmNetworkResult> Networks;
        public readonly string? Org;
        public readonly string OsType;
        public readonly string PlacementPolicyId;
        public readonly string SizingPolicyId;
        public readonly int Status;
        public readonly string StatusText;
        public readonly string StorageProfile;
        public readonly string VappName;
        public readonly string? Vdc;
        public readonly string VmType;

        [OutputConstructor]
        private GetVmResult(
            string computerName,

            int cpuCores,

            bool cpuHotAddEnabled,

            int cpuLimit,

            string cpuPriority,

            int cpuReservation,

            int cpuShares,

            int cpus,

            ImmutableArray<Outputs.GetVmCustomizationResult> customizations,

            string description,

            ImmutableArray<Outputs.GetVmDiskResult> disks,

            bool exposeHardwareVirtualization,

            ImmutableDictionary<string, object> guestProperties,

            string hardwareVersion,

            string href,

            string id,

            ImmutableArray<Outputs.GetVmInternalDiskResult> internalDisks,

            int memory,

            bool memoryHotAddEnabled,

            int memoryLimit,

            string memoryPriority,

            int memoryReservation,

            int memoryShares,

            ImmutableDictionary<string, object> metadata,

            ImmutableArray<Outputs.GetVmMetadataEntryResult> metadataEntries,

            string name,

            int? networkDhcpWaitSeconds,

            ImmutableArray<Outputs.GetVmNetworkResult> networks,

            string? org,

            string osType,

            string placementPolicyId,

            string sizingPolicyId,

            int status,

            string statusText,

            string storageProfile,

            string vappName,

            string? vdc,

            string vmType)
        {
            ComputerName = computerName;
            CpuCores = cpuCores;
            CpuHotAddEnabled = cpuHotAddEnabled;
            CpuLimit = cpuLimit;
            CpuPriority = cpuPriority;
            CpuReservation = cpuReservation;
            CpuShares = cpuShares;
            Cpus = cpus;
            Customizations = customizations;
            Description = description;
            Disks = disks;
            ExposeHardwareVirtualization = exposeHardwareVirtualization;
            GuestProperties = guestProperties;
            HardwareVersion = hardwareVersion;
            Href = href;
            Id = id;
            InternalDisks = internalDisks;
            Memory = memory;
            MemoryHotAddEnabled = memoryHotAddEnabled;
            MemoryLimit = memoryLimit;
            MemoryPriority = memoryPriority;
            MemoryReservation = memoryReservation;
            MemoryShares = memoryShares;
            Metadata = metadata;
            MetadataEntries = metadataEntries;
            Name = name;
            NetworkDhcpWaitSeconds = networkDhcpWaitSeconds;
            Networks = networks;
            Org = org;
            OsType = osType;
            PlacementPolicyId = placementPolicyId;
            SizingPolicyId = sizingPolicyId;
            Status = status;
            StatusText = statusText;
            StorageProfile = storageProfile;
            VappName = vappName;
            Vdc = vdc;
            VmType = vmType;
        }
    }
}
