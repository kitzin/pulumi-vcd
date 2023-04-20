// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getPortgroup(args: GetPortgroupArgs, opts?: pulumi.InvokeOptions): Promise<GetPortgroupResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vcd:index/getPortgroup:getPortgroup", {
        "name": args.name,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getPortgroup.
 */
export interface GetPortgroupArgs {
    name: string;
    type: string;
}

/**
 * A collection of values returned by getPortgroup.
 */
export interface GetPortgroupResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly type: string;
}

export function getPortgroupOutput(args: GetPortgroupOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPortgroupResult> {
    return pulumi.output(args).apply(a => getPortgroup(a, opts))
}

/**
 * A collection of arguments for invoking getPortgroup.
 */
export interface GetPortgroupOutputArgs {
    name: pulumi.Input<string>;
    type: pulumi.Input<string>;
}