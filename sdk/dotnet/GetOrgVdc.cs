// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetOrgVdc
    {
        public static Task<GetOrgVdcResult> InvokeAsync(GetOrgVdcArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOrgVdcResult>("vcd:index/getOrgVdc:getOrgVdc", args ?? new GetOrgVdcArgs(), options.WithDefaults());

        public static Output<GetOrgVdcResult> Invoke(GetOrgVdcInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetOrgVdcResult>("vcd:index/getOrgVdc:getOrgVdc", args ?? new GetOrgVdcInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrgVdcArgs : global::Pulumi.InvokeArgs
    {
        [Input("metadataEntries")]
        private List<Inputs.GetOrgVdcMetadataEntryArgs>? _metadataEntries;
        public List<Inputs.GetOrgVdcMetadataEntryArgs> MetadataEntries
        {
            get => _metadataEntries ?? (_metadataEntries = new List<Inputs.GetOrgVdcMetadataEntryArgs>());
            set => _metadataEntries = value;
        }

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("org")]
        public string? Org { get; set; }

        public GetOrgVdcArgs()
        {
        }
        public static new GetOrgVdcArgs Empty => new GetOrgVdcArgs();
    }

    public sealed class GetOrgVdcInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("metadataEntries")]
        private InputList<Inputs.GetOrgVdcMetadataEntryInputArgs>? _metadataEntries;
        public InputList<Inputs.GetOrgVdcMetadataEntryInputArgs> MetadataEntries
        {
            get => _metadataEntries ?? (_metadataEntries = new InputList<Inputs.GetOrgVdcMetadataEntryInputArgs>());
            set => _metadataEntries = value;
        }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("org")]
        public Input<string>? Org { get; set; }

        public GetOrgVdcInvokeArgs()
        {
        }
        public static new GetOrgVdcInvokeArgs Empty => new GetOrgVdcInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrgVdcResult
    {
        public readonly string AllocationModel;
        public readonly bool AllowOverCommit;
        public readonly ImmutableArray<Outputs.GetOrgVdcComputeCapacityResult> ComputeCapacities;
        public readonly double CpuGuaranteed;
        public readonly int CpuSpeed;
        public readonly string DefaultComputePolicyId;
        public readonly string DefaultVmSizingPolicyId;
        public readonly string Description;
        public readonly string EdgeClusterId;
        public readonly bool Elasticity;
        public readonly bool EnableFastProvisioning;
        public readonly bool EnableThinProvisioning;
        public readonly bool EnableVmDiscovery;
        public readonly bool Enabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool IncludeVmMemoryOverhead;
        public readonly double MemoryGuaranteed;
        public readonly ImmutableDictionary<string, object> Metadata;
        public readonly ImmutableArray<Outputs.GetOrgVdcMetadataEntryResult> MetadataEntries;
        public readonly string Name;
        public readonly string NetworkPoolName;
        public readonly int NetworkQuota;
        public readonly int NicQuota;
        public readonly string? Org;
        public readonly string ProviderVdcName;
        public readonly ImmutableArray<Outputs.GetOrgVdcStorageProfileResult> StorageProfiles;
        public readonly ImmutableArray<string> VmPlacementPolicyIds;
        public readonly int VmQuota;
        public readonly ImmutableArray<string> VmSizingPolicyIds;

        [OutputConstructor]
        private GetOrgVdcResult(
            string allocationModel,

            bool allowOverCommit,

            ImmutableArray<Outputs.GetOrgVdcComputeCapacityResult> computeCapacities,

            double cpuGuaranteed,

            int cpuSpeed,

            string defaultComputePolicyId,

            string defaultVmSizingPolicyId,

            string description,

            string edgeClusterId,

            bool elasticity,

            bool enableFastProvisioning,

            bool enableThinProvisioning,

            bool enableVmDiscovery,

            bool enabled,

            string id,

            bool includeVmMemoryOverhead,

            double memoryGuaranteed,

            ImmutableDictionary<string, object> metadata,

            ImmutableArray<Outputs.GetOrgVdcMetadataEntryResult> metadataEntries,

            string name,

            string networkPoolName,

            int networkQuota,

            int nicQuota,

            string? org,

            string providerVdcName,

            ImmutableArray<Outputs.GetOrgVdcStorageProfileResult> storageProfiles,

            ImmutableArray<string> vmPlacementPolicyIds,

            int vmQuota,

            ImmutableArray<string> vmSizingPolicyIds)
        {
            AllocationModel = allocationModel;
            AllowOverCommit = allowOverCommit;
            ComputeCapacities = computeCapacities;
            CpuGuaranteed = cpuGuaranteed;
            CpuSpeed = cpuSpeed;
            DefaultComputePolicyId = defaultComputePolicyId;
            DefaultVmSizingPolicyId = defaultVmSizingPolicyId;
            Description = description;
            EdgeClusterId = edgeClusterId;
            Elasticity = elasticity;
            EnableFastProvisioning = enableFastProvisioning;
            EnableThinProvisioning = enableThinProvisioning;
            EnableVmDiscovery = enableVmDiscovery;
            Enabled = enabled;
            Id = id;
            IncludeVmMemoryOverhead = includeVmMemoryOverhead;
            MemoryGuaranteed = memoryGuaranteed;
            Metadata = metadata;
            MetadataEntries = metadataEntries;
            Name = name;
            NetworkPoolName = networkPoolName;
            NetworkQuota = networkQuota;
            NicQuota = nicQuota;
            Org = org;
            ProviderVdcName = providerVdcName;
            StorageProfiles = storageProfiles;
            VmPlacementPolicyIds = vmPlacementPolicyIds;
            VmQuota = vmQuota;
            VmSizingPolicyIds = vmSizingPolicyIds;
        }
    }
}