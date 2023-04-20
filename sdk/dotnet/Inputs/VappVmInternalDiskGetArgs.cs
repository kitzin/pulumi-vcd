// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class VappVmInternalDiskGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("busNumber")]
        public Input<int>? BusNumber { get; set; }

        [Input("busType")]
        public Input<string>? BusType { get; set; }

        [Input("diskId")]
        public Input<string>? DiskId { get; set; }

        [Input("iops")]
        public Input<int>? Iops { get; set; }

        [Input("sizeInMb")]
        public Input<int>? SizeInMb { get; set; }

        [Input("storageProfile")]
        public Input<string>? StorageProfile { get; set; }

        [Input("thinProvisioned")]
        public Input<bool>? ThinProvisioned { get; set; }

        [Input("unitNumber")]
        public Input<int>? UnitNumber { get; set; }

        public VappVmInternalDiskGetArgs()
        {
        }
        public static new VappVmInternalDiskGetArgs Empty => new VappVmInternalDiskGetArgs();
    }
}
