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
    public sealed class GetLbServerPoolMemberResult
    {
        public readonly string Condition;
        public readonly string Id;
        public readonly string IpAddress;
        public readonly int MaxConnections;
        public readonly int MinConnections;
        public readonly int MonitorPort;
        public readonly string Name;
        public readonly int Port;
        public readonly int Weight;

        [OutputConstructor]
        private GetLbServerPoolMemberResult(
            string condition,

            string id,

            string ipAddress,

            int maxConnections,

            int minConnections,

            int monitorPort,

            string name,

            int port,

            int weight)
        {
            Condition = condition;
            Id = id;
            IpAddress = ipAddress;
            MaxConnections = maxConnections;
            MinConnections = minConnections;
            MonitorPort = monitorPort;
            Name = name;
            Port = port;
            Weight = weight;
        }
    }
}
