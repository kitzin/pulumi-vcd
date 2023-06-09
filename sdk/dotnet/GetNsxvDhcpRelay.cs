// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetNsxvDhcpRelay
    {
        public static Task<GetNsxvDhcpRelayResult> InvokeAsync(GetNsxvDhcpRelayArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNsxvDhcpRelayResult>("vcd:index/getNsxvDhcpRelay:getNsxvDhcpRelay", args ?? new GetNsxvDhcpRelayArgs(), options.WithDefaults());

        public static Output<GetNsxvDhcpRelayResult> Invoke(GetNsxvDhcpRelayInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNsxvDhcpRelayResult>("vcd:index/getNsxvDhcpRelay:getNsxvDhcpRelay", args ?? new GetNsxvDhcpRelayInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNsxvDhcpRelayArgs : global::Pulumi.InvokeArgs
    {
        [Input("edgeGateway", required: true)]
        public string EdgeGateway { get; set; } = null!;

        [Input("org")]
        public string? Org { get; set; }

        [Input("vdc")]
        public string? Vdc { get; set; }

        public GetNsxvDhcpRelayArgs()
        {
        }
        public static new GetNsxvDhcpRelayArgs Empty => new GetNsxvDhcpRelayArgs();
    }

    public sealed class GetNsxvDhcpRelayInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("edgeGateway", required: true)]
        public Input<string> EdgeGateway { get; set; } = null!;

        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public GetNsxvDhcpRelayInvokeArgs()
        {
        }
        public static new GetNsxvDhcpRelayInvokeArgs Empty => new GetNsxvDhcpRelayInvokeArgs();
    }


    [OutputType]
    public sealed class GetNsxvDhcpRelayResult
    {
        public readonly ImmutableArray<string> DomainNames;
        public readonly string EdgeGateway;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> IpAddresses;
        public readonly ImmutableArray<string> IpSets;
        public readonly string? Org;
        public readonly ImmutableArray<Outputs.GetNsxvDhcpRelayRelayAgentResult> RelayAgents;
        public readonly string? Vdc;

        [OutputConstructor]
        private GetNsxvDhcpRelayResult(
            ImmutableArray<string> domainNames,

            string edgeGateway,

            string id,

            ImmutableArray<string> ipAddresses,

            ImmutableArray<string> ipSets,

            string? org,

            ImmutableArray<Outputs.GetNsxvDhcpRelayRelayAgentResult> relayAgents,

            string? vdc)
        {
            DomainNames = domainNames;
            EdgeGateway = edgeGateway;
            Id = id;
            IpAddresses = ipAddresses;
            IpSets = ipSets;
            Org = org;
            RelayAgents = relayAgents;
            Vdc = vdc;
        }
    }
}
