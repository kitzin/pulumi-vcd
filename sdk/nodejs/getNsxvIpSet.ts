// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getNsxvIpSet(args: GetNsxvIpSetArgs, opts?: pulumi.InvokeOptions): Promise<GetNsxvIpSetResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vcd:index/getNsxvIpSet:getNsxvIpSet", {
        "name": args.name,
        "org": args.org,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getNsxvIpSet.
 */
export interface GetNsxvIpSetArgs {
    name: string;
    org?: string;
    vdc?: string;
}

/**
 * A collection of values returned by getNsxvIpSet.
 */
export interface GetNsxvIpSetResult {
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipAddresses: string[];
    readonly isInheritanceAllowed: boolean;
    readonly name: string;
    readonly org?: string;
    readonly vdc?: string;
}

export function getNsxvIpSetOutput(args: GetNsxvIpSetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNsxvIpSetResult> {
    return pulumi.output(args).apply(a => getNsxvIpSet(a, opts))
}

/**
 * A collection of arguments for invoking getNsxvIpSet.
 */
export interface GetNsxvIpSetOutputArgs {
    name: pulumi.Input<string>;
    org?: pulumi.Input<string>;
    vdc?: pulumi.Input<string>;
}
