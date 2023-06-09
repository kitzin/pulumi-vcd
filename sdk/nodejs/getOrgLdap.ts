// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getOrgLdap(args: GetOrgLdapArgs, opts?: pulumi.InvokeOptions): Promise<GetOrgLdapResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vcd:index/getOrgLdap:getOrgLdap", {
        "orgId": args.orgId,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrgLdap.
 */
export interface GetOrgLdapArgs {
    orgId: string;
}

/**
 * A collection of values returned by getOrgLdap.
 */
export interface GetOrgLdapResult {
    readonly customSettings: outputs.GetOrgLdapCustomSetting[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ldapMode: string;
    readonly orgId: string;
}

export function getOrgLdapOutput(args: GetOrgLdapOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOrgLdapResult> {
    return pulumi.output(args).apply(a => getOrgLdap(a, opts))
}

/**
 * A collection of arguments for invoking getOrgLdap.
 */
export interface GetOrgLdapOutputArgs {
    orgId: pulumi.Input<string>;
}
