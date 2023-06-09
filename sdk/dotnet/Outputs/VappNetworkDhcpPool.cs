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
    public sealed class VappNetworkDhcpPool
    {
        public readonly int? DefaultLeaseTime;
        public readonly bool? Enabled;
        public readonly string? EndAddress;
        public readonly int? MaxLeaseTime;
        public readonly string StartAddress;

        [OutputConstructor]
        private VappNetworkDhcpPool(
            int? defaultLeaseTime,

            bool? enabled,

            string? endAddress,

            int? maxLeaseTime,

            string startAddress)
        {
            DefaultLeaseTime = defaultLeaseTime;
            Enabled = enabled;
            EndAddress = endAddress;
            MaxLeaseTime = maxLeaseTime;
            StartAddress = startAddress;
        }
    }
}
