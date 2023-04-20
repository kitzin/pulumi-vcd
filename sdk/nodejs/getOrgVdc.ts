// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getOrgVdc(args: GetOrgVdcArgs, opts?: pulumi.InvokeOptions): Promise<GetOrgVdcResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vcd:index/getOrgVdc:getOrgVdc", {
        "metadataEntries": args.metadataEntries,
        "name": args.name,
        "org": args.org,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrgVdc.
 */
export interface GetOrgVdcArgs {
    metadataEntries?: inputs.GetOrgVdcMetadataEntry[];
    name: string;
    org?: string;
}

/**
 * A collection of values returned by getOrgVdc.
 */
export interface GetOrgVdcResult {
    readonly allocationModel: string;
    readonly allowOverCommit: boolean;
    readonly computeCapacities: outputs.GetOrgVdcComputeCapacity[];
    readonly cpuGuaranteed: number;
    readonly cpuSpeed: number;
    readonly defaultComputePolicyId: string;
    /**
     * @deprecated Use `default_compute_policy_id` attribute instead, which can support VM Sizing Policies, VM Placement Policies and vGPU Policies
     */
    readonly defaultVmSizingPolicyId: string;
    readonly description: string;
    readonly edgeClusterId: string;
    readonly elasticity: boolean;
    readonly enableFastProvisioning: boolean;
    readonly enableThinProvisioning: boolean;
    readonly enableVmDiscovery: boolean;
    readonly enabled: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly includeVmMemoryOverhead: boolean;
    readonly memoryGuaranteed: number;
    /**
     * @deprecated Use metadata_entry instead
     */
    readonly metadata: {[key: string]: any};
    readonly metadataEntries: outputs.GetOrgVdcMetadataEntry[];
    readonly name: string;
    readonly networkPoolName: string;
    readonly networkQuota: number;
    readonly nicQuota: number;
    readonly org?: string;
    readonly providerVdcName: string;
    readonly storageProfiles: outputs.GetOrgVdcStorageProfile[];
    readonly vmPlacementPolicyIds: string[];
    readonly vmQuota: number;
    readonly vmSizingPolicyIds: string[];
}

export function getOrgVdcOutput(args: GetOrgVdcOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOrgVdcResult> {
    return pulumi.output(args).apply(a => getOrgVdc(a, opts))
}

/**
 * A collection of arguments for invoking getOrgVdc.
 */
export interface GetOrgVdcOutputArgs {
    metadataEntries?: pulumi.Input<pulumi.Input<inputs.GetOrgVdcMetadataEntryArgs>[]>;
    name: pulumi.Input<string>;
    org?: pulumi.Input<string>;
}