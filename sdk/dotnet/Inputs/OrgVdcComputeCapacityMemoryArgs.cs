// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class OrgVdcComputeCapacityMemoryArgs : global::Pulumi.ResourceArgs
    {
        [Input("allocated")]
        public Input<int>? Allocated { get; set; }

        [Input("limit")]
        public Input<int>? Limit { get; set; }

        [Input("reserved")]
        public Input<int>? Reserved { get; set; }

        [Input("used")]
        public Input<int>? Used { get; set; }

        public OrgVdcComputeCapacityMemoryArgs()
        {
        }
        public static new OrgVdcComputeCapacityMemoryArgs Empty => new OrgVdcComputeCapacityMemoryArgs();
    }
}