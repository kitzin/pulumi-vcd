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
    public sealed class EdgeGatewayVpnLocalSubnet
    {
        public readonly string LocalSubnetGateway;
        public readonly string LocalSubnetMask;
        public readonly string LocalSubnetName;

        [OutputConstructor]
        private EdgeGatewayVpnLocalSubnet(
            string localSubnetGateway,

            string localSubnetMask,

            string localSubnetName)
        {
            LocalSubnetGateway = localSubnetGateway;
            LocalSubnetMask = localSubnetMask;
            LocalSubnetName = localSubnetName;
        }
    }
}
