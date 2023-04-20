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
    public sealed class GetProviderVdcComputeCapacityResult
    {
        public readonly ImmutableArray<Outputs.GetProviderVdcComputeCapacityCpusResult> Cpus;
        public readonly bool IsElastic;
        public readonly bool IsHa;
        public readonly ImmutableArray<Outputs.GetProviderVdcComputeCapacityMemoryResult> Memories;

        [OutputConstructor]
        private GetProviderVdcComputeCapacityResult(
            ImmutableArray<Outputs.GetProviderVdcComputeCapacityCpusResult> cpus,

            bool isElastic,

            bool isHa,

            ImmutableArray<Outputs.GetProviderVdcComputeCapacityMemoryResult> memories)
        {
            Cpus = cpus;
            IsElastic = isElastic;
            IsHa = isHa;
            Memories = memories;
        }
    }
}
