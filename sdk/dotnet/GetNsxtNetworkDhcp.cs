// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetNsxtNetworkDhcp
    {
        public static Task<GetNsxtNetworkDhcpResult> InvokeAsync(GetNsxtNetworkDhcpArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNsxtNetworkDhcpResult>("vcd:index/getNsxtNetworkDhcp:getNsxtNetworkDhcp", args ?? new GetNsxtNetworkDhcpArgs(), options.WithDefaults());

        public static Output<GetNsxtNetworkDhcpResult> Invoke(GetNsxtNetworkDhcpInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetNsxtNetworkDhcpResult>("vcd:index/getNsxtNetworkDhcp:getNsxtNetworkDhcp", args ?? new GetNsxtNetworkDhcpInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNsxtNetworkDhcpArgs : global::Pulumi.InvokeArgs
    {
        [Input("dnsServers")]
        private List<string>? _dnsServers;
        public List<string> DnsServers
        {
            get => _dnsServers ?? (_dnsServers = new List<string>());
            set => _dnsServers = value;
        }

        [Input("org")]
        public string? Org { get; set; }

        [Input("orgNetworkId", required: true)]
        public string OrgNetworkId { get; set; } = null!;

        [Input("vdc")]
        public string? Vdc { get; set; }

        public GetNsxtNetworkDhcpArgs()
        {
        }
        public static new GetNsxtNetworkDhcpArgs Empty => new GetNsxtNetworkDhcpArgs();
    }

    public sealed class GetNsxtNetworkDhcpInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("dnsServers")]
        private InputList<string>? _dnsServers;
        public InputList<string> DnsServers
        {
            get => _dnsServers ?? (_dnsServers = new InputList<string>());
            set => _dnsServers = value;
        }

        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("orgNetworkId", required: true)]
        public Input<string> OrgNetworkId { get; set; } = null!;

        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public GetNsxtNetworkDhcpInvokeArgs()
        {
        }
        public static new GetNsxtNetworkDhcpInvokeArgs Empty => new GetNsxtNetworkDhcpInvokeArgs();
    }


    [OutputType]
    public sealed class GetNsxtNetworkDhcpResult
    {
        public readonly ImmutableArray<string> DnsServers;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int LeaseTime;
        public readonly string ListenerIpAddress;
        public readonly string Mode;
        public readonly string? Org;
        public readonly string OrgNetworkId;
        public readonly ImmutableArray<Outputs.GetNsxtNetworkDhcpPoolResult> Pools;
        public readonly string? Vdc;

        [OutputConstructor]
        private GetNsxtNetworkDhcpResult(
            ImmutableArray<string> dnsServers,

            string id,

            int leaseTime,

            string listenerIpAddress,

            string mode,

            string? org,

            string orgNetworkId,

            ImmutableArray<Outputs.GetNsxtNetworkDhcpPoolResult> pools,

            string? vdc)
        {
            DnsServers = dnsServers;
            Id = id;
            LeaseTime = leaseTime;
            ListenerIpAddress = listenerIpAddress;
            Mode = mode;
            Org = org;
            OrgNetworkId = orgNetworkId;
            Pools = pools;
            Vdc = vdc;
        }
    }
}
