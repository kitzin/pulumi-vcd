// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getVappNetwork(args: GetVappNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetVappNetworkResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vcd:index/getVappNetwork:getVappNetwork", {
        "name": args.name,
        "org": args.org,
        "vappName": args.vappName,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getVappNetwork.
 */
export interface GetVappNetworkArgs {
    name: string;
    org?: string;
    vappName: string;
    vdc?: string;
}

/**
 * A collection of values returned by getVappNetwork.
 */
export interface GetVappNetworkResult {
    readonly description: string;
    readonly dhcpPools: outputs.GetVappNetworkDhcpPool[];
    readonly dns1: string;
    readonly dns2: string;
    readonly dnsSuffix: string;
    readonly gateway: string;
    readonly guestVlanAllowed: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly netmask: string;
    readonly org?: string;
    readonly orgNetworkName: string;
    readonly retainIpMacEnabled: boolean;
    readonly staticIpPools: outputs.GetVappNetworkStaticIpPool[];
    readonly vappName: string;
    readonly vdc?: string;
}

export function getVappNetworkOutput(args: GetVappNetworkOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVappNetworkResult> {
    return pulumi.output(args).apply(a => getVappNetwork(a, opts))
}

/**
 * A collection of arguments for invoking getVappNetwork.
 */
export interface GetVappNetworkOutputArgs {
    name: pulumi.Input<string>;
    org?: pulumi.Input<string>;
    vappName: pulumi.Input<string>;
    vdc?: pulumi.Input<string>;
}