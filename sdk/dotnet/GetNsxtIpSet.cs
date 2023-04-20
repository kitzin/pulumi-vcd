// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetNsxtIpSet
    {
        public static Task<GetNsxtIpSetResult> InvokeAsync(GetNsxtIpSetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNsxtIpSetResult>("vcd:index/getNsxtIpSet:getNsxtIpSet", args ?? new GetNsxtIpSetArgs(), options.WithDefaults());

        public static Output<GetNsxtIpSetResult> Invoke(GetNsxtIpSetInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNsxtIpSetResult>("vcd:index/getNsxtIpSet:getNsxtIpSet", args ?? new GetNsxtIpSetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNsxtIpSetArgs : global::Pulumi.InvokeArgs
    {
        [Input("edgeGatewayId", required: true)]
        public string EdgeGatewayId { get; set; } = null!;

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("org")]
        public string? Org { get; set; }

        [Input("vdc")]
        public string? Vdc { get; set; }

        public GetNsxtIpSetArgs()
        {
        }
        public static new GetNsxtIpSetArgs Empty => new GetNsxtIpSetArgs();
    }

    public sealed class GetNsxtIpSetInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("edgeGatewayId", required: true)]
        public Input<string> EdgeGatewayId { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public GetNsxtIpSetInvokeArgs()
        {
        }
        public static new GetNsxtIpSetInvokeArgs Empty => new GetNsxtIpSetInvokeArgs();
    }


    [OutputType]
    public sealed class GetNsxtIpSetResult
    {
        public readonly string Description;
        public readonly string EdgeGatewayId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> IpAddresses;
        public readonly string Name;
        public readonly string? Org;
        public readonly string OwnerId;
        public readonly string? Vdc;

        [OutputConstructor]
        private GetNsxtIpSetResult(
            string description,

            string edgeGatewayId,

            string id,

            ImmutableArray<string> ipAddresses,

            string name,

            string? org,

            string ownerId,

            string? vdc)
        {
            Description = description;
            EdgeGatewayId = edgeGatewayId;
            Id = id;
            IpAddresses = ipAddresses;
            Name = name;
            Org = org;
            OwnerId = ownerId;
            Vdc = vdc;
        }
    }
}
