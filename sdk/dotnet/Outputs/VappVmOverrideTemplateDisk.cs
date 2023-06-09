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
    public sealed class VappVmOverrideTemplateDisk
    {
        public readonly int BusNumber;
        public readonly string BusType;
        public readonly int? Iops;
        public readonly int SizeInMb;
        public readonly string? StorageProfile;
        public readonly int UnitNumber;

        [OutputConstructor]
        private VappVmOverrideTemplateDisk(
            int busNumber,

            string busType,

            int? iops,

            int sizeInMb,

            string? storageProfile,

            int unitNumber)
        {
            BusNumber = busNumber;
            BusType = busType;
            Iops = iops;
            SizeInMb = sizeInMb;
            StorageProfile = storageProfile;
            UnitNumber = unitNumber;
        }
    }
}
