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
    public sealed class ExternalNetworkV2VsphereNetwork
    {
        public readonly string PortgroupId;
        public readonly string VcenterId;

        [OutputConstructor]
        private ExternalNetworkV2VsphereNetwork(
            string portgroupId,

            string vcenterId)
        {
            PortgroupId = portgroupId;
            VcenterId = vcenterId;
        }
    }
}
