// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetNetworkIsolatedV2
    {
        public static Task<GetNetworkIsolatedV2Result> InvokeAsync(GetNetworkIsolatedV2Args? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetworkIsolatedV2Result>("vcd:index/getNetworkIsolatedV2:getNetworkIsolatedV2", args ?? new GetNetworkIsolatedV2Args(), options.WithDefaults());

        public static Output<GetNetworkIsolatedV2Result> Invoke(GetNetworkIsolatedV2InvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNetworkIsolatedV2Result>("vcd:index/getNetworkIsolatedV2:getNetworkIsolatedV2", args ?? new GetNetworkIsolatedV2InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkIsolatedV2Args : global::Pulumi.InvokeArgs
    {
        [Input("filter")]
        public Inputs.GetNetworkIsolatedV2FilterArgs? Filter { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("org")]
        public string? Org { get; set; }

        [Input("ownerId")]
        public string? OwnerId { get; set; }

        [Input("vdc")]
        public string? Vdc { get; set; }

        public GetNetworkIsolatedV2Args()
        {
        }
        public static new GetNetworkIsolatedV2Args Empty => new GetNetworkIsolatedV2Args();
    }

    public sealed class GetNetworkIsolatedV2InvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("filter")]
        public Input<Inputs.GetNetworkIsolatedV2FilterInputArgs>? Filter { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public GetNetworkIsolatedV2InvokeArgs()
        {
        }
        public static new GetNetworkIsolatedV2InvokeArgs Empty => new GetNetworkIsolatedV2InvokeArgs();
    }


    [OutputType]
    public sealed class GetNetworkIsolatedV2Result
    {
        public readonly string Description;
        public readonly string Dns1;
        public readonly string Dns2;
        public readonly string DnsSuffix;
        public readonly Outputs.GetNetworkIsolatedV2FilterResult? Filter;
        public readonly string Gateway;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool IsShared;
        public readonly ImmutableDictionary<string, object> Metadata;
        public readonly ImmutableArray<Outputs.GetNetworkIsolatedV2MetadataEntryResult> MetadataEntries;
        public readonly string? Name;
        public readonly string? Org;
        public readonly string OwnerId;
        public readonly int PrefixLength;
        public readonly ImmutableArray<Outputs.GetNetworkIsolatedV2StaticIpPoolResult> StaticIpPools;
        public readonly string Vdc;

        [OutputConstructor]
        private GetNetworkIsolatedV2Result(
            string description,

            string dns1,

            string dns2,

            string dnsSuffix,

            Outputs.GetNetworkIsolatedV2FilterResult? filter,

            string gateway,

            string id,

            bool isShared,

            ImmutableDictionary<string, object> metadata,

            ImmutableArray<Outputs.GetNetworkIsolatedV2MetadataEntryResult> metadataEntries,

            string? name,

            string? org,

            string ownerId,

            int prefixLength,

            ImmutableArray<Outputs.GetNetworkIsolatedV2StaticIpPoolResult> staticIpPools,

            string vdc)
        {
            Description = description;
            Dns1 = dns1;
            Dns2 = dns2;
            DnsSuffix = dnsSuffix;
            Filter = filter;
            Gateway = gateway;
            Id = id;
            IsShared = isShared;
            Metadata = metadata;
            MetadataEntries = metadataEntries;
            Name = name;
            Org = org;
            OwnerId = ownerId;
            PrefixLength = prefixLength;
            StaticIpPools = staticIpPools;
            Vdc = vdc;
        }
    }
}