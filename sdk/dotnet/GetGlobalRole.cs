// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetGlobalRole
    {
        public static Task<GetGlobalRoleResult> InvokeAsync(GetGlobalRoleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetGlobalRoleResult>("vcd:index/getGlobalRole:getGlobalRole", args ?? new GetGlobalRoleArgs(), options.WithDefaults());

        public static Output<GetGlobalRoleResult> Invoke(GetGlobalRoleInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetGlobalRoleResult>("vcd:index/getGlobalRole:getGlobalRole", args ?? new GetGlobalRoleInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetGlobalRoleArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetGlobalRoleArgs()
        {
        }
        public static new GetGlobalRoleArgs Empty => new GetGlobalRoleArgs();
    }

    public sealed class GetGlobalRoleInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetGlobalRoleInvokeArgs()
        {
        }
        public static new GetGlobalRoleInvokeArgs Empty => new GetGlobalRoleInvokeArgs();
    }


    [OutputType]
    public sealed class GetGlobalRoleResult
    {
        public readonly string BundleKey;
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly bool PublishToAllTenants;
        public readonly bool ReadOnly;
        public readonly ImmutableArray<string> Rights;
        public readonly ImmutableArray<string> Tenants;

        [OutputConstructor]
        private GetGlobalRoleResult(
            string bundleKey,

            string description,

            string id,

            string name,

            bool publishToAllTenants,

            bool readOnly,

            ImmutableArray<string> rights,

            ImmutableArray<string> tenants)
        {
            BundleKey = bundleKey;
            Description = description;
            Id = id;
            Name = name;
            PublishToAllTenants = publishToAllTenants;
            ReadOnly = readOnly;
            Rights = rights;
            Tenants = tenants;
        }
    }
}