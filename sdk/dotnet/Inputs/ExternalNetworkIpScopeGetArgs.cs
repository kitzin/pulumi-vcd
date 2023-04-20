// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class ExternalNetworkIpScopeGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("dns1")]
        public Input<string>? Dns1 { get; set; }

        [Input("dns2")]
        public Input<string>? Dns2 { get; set; }

        [Input("dnsSuffix")]
        public Input<string>? DnsSuffix { get; set; }

        [Input("gateway", required: true)]
        public Input<string> Gateway { get; set; } = null!;

        [Input("netmask", required: true)]
        public Input<string> Netmask { get; set; } = null!;

        [Input("staticIpPools")]
        private InputList<Inputs.ExternalNetworkIpScopeStaticIpPoolGetArgs>? _staticIpPools;
        public InputList<Inputs.ExternalNetworkIpScopeStaticIpPoolGetArgs> StaticIpPools
        {
            get => _staticIpPools ?? (_staticIpPools = new InputList<Inputs.ExternalNetworkIpScopeStaticIpPoolGetArgs>());
            set => _staticIpPools = value;
        }

        public ExternalNetworkIpScopeGetArgs()
        {
        }
        public static new ExternalNetworkIpScopeGetArgs Empty => new ExternalNetworkIpScopeGetArgs();
    }
}