// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class VmNetworkGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("adapterType")]
        public Input<string>? AdapterType { get; set; }

        [Input("connected")]
        public Input<bool>? Connected { get; set; }

        [Input("ip")]
        public Input<string>? Ip { get; set; }

        [Input("ipAllocationMode")]
        public Input<string>? IpAllocationMode { get; set; }

        [Input("isPrimary")]
        public Input<bool>? IsPrimary { get; set; }

        [Input("mac")]
        public Input<string>? Mac { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public VmNetworkGetArgs()
        {
        }
        public static new VmNetworkGetArgs Empty => new VmNetworkGetArgs();
    }
}
