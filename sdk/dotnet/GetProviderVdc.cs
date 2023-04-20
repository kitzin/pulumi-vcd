// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetProviderVdc
    {
        public static Task<GetProviderVdcResult> InvokeAsync(GetProviderVdcArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProviderVdcResult>("vcd:index/getProviderVdc:getProviderVdc", args ?? new GetProviderVdcArgs(), options.WithDefaults());

        public static Output<GetProviderVdcResult> Invoke(GetProviderVdcInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetProviderVdcResult>("vcd:index/getProviderVdc:getProviderVdc", args ?? new GetProviderVdcInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProviderVdcArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetProviderVdcArgs()
        {
        }
        public static new GetProviderVdcArgs Empty => new GetProviderVdcArgs();
    }

    public sealed class GetProviderVdcInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetProviderVdcInvokeArgs()
        {
        }
        public static new GetProviderVdcInvokeArgs Empty => new GetProviderVdcInvokeArgs();
    }


    [OutputType]
    public sealed class GetProviderVdcResult
    {
        public readonly ImmutableArray<string> Capabilities;
        public readonly ImmutableArray<Outputs.GetProviderVdcComputeCapacityResult> ComputeCapacities;
        public readonly string ComputeProviderScope;
        public readonly string Description;
        public readonly ImmutableArray<string> ExternalNetworkIds;
        public readonly string HighestSupportedHardwareVersion;
        public readonly ImmutableArray<string> HostIds;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool IsEnabled;
        public readonly ImmutableDictionary<string, object> Metadata;
        public readonly ImmutableArray<Outputs.GetProviderVdcMetadataEntryResult> MetadataEntries;
        public readonly string Name;
        public readonly ImmutableArray<string> NetworkPoolIds;
        public readonly string NsxtManagerId;
        public readonly ImmutableArray<string> ResourcePoolIds;
        public readonly int Status;
        public readonly ImmutableArray<string> StorageContainerIds;
        public readonly ImmutableArray<string> StorageProfileIds;
        public readonly string UniversalNetworkPoolId;
        public readonly string VcenterId;

        [OutputConstructor]
        private GetProviderVdcResult(
            ImmutableArray<string> capabilities,

            ImmutableArray<Outputs.GetProviderVdcComputeCapacityResult> computeCapacities,

            string computeProviderScope,

            string description,

            ImmutableArray<string> externalNetworkIds,

            string highestSupportedHardwareVersion,

            ImmutableArray<string> hostIds,

            string id,

            bool isEnabled,

            ImmutableDictionary<string, object> metadata,

            ImmutableArray<Outputs.GetProviderVdcMetadataEntryResult> metadataEntries,

            string name,

            ImmutableArray<string> networkPoolIds,

            string nsxtManagerId,

            ImmutableArray<string> resourcePoolIds,

            int status,

            ImmutableArray<string> storageContainerIds,

            ImmutableArray<string> storageProfileIds,

            string universalNetworkPoolId,

            string vcenterId)
        {
            Capabilities = capabilities;
            ComputeCapacities = computeCapacities;
            ComputeProviderScope = computeProviderScope;
            Description = description;
            ExternalNetworkIds = externalNetworkIds;
            HighestSupportedHardwareVersion = highestSupportedHardwareVersion;
            HostIds = hostIds;
            Id = id;
            IsEnabled = isEnabled;
            Metadata = metadata;
            MetadataEntries = metadataEntries;
            Name = name;
            NetworkPoolIds = networkPoolIds;
            NsxtManagerId = nsxtManagerId;
            ResourcePoolIds = resourcePoolIds;
            Status = status;
            StorageContainerIds = storageContainerIds;
            StorageProfileIds = storageProfileIds;
            UniversalNetworkPoolId = universalNetworkPoolId;
            VcenterId = vcenterId;
        }
    }
}