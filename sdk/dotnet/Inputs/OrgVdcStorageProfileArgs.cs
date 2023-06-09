// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class OrgVdcStorageProfileArgs : global::Pulumi.ResourceArgs
    {
        [Input("default", required: true)]
        public Input<bool> Default { get; set; } = null!;

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("limit", required: true)]
        public Input<int> Limit { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("storageUsedInMb")]
        public Input<int>? StorageUsedInMb { get; set; }

        public OrgVdcStorageProfileArgs()
        {
        }
        public static new OrgVdcStorageProfileArgs Empty => new OrgVdcStorageProfileArgs();
    }
}
