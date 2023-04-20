// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetVappNetwork
    {
        public static Task<GetVappNetworkResult> InvokeAsync(GetVappNetworkArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVappNetworkResult>("vcd:index/getVappNetwork:getVappNetwork", args ?? new GetVappNetworkArgs(), options.WithDefaults());

        public static Output<GetVappNetworkResult> Invoke(GetVappNetworkInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetVappNetworkResult>("vcd:index/getVappNetwork:getVappNetwork", args ?? new GetVappNetworkInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetVappNetworkArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("org")]
        public string? Org { get; set; }

        [Input("vappName", required: true)]
        public string VappName { get; set; } = null!;

        [Input("vdc")]
        public string? Vdc { get; set; }

        public GetVappNetworkArgs()
        {
        }
        public static new GetVappNetworkArgs Empty => new GetVappNetworkArgs();
    }

    public sealed class GetVappNetworkInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("vappName", required: true)]
        public Input<string> VappName { get; set; } = null!;

        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public GetVappNetworkInvokeArgs()
        {
        }
        public static new GetVappNetworkInvokeArgs Empty => new GetVappNetworkInvokeArgs();
    }


    [OutputType]
    public sealed class GetVappNetworkResult
    {
        public readonly string Description;
        public readonly ImmutableArray<Outputs.GetVappNetworkDhcpPoolResult> DhcpPools;
        public readonly string Dns1;
        public readonly string Dns2;
        public readonly string DnsSuffix;
        public readonly string Gateway;
        public readonly bool GuestVlanAllowed;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string Netmask;
        public readonly string? Org;
        public readonly string OrgNetworkName;
        public readonly bool RetainIpMacEnabled;
        public readonly ImmutableArray<Outputs.GetVappNetworkStaticIpPoolResult> StaticIpPools;
        public readonly string VappName;
        public readonly string? Vdc;

        [OutputConstructor]
        private GetVappNetworkResult(
            string description,

            ImmutableArray<Outputs.GetVappNetworkDhcpPoolResult> dhcpPools,

            string dns1,

            string dns2,

            string dnsSuffix,

            string gateway,

            bool guestVlanAllowed,

            string id,

            string name,

            string netmask,

            string? org,

            string orgNetworkName,

            bool retainIpMacEnabled,

            ImmutableArray<Outputs.GetVappNetworkStaticIpPoolResult> staticIpPools,

            string vappName,

            string? vdc)
        {
            Description = description;
            DhcpPools = dhcpPools;
            Dns1 = dns1;
            Dns2 = dns2;
            DnsSuffix = dnsSuffix;
            Gateway = gateway;
            GuestVlanAllowed = guestVlanAllowed;
            Id = id;
            Name = name;
            Netmask = netmask;
            Org = org;
            OrgNetworkName = orgNetworkName;
            RetainIpMacEnabled = retainIpMacEnabled;
            StaticIpPools = staticIpPools;
            VappName = vappName;
            Vdc = vdc;
        }
    }
}