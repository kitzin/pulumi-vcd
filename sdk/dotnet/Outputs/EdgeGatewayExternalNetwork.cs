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
    public sealed class EdgeGatewayExternalNetwork
    {
        public readonly bool? EnableRateLimit;
        public readonly double? IncomingRateLimit;
        public readonly string Name;
        public readonly double? OutgoingRateLimit;
        public readonly ImmutableArray<Outputs.EdgeGatewayExternalNetworkSubnet> Subnets;

        [OutputConstructor]
        private EdgeGatewayExternalNetwork(
            bool? enableRateLimit,

            double? incomingRateLimit,

            string name,

            double? outgoingRateLimit,

            ImmutableArray<Outputs.EdgeGatewayExternalNetworkSubnet> subnets)
        {
            EnableRateLimit = enableRateLimit;
            IncomingRateLimit = incomingRateLimit;
            Name = name;
            OutgoingRateLimit = outgoingRateLimit;
            Subnets = subnets;
        }
    }
}
