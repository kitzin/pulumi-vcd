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
    public sealed class OrgLdapCustomSettingsUserAttributes
    {
        public readonly string DisplayName;
        public readonly string Email;
        public readonly string GivenName;
        public readonly string? GroupBackLinkIdentifier;
        public readonly string GroupMembershipIdentifier;
        public readonly string ObjectClass;
        public readonly string Surname;
        public readonly string Telephone;
        public readonly string UniqueIdentifier;
        public readonly string Username;

        [OutputConstructor]
        private OrgLdapCustomSettingsUserAttributes(
            string displayName,

            string email,

            string givenName,

            string? groupBackLinkIdentifier,

            string groupMembershipIdentifier,

            string objectClass,

            string surname,

            string telephone,

            string uniqueIdentifier,

            string username)
        {
            DisplayName = displayName;
            Email = email;
            GivenName = givenName;
            GroupBackLinkIdentifier = groupBackLinkIdentifier;
            GroupMembershipIdentifier = groupMembershipIdentifier;
            ObjectClass = objectClass;
            Surname = surname;
            Telephone = telephone;
            UniqueIdentifier = uniqueIdentifier;
            Username = username;
        }
    }
}
