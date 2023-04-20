// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getProviderVdc(args: GetProviderVdcArgs, opts?: pulumi.InvokeOptions): Promise<GetProviderVdcResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vcd:index/getProviderVdc:getProviderVdc", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getProviderVdc.
 */
export interface GetProviderVdcArgs {
    name: string;
}

/**
 * A collection of values returned by getProviderVdc.
 */
export interface GetProviderVdcResult {
    readonly capabilities: string[];
    readonly computeCapacities: outputs.GetProviderVdcComputeCapacity[];
    readonly computeProviderScope: string;
    readonly description: string;
    readonly externalNetworkIds: string[];
    readonly highestSupportedHardwareVersion: string;
    readonly hostIds: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly isEnabled: boolean;
    /**
     * @deprecated Use metadata_entry instead
     */
    readonly metadata: {[key: string]: any};
    readonly metadataEntries: outputs.GetProviderVdcMetadataEntry[];
    readonly name: string;
    readonly networkPoolIds: string[];
    readonly nsxtManagerId: string;
    readonly resourcePoolIds: string[];
    readonly status: number;
    readonly storageContainerIds: string[];
    readonly storageProfileIds: string[];
    readonly universalNetworkPoolId: string;
    readonly vcenterId: string;
}

export function getProviderVdcOutput(args: GetProviderVdcOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetProviderVdcResult> {
    return pulumi.output(args).apply(a => getProviderVdc(a, opts))
}

/**
 * A collection of arguments for invoking getProviderVdc.
 */
export interface GetProviderVdcOutputArgs {
    name: pulumi.Input<string>;
}
