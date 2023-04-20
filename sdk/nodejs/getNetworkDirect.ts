// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getNetworkDirect(args?: GetNetworkDirectArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkDirectResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vcd:index/getNetworkDirect:getNetworkDirect", {
        "filter": args.filter,
        "name": args.name,
        "org": args.org,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetworkDirect.
 */
export interface GetNetworkDirectArgs {
    filter?: inputs.GetNetworkDirectFilter;
    name?: string;
    org?: string;
    vdc?: string;
}

/**
 * A collection of values returned by getNetworkDirect.
 */
export interface GetNetworkDirectResult {
    readonly description: string;
    readonly externalNetwork: string;
    readonly externalNetworkDns1: string;
    readonly externalNetworkDns2: string;
    readonly externalNetworkDnsSuffix: string;
    readonly externalNetworkGateway: string;
    readonly externalNetworkNetmask: string;
    readonly filter?: outputs.GetNetworkDirectFilter;
    readonly href: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * @deprecated Use metadata_entry instead
     */
    readonly metadata: {[key: string]: any};
    readonly metadataEntries: outputs.GetNetworkDirectMetadataEntry[];
    readonly name?: string;
    readonly org?: string;
    readonly shared: boolean;
    readonly vdc?: string;
}

export function getNetworkDirectOutput(args?: GetNetworkDirectOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetworkDirectResult> {
    return pulumi.output(args).apply(a => getNetworkDirect(a, opts))
}

/**
 * A collection of arguments for invoking getNetworkDirect.
 */
export interface GetNetworkDirectOutputArgs {
    filter?: pulumi.Input<inputs.GetNetworkDirectFilterArgs>;
    name?: pulumi.Input<string>;
    org?: pulumi.Input<string>;
    vdc?: pulumi.Input<string>;
}
