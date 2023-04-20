// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class VmDiskArgs : global::Pulumi.ResourceArgs
    {
        [Input("busNumber", required: true)]
        public Input<string> BusNumber { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("sizeInMb")]
        public Input<int>? SizeInMb { get; set; }

        [Input("unitNumber", required: true)]
        public Input<string> UnitNumber { get; set; } = null!;

        public VmDiskArgs()
        {
        }
        public static new VmDiskArgs Empty => new VmDiskArgs();
    }
}
