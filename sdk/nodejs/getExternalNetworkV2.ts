// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getExternalNetworkV2(args: GetExternalNetworkV2Args, opts?: pulumi.InvokeOptions): Promise<GetExternalNetworkV2Result> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vcd:index/getExternalNetworkV2:getExternalNetworkV2", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getExternalNetworkV2.
 */
export interface GetExternalNetworkV2Args {
    name: string;
}

/**
 * A collection of values returned by getExternalNetworkV2.
 */
export interface GetExternalNetworkV2Result {
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipScopes: outputs.GetExternalNetworkV2IpScope[];
    readonly name: string;
    readonly nsxtNetworks: outputs.GetExternalNetworkV2NsxtNetwork[];
    readonly vsphereNetworks: outputs.GetExternalNetworkV2VsphereNetwork[];
}

export function getExternalNetworkV2Output(args: GetExternalNetworkV2OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetExternalNetworkV2Result> {
    return pulumi.output(args).apply(a => getExternalNetworkV2(a, opts))
}

/**
 * A collection of arguments for invoking getExternalNetworkV2.
 */
export interface GetExternalNetworkV2OutputArgs {
    name: pulumi.Input<string>;
}
