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
    public sealed class GetOrgVappLeaseResult
    {
        public readonly bool DeleteOnStorageLeaseExpiration;
        public readonly int MaximumRuntimeLeaseInSec;
        public readonly int MaximumStorageLeaseInSec;
        public readonly bool PowerOffOnRuntimeLeaseExpiration;

        [OutputConstructor]
        private GetOrgVappLeaseResult(
            bool deleteOnStorageLeaseExpiration,

            int maximumRuntimeLeaseInSec,

            int maximumStorageLeaseInSec,

            bool powerOffOnRuntimeLeaseExpiration)
        {
            DeleteOnStorageLeaseExpiration = deleteOnStorageLeaseExpiration;
            MaximumRuntimeLeaseInSec = maximumRuntimeLeaseInSec;
            MaximumStorageLeaseInSec = maximumStorageLeaseInSec;
            PowerOffOnRuntimeLeaseExpiration = powerOffOnRuntimeLeaseExpiration;
        }
    }
}
