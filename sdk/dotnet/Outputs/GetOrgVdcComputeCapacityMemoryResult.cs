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
    public sealed class GetOrgVdcComputeCapacityMemoryResult
    {
        public readonly int Allocated;
        public readonly int Limit;
        public readonly int Reserved;
        public readonly int Used;

        [OutputConstructor]
        private GetOrgVdcComputeCapacityMemoryResult(
            int allocated,

            int limit,

            int reserved,

            int used)
        {
            Allocated = allocated;
            Limit = limit;
            Reserved = reserved;
            Used = used;
        }
    }
}
