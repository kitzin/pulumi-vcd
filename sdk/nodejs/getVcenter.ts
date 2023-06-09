// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export function getVcenter(args: GetVcenterArgs, opts?: pulumi.InvokeOptions): Promise<GetVcenterResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vcd:index/getVcenter:getVcenter", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getVcenter.
 */
export interface GetVcenterArgs {
    name: string;
}

/**
 * A collection of values returned by getVcenter.
 */
export interface GetVcenterResult {
    readonly connectionStatus: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly isEnabled: boolean;
    readonly name: string;
    readonly status: string;
    readonly vcenterHost: string;
    readonly vcenterVersion: string;
}

export function getVcenterOutput(args: GetVcenterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVcenterResult> {
    return pulumi.output(args).apply(a => getVcenter(a, opts))
}

/**
 * A collection of arguments for invoking getVcenter.
 */
export interface GetVcenterOutputArgs {
    name: pulumi.Input<string>;
}
