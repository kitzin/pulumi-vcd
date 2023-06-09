// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NsxtEdgegatewaySubnetGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("allocatedIps", required: true)]
        private InputList<Inputs.NsxtEdgegatewaySubnetAllocatedIpGetArgs>? _allocatedIps;
        public InputList<Inputs.NsxtEdgegatewaySubnetAllocatedIpGetArgs> AllocatedIps
        {
            get => _allocatedIps ?? (_allocatedIps = new InputList<Inputs.NsxtEdgegatewaySubnetAllocatedIpGetArgs>());
            set => _allocatedIps = value;
        }

        [Input("gateway", required: true)]
        public Input<string> Gateway { get; set; } = null!;

        [Input("prefixLength", required: true)]
        public Input<int> PrefixLength { get; set; } = null!;

        [Input("primaryIp")]
        public Input<string>? PrimaryIp { get; set; }

        public NsxtEdgegatewaySubnetGetArgs()
        {
        }
        public static new NsxtEdgegatewaySubnetGetArgs Empty => new NsxtEdgegatewaySubnetGetArgs();
    }
}
