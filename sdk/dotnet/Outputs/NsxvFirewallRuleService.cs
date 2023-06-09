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
    public sealed class NsxvFirewallRuleService
    {
        public readonly string? Port;
        public readonly string Protocol;
        public readonly string? SourcePort;

        [OutputConstructor]
        private NsxvFirewallRuleService(
            string? port,

            string protocol,

            string? sourcePort)
        {
            Port = port;
            Protocol = protocol;
            SourcePort = sourcePort;
        }
    }
}
