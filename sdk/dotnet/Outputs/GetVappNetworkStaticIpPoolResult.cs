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
    public sealed class GetVappNetworkStaticIpPoolResult
    {
        public readonly string EndAddress;
        public readonly string StartAddress;

        [OutputConstructor]
        private GetVappNetworkStaticIpPoolResult(
            string endAddress,

            string startAddress)
        {
            EndAddress = endAddress;
            StartAddress = startAddress;
        }
    }
}
