// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class VappFirewallRulesRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("destinationIp")]
        public Input<string>? DestinationIp { get; set; }

        [Input("destinationPort")]
        public Input<string>? DestinationPort { get; set; }

        [Input("destinationVmId")]
        public Input<string>? DestinationVmId { get; set; }

        [Input("destinationVmIpType")]
        public Input<string>? DestinationVmIpType { get; set; }

        [Input("destinationVmNicId")]
        public Input<int>? DestinationVmNicId { get; set; }

        [Input("enableLogging")]
        public Input<bool>? EnableLogging { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policy")]
        public Input<string>? Policy { get; set; }

        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        [Input("sourcePort")]
        public Input<string>? SourcePort { get; set; }

        [Input("sourceVmId")]
        public Input<string>? SourceVmId { get; set; }

        [Input("sourceVmIpType")]
        public Input<string>? SourceVmIpType { get; set; }

        [Input("sourceVmNicId")]
        public Input<int>? SourceVmNicId { get; set; }

        public VappFirewallRulesRuleArgs()
        {
        }
        public static new VappFirewallRulesRuleArgs Empty => new VappFirewallRulesRuleArgs();
    }
}
