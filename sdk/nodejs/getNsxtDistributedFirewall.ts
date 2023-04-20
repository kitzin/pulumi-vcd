// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getNsxtDistributedFirewall(args: GetNsxtDistributedFirewallArgs, opts?: pulumi.InvokeOptions): Promise<GetNsxtDistributedFirewallResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vcd:index/getNsxtDistributedFirewall:getNsxtDistributedFirewall", {
        "org": args.org,
        "vdcGroupId": args.vdcGroupId,
    }, opts);
}

/**
 * A collection of arguments for invoking getNsxtDistributedFirewall.
 */
export interface GetNsxtDistributedFirewallArgs {
    org?: string;
    vdcGroupId: string;
}

/**
 * A collection of values returned by getNsxtDistributedFirewall.
 */
export interface GetNsxtDistributedFirewallResult {
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly org?: string;
    readonly rules: outputs.GetNsxtDistributedFirewallRule[];
    readonly vdcGroupId: string;
}

export function getNsxtDistributedFirewallOutput(args: GetNsxtDistributedFirewallOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNsxtDistributedFirewallResult> {
    return pulumi.output(args).apply(a => getNsxtDistributedFirewall(a, opts))
}

/**
 * A collection of arguments for invoking getNsxtDistributedFirewall.
 */
export interface GetNsxtDistributedFirewallOutputArgs {
    org?: pulumi.Input<string>;
    vdcGroupId: pulumi.Input<string>;
}
