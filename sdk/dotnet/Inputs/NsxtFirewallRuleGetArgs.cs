// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NsxtFirewallRuleGetArgs : global::Pulumi.ResourceArgs
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

        [Input("destinationIds")]
        private InputList<string>? _destinationIds;
        public InputList<string> DestinationIds
        {
            get => _destinationIds ?? (_destinationIds = new InputList<string>());
            set => _destinationIds = value;
        }

        [Input("direction", required: true)]
        public Input<string> Direction { get; set; } = null!;

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("ipProtocol", required: true)]
        public Input<string> IpProtocol { get; set; } = null!;

        [Input("logging")]
        public Input<bool>? Logging { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("sourceIds")]
        private InputList<string>? _sourceIds;
        public InputList<string> SourceIds
        {
            get => _sourceIds ?? (_sourceIds = new InputList<string>());
            set => _sourceIds = value;
        }

        public NsxtFirewallRuleGetArgs()
        {
        }
        public static new NsxtFirewallRuleGetArgs Empty => new NsxtFirewallRuleGetArgs();
    }
}
