// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NsxvDhcpRelayRelayAgentGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("gatewayIpAddress")]
        public Input<string>? GatewayIpAddress { get; set; }

        [Input("networkName", required: true)]
        public Input<string> NetworkName { get; set; } = null!;

        public NsxvDhcpRelayRelayAgentGetArgs()
        {
        }
        public static new NsxvDhcpRelayRelayAgentGetArgs Empty => new NsxvDhcpRelayRelayAgentGetArgs();
    }
}