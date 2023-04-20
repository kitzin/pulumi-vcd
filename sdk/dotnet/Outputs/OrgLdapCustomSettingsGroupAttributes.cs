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
    public sealed class OrgLdapCustomSettingsGroupAttributes
    {
        public readonly string? GroupBackLinkIdentifier;
        public readonly string GroupMembershipIdentifier;
        public readonly string Membership;
        public readonly string Name;
        public readonly string ObjectClass;
        public readonly string UniqueIdentifier;

        [OutputConstructor]
        private OrgLdapCustomSettingsGroupAttributes(
            string? groupBackLinkIdentifier,

            string groupMembershipIdentifier,

            string membership,

            string name,

            string objectClass,

            string uniqueIdentifier)
        {
            GroupBackLinkIdentifier = groupBackLinkIdentifier;
            GroupMembershipIdentifier = groupMembershipIdentifier;
            Membership = membership;
            Name = name;
            ObjectClass = objectClass;
            UniqueIdentifier = uniqueIdentifier;
        }
    }
}
