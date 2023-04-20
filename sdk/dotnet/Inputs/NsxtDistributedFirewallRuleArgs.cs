// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NsxtDistributedFirewallRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        [Input("appPortProfileIds")]
        private InputList<string>? _appPortProfileIds;
        public InputList<string> AppPortProfileIds
        {
            get => _appPortProfileIds ?? (_appPortProfileIds = new InputList<string>());
            set => _appPortProfileIds = value;
        }

        [Input("comment")]
        public Input<string>? Comment { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("destinationGroupsExcluded")]
        public Input<bool>? DestinationGroupsExcluded { get; set; }

        [Input("destinationIds")]
        private InputList<string>? _destinationIds;
        public InputList<string> DestinationIds
        {
            get => _destinationIds ?? (_destinationIds = new InputList<string>());
            set => _destinationIds = value;
        }

        [Input("direction")]
        public Input<string>? Direction { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("ipProtocol")]
        public Input<string>? IpProtocol { get; set; }

        [Input("logging")]
        public Input<bool>? Logging { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("networkContextProfileIds")]
        private InputList<string>? _networkContextProfileIds;
        public InputList<string> NetworkContextProfileIds
        {
            get => _networkContextProfileIds ?? (_networkContextProfileIds = new InputList<string>());
            set => _networkContextProfileIds = value;
        }

        [Input("sourceGroupsExcluded")]
        public Input<bool>? SourceGroupsExcluded { get; set; }

        [Input("sourceIds")]
        private InputList<string>? _sourceIds;
        public InputList<string> SourceIds
        {
            get => _sourceIds ?? (_sourceIds = new InputList<string>());
            set => _sourceIds = value;
        }

        public NsxtDistributedFirewallRuleArgs()
        {
        }
        public static new NsxtDistributedFirewallRuleArgs Empty => new NsxtDistributedFirewallRuleArgs();
    }
}
