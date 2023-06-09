// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class VappVmCustomizationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("adminPassword")]
        public Input<string>? AdminPassword { get; set; }

        [Input("allowLocalAdminPassword")]
        public Input<bool>? AllowLocalAdminPassword { get; set; }

        [Input("autoGeneratePassword")]
        public Input<bool>? AutoGeneratePassword { get; set; }

        [Input("changeSid")]
        public Input<bool>? ChangeSid { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("force")]
        public Input<bool>? Force { get; set; }

        [Input("initscript")]
        public Input<string>? Initscript { get; set; }

        [Input("joinDomain")]
        public Input<bool>? JoinDomain { get; set; }

        [Input("joinDomainAccountOu")]
        public Input<string>? JoinDomainAccountOu { get; set; }

        [Input("joinDomainName")]
        public Input<string>? JoinDomainName { get; set; }

        [Input("joinDomainPassword")]
        public Input<string>? JoinDomainPassword { get; set; }

        [Input("joinDomainUser")]
        public Input<string>? JoinDomainUser { get; set; }

        [Input("joinOrgDomain")]
        public Input<bool>? JoinOrgDomain { get; set; }

        [Input("mustChangePasswordOnFirstLogin")]
        public Input<bool>? MustChangePasswordOnFirstLogin { get; set; }

        [Input("numberOfAutoLogons")]
        public Input<int>? NumberOfAutoLogons { get; set; }

        public VappVmCustomizationGetArgs()
        {
        }
        public static new VappVmCustomizationGetArgs Empty => new VappVmCustomizationGetArgs();
    }
}
