// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NsxtDynamicSecurityGroupCriteriaArgs : global::Pulumi.ResourceArgs
    {
        [Input("rules")]
        private InputList<Inputs.NsxtDynamicSecurityGroupCriteriaRuleArgs>? _rules;
        public InputList<Inputs.NsxtDynamicSecurityGroupCriteriaRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.NsxtDynamicSecurityGroupCriteriaRuleArgs>());
            set => _rules = value;
        }

        public NsxtDynamicSecurityGroupCriteriaArgs()
        {
        }
        public static new NsxtDynamicSecurityGroupCriteriaArgs Empty => new NsxtDynamicSecurityGroupCriteriaArgs();
    }
}