// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getOrgGroup(args: GetOrgGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetOrgGroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vcd:index/getOrgGroup:getOrgGroup", {
        "name": args.name,
        "org": args.org,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrgGroup.
 */
export interface GetOrgGroupArgs {
    name: string;
    org?: string;
}

/**
 * A collection of values returned by getOrgGroup.
 */
export interface GetOrgGroupResult {
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly org?: string;
    readonly providerType: string;
    readonly role: string;
    readonly userNames: string[];
}

export function getOrgGroupOutput(args: GetOrgGroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOrgGroupResult> {
    return pulumi.output(args).apply(a => getOrgGroup(a, opts))
}

/**
 * A collection of arguments for invoking getOrgGroup.
 */
export interface GetOrgGroupOutputArgs {
    name: pulumi.Input<string>;
    org?: pulumi.Input<string>;
}
