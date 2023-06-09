// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Outputs
{

    [OutputType]
    public sealed class GetVappVmDiskResult
    {
        public readonly string BusNumber;
        public readonly string Name;
        public readonly int SizeInMb;
        public readonly string UnitNumber;

        [OutputConstructor]
        private GetVappVmDiskResult(
            string busNumber,

            string name,

            int sizeInMb,

            string unitNumber)
        {
            BusNumber = busNumber;
            Name = name;
            SizeInMb = sizeInMb;
            UnitNumber = unitNumber;
        }
    }
}
