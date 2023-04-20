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
    public sealed class VappLease
    {
        public readonly int RuntimeLeaseInSec;
        public readonly int StorageLeaseInSec;

        [OutputConstructor]
        private VappLease(
            int runtimeLeaseInSec,

            int storageLeaseInSec)
        {
            RuntimeLeaseInSec = runtimeLeaseInSec;
            StorageLeaseInSec = storageLeaseInSec;
        }
    }
}