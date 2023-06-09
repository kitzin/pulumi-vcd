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
    public sealed class NsxtAlbVirtualServiceServicePort
    {
        public readonly int? EndPort;
        public readonly bool? SslEnabled;
        public readonly int StartPort;
        public readonly string Type;

        [OutputConstructor]
        private NsxtAlbVirtualServiceServicePort(
            int? endPort,

            bool? sslEnabled,

            int startPort,

            string type)
        {
            EndPort = endPort;
            SslEnabled = sslEnabled;
            StartPort = startPort;
            Type = type;
        }
    }
}
