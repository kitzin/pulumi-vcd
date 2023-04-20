// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NsxtAlbPoolMemberGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("detailedHealthMessage")]
        public Input<string>? DetailedHealthMessage { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("healthStatus")]
        public Input<string>? HealthStatus { get; set; }

        [Input("ipAddress", required: true)]
        public Input<string> IpAddress { get; set; } = null!;

        [Input("markedDownBies")]
        private InputList<string>? _markedDownBies;
        public InputList<string> MarkedDownBies
        {
            get => _markedDownBies ?? (_markedDownBies = new InputList<string>());
            set => _markedDownBies = value;
        }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("ratio")]
        public Input<int>? Ratio { get; set; }

        public NsxtAlbPoolMemberGetArgs()
        {
        }
        public static new NsxtAlbPoolMemberGetArgs Empty => new NsxtAlbPoolMemberGetArgs();
    }
}