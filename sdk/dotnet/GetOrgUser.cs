// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetOrgUser
    {
        public static Task<GetOrgUserResult> InvokeAsync(GetOrgUserArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetOrgUserResult>("vcd:index/getOrgUser:getOrgUser", args ?? new GetOrgUserArgs(), options.WithDefaults());

        public static Output<GetOrgUserResult> Invoke(GetOrgUserInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetOrgUserResult>("vcd:index/getOrgUser:getOrgUser", args ?? new GetOrgUserInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrgUserArgs : global::Pulumi.InvokeArgs
    {
        [Input("name")]
        public string? Name { get; set; }

        [Input("org")]
        public string? Org { get; set; }

        [Input("userId")]
        public string? UserId { get; set; }

        public GetOrgUserArgs()
        {
        }
        public static new GetOrgUserArgs Empty => new GetOrgUserArgs();
    }

    public sealed class GetOrgUserInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public GetOrgUserInvokeArgs()
        {
        }
        public static new GetOrgUserInvokeArgs Empty => new GetOrgUserInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrgUserResult
    {
        public readonly int DeployedVmQuota;
        public readonly string Description;
        public readonly string EmailAddress;
        public readonly bool Enabled;
        public readonly string FullName;
        public readonly ImmutableArray<string> GroupNames;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstantMessaging;
        public readonly bool IsExternal;
        public readonly bool IsGroupRole;
        public readonly bool IsLocked;
        public readonly string? Name;
        public readonly string? Org;
        public readonly string ProviderType;
        public readonly string Role;
        public readonly int StoredVmQuota;
        public readonly string Telephone;
        public readonly string? UserId;

        [OutputConstructor]
        private GetOrgUserResult(
            int deployedVmQuota,

            string description,

            string emailAddress,

            bool enabled,

            string fullName,

            ImmutableArray<string> groupNames,

            string id,

            string instantMessaging,

            bool isExternal,

            bool isGroupRole,

            bool isLocked,

            string? name,

            string? org,

            string providerType,

            string role,

            int storedVmQuota,

            string telephone,

            string? userId)
        {
            DeployedVmQuota = deployedVmQuota;
            Description = description;
            EmailAddress = emailAddress;
            Enabled = enabled;
            FullName = fullName;
            GroupNames = groupNames;
            Id = id;
            InstantMessaging = instantMessaging;
            IsExternal = isExternal;
            IsGroupRole = isGroupRole;
            IsLocked = isLocked;
            Name = name;
            Org = org;
            ProviderType = providerType;
            Role = role;
            StoredVmQuota = storedVmQuota;
            Telephone = telephone;
            UserId = userId;
        }
    }
}
