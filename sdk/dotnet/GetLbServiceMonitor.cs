// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetLbServiceMonitor
    {
        public static Task<GetLbServiceMonitorResult> InvokeAsync(GetLbServiceMonitorArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLbServiceMonitorResult>("vcd:index/getLbServiceMonitor:getLbServiceMonitor", args ?? new GetLbServiceMonitorArgs(), options.WithDefaults());

        public static Output<GetLbServiceMonitorResult> Invoke(GetLbServiceMonitorInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetLbServiceMonitorResult>("vcd:index/getLbServiceMonitor:getLbServiceMonitor", args ?? new GetLbServiceMonitorInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLbServiceMonitorArgs : global::Pulumi.InvokeArgs
    {
        [Input("edgeGateway", required: true)]
        public string EdgeGateway { get; set; } = null!;

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("org")]
        public string? Org { get; set; }

        [Input("vdc")]
        public string? Vdc { get; set; }

        public GetLbServiceMonitorArgs()
        {
        }
        public static new GetLbServiceMonitorArgs Empty => new GetLbServiceMonitorArgs();
    }

    public sealed class GetLbServiceMonitorInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("edgeGateway", required: true)]
        public Input<string> EdgeGateway { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public GetLbServiceMonitorInvokeArgs()
        {
        }
        public static new GetLbServiceMonitorInvokeArgs Empty => new GetLbServiceMonitorInvokeArgs();
    }


    [OutputType]
    public sealed class GetLbServiceMonitorResult
    {
        public readonly string EdgeGateway;
        public readonly string Expected;
        public readonly ImmutableDictionary<string, object> Extension;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int Interval;
        public readonly int MaxRetries;
        public readonly string Method;
        public readonly string Name;
        public readonly string? Org;
        public readonly string Receive;
        public readonly string Send;
        public readonly int Timeout;
        public readonly string Type;
        public readonly string Url;
        public readonly string? Vdc;

        [OutputConstructor]
        private GetLbServiceMonitorResult(
            string edgeGateway,

            string expected,

            ImmutableDictionary<string, object> extension,

            string id,

            int interval,

            int maxRetries,

            string method,

            string name,

            string? org,

            string receive,

            string send,

            int timeout,

            string type,

            string url,

            string? vdc)
        {
            EdgeGateway = edgeGateway;
            Expected = expected;
            Extension = extension;
            Id = id;
            Interval = interval;
            MaxRetries = maxRetries;
            Method = method;
            Name = name;
            Org = org;
            Receive = receive;
            Send = send;
            Timeout = timeout;
            Type = type;
            Url = url;
            Vdc = vdc;
        }
    }
}
