// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class VdcGroupParticipatingOrgVdcGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("faultDomainTag")]
        public Input<string>? FaultDomainTag { get; set; }

        [Input("isRemoteOrg")]
        public Input<bool>? IsRemoteOrg { get; set; }

        [Input("networkProviderScope")]
        public Input<string>? NetworkProviderScope { get; set; }

        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        [Input("orgName")]
        public Input<string>? OrgName { get; set; }

        [Input("siteId")]
        public Input<string>? SiteId { get; set; }

        [Input("siteName")]
        public Input<string>? SiteName { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("vdcId")]
        public Input<string>? VdcId { get; set; }

        [Input("vdcName")]
        public Input<string>? VdcName { get; set; }

        public VdcGroupParticipatingOrgVdcGetArgs()
        {
        }
        public static new VdcGroupParticipatingOrgVdcGetArgs Empty => new VdcGroupParticipatingOrgVdcGetArgs();
    }
}
