// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class VappNetworkStaticIpPoolArgs : global::Pulumi.ResourceArgs
    {
        [Input("endAddress", required: true)]
        public Input<string> EndAddress { get; set; } = null!;

        [Input("startAddress", required: true)]
        public Input<string> StartAddress { get; set; } = null!;

        public VappNetworkStaticIpPoolArgs()
        {
        }
        public static new VappNetworkStaticIpPoolArgs Empty => new VappNetworkStaticIpPoolArgs();
    }
}