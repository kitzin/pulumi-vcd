// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getNetworkIsolatedV2(args?: GetNetworkIsolatedV2Args, opts?: pulumi.InvokeOptions): Promise<GetNetworkIsolatedV2Result> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vcd:index/getNetworkIsolatedV2:getNetworkIsolatedV2", {
        "filter": args.filter,
        "name": args.name,
        "org": args.org,
        "ownerId": args.ownerId,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetworkIsolatedV2.
 */
export interface GetNetworkIsolatedV2Args {
    filter?: inputs.GetNetworkIsolatedV2Filter;
    name?: string;
    org?: string;
    ownerId?: string;
    /**
     * @deprecated This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
     */
    vdc?: string;
}

/**
 * A collection of values returned by getNetworkIsolatedV2.
 */
export interface GetNetworkIsolatedV2Result {
    readonly description: string;
    readonly dns1: string;
    readonly dns2: string;
    readonly dnsSuffix: string;
    readonly filter?: outputs.GetNetworkIsolatedV2Filter;
    readonly gateway: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly isShared: boolean;
    /**
     * @deprecated Use metadata_entry instead
     */
    readonly metadata: {[key: string]: any};
    readonly metadataEntries: outputs.GetNetworkIsolatedV2MetadataEntry[];
    readonly name?: string;
    readonly org?: string;
    readonly ownerId: string;
    readonly prefixLength: number;
    readonly staticIpPools: outputs.GetNetworkIsolatedV2StaticIpPool[];
    /**
     * @deprecated This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
     */
    readonly vdc: string;
}

export function getNetworkIsolatedV2Output(args?: GetNetworkIsolatedV2OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNetworkIsolatedV2Result> {
    return pulumi.output(args).apply(a => getNetworkIsolatedV2(a, opts))
}

/**
 * A collection of arguments for invoking getNetworkIsolatedV2.
 */
export interface GetNetworkIsolatedV2OutputArgs {
    filter?: pulumi.Input<inputs.GetNetworkIsolatedV2FilterArgs>;
    name?: pulumi.Input<string>;
    org?: pulumi.Input<string>;
    ownerId?: pulumi.Input<string>;
    /**
     * @deprecated This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
     */
    vdc?: pulumi.Input<string>;
}
