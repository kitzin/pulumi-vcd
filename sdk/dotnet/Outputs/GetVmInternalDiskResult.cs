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
    public sealed class GetVmInternalDiskResult
    {
        public readonly int BusNumber;
        public readonly string BusType;
        public readonly string DiskId;
        public readonly int Iops;
        public readonly int SizeInMb;
        public readonly string StorageProfile;
        public readonly bool ThinProvisioned;
        public readonly int UnitNumber;

        [OutputConstructor]
        private GetVmInternalDiskResult(
            int busNumber,

            string busType,

            string diskId,

            int iops,

            int sizeInMb,

            string storageProfile,

            bool thinProvisioned,

            int unitNumber)
        {
            BusNumber = busNumber;
            BusType = busType;
            DiskId = diskId;
            Iops = iops;
            SizeInMb = sizeInMb;
            StorageProfile = storageProfile;
            ThinProvisioned = thinProvisioned;
            UnitNumber = unitNumber;
        }
    }
}
