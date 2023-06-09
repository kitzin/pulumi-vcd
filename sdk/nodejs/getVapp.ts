// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getVapp(args: GetVappArgs, opts?: pulumi.InvokeOptions): Promise<GetVappResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vcd:index/getVapp:getVapp", {
        "name": args.name,
        "org": args.org,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getVapp.
 */
export interface GetVappArgs {
    name: string;
    org?: string;
    vdc?: string;
}

/**
 * A collection of values returned by getVapp.
 */
export interface GetVappResult {
    readonly description: string;
    readonly guestProperties: {[key: string]: any};
    readonly href: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly leases: outputs.GetVappLease[];
    /**
     * @deprecated Use metadata_entry instead
     */
    readonly metadata: {[key: string]: any};
    readonly metadataEntries: outputs.GetVappMetadataEntry[];
    readonly name: string;
    readonly org?: string;
    readonly status: number;
    readonly statusText: string;
    readonly vdc?: string;
}

export function getVappOutput(args: GetVappOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVappResult> {
    return pulumi.output(args).apply(a => getVapp(a, opts))
}

/**
 * A collection of arguments for invoking getVapp.
 */
export interface GetVappOutputArgs {
    name: pulumi.Input<string>;
    org?: pulumi.Input<string>;
    vdc?: pulumi.Input<string>;
}
