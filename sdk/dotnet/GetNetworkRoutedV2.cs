// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetNetworkRoutedV2
    {
        public static Task<GetNetworkRoutedV2Result> InvokeAsync(GetNetworkRoutedV2Args? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNetworkRoutedV2Result>("vcd:index/getNetworkRoutedV2:getNetworkRoutedV2", args ?? new GetNetworkRoutedV2Args(), options.WithDefaults());

        public static Output<GetNetworkRoutedV2Result> Invoke(GetNetworkRoutedV2InvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNetworkRoutedV2Result>("vcd:index/getNetworkRoutedV2:getNetworkRoutedV2", args ?? new GetNetworkRoutedV2InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkRoutedV2Args : global::Pulumi.InvokeArgs
    {
        [Input("edgeGatewayId")]
        public string? EdgeGatewayId { get; set; }

        [Input("filter")]
        public Inputs.GetNetworkRoutedV2FilterArgs? Filter { get; set; }

        [Input("name")]
        public string? Name { get; set; }

        [Input("org")]
        public string? Org { get; set; }

        [Input("vdc")]
        public string? Vdc { get; set; }

        public GetNetworkRoutedV2Args()
        {
        }
        public static new GetNetworkRoutedV2Args Empty => new GetNetworkRoutedV2Args();
    }

    public sealed class GetNetworkRoutedV2InvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("edgeGatewayId")]
        public Input<string>? EdgeGatewayId { get; set; }

        [Input("filter")]
        public Input<Inputs.GetNetworkRoutedV2FilterInputArgs>? Filter { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public GetNetworkRoutedV2InvokeArgs()
        {
        }
        public static new GetNetworkRoutedV2InvokeArgs Empty => new GetNetworkRoutedV2InvokeArgs();
    }


    [OutputType]
    public sealed class GetNetworkRoutedV2Result
    {
        public readonly string Description;
        public readonly string Dns1;
        public readonly string Dns2;
        public readonly string DnsSuffix;
        public readonly string EdgeGatewayId;
        public readonly Outputs.GetNetworkRoutedV2FilterResult? Filter;
        public readonly string Gateway;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InterfaceType;
        public readonly ImmutableDictionary<string, object> Metadata;
        public readonly ImmutableArray<Outputs.GetNetworkRoutedV2MetadataEntryResult> MetadataEntries;
        public readonly string? Name;
        public readonly string? Org;
        public readonly string OwnerId;
        public readonly int PrefixLength;
        public readonly ImmutableArray<Outputs.GetNetworkRoutedV2StaticIpPoolResult> StaticIpPools;
        public readonly string? Vdc;

        [OutputConstructor]
        private GetNetworkRoutedV2Result(
            string description,

            string dns1,

            string dns2,

            string dnsSuffix,

            string edgeGatewayId,

            Outputs.GetNetworkRoutedV2FilterResult? filter,

            string gateway,

            string id,

            string interfaceType,

            ImmutableDictionary<string, object> metadata,

            ImmutableArray<Outputs.GetNetworkRoutedV2MetadataEntryResult> metadataEntries,

            string? name,

            string? org,

            string ownerId,

            int prefixLength,

            ImmutableArray<Outputs.GetNetworkRoutedV2StaticIpPoolResult> staticIpPools,

            string? vdc)
        {
            Description = description;
            Dns1 = dns1;
            Dns2 = dns2;
            DnsSuffix = dnsSuffix;
            EdgeGatewayId = edgeGatewayId;
            Filter = filter;
            Gateway = gateway;
            Id = id;
            InterfaceType = interfaceType;
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
