// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export function getNsxtIpsecVpnTunnel(args: GetNsxtIpsecVpnTunnelArgs, opts?: pulumi.InvokeOptions): Promise<GetNsxtIpsecVpnTunnelResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vcd:index/getNsxtIpsecVpnTunnel:getNsxtIpsecVpnTunnel", {
        "edgeGatewayId": args.edgeGatewayId,
        "name": args.name,
        "org": args.org,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getNsxtIpsecVpnTunnel.
 */
export interface GetNsxtIpsecVpnTunnelArgs {
    edgeGatewayId: string;
    name: string;
    org?: string;
    /**
     * @deprecated Edge Gateway will be looked up based on 'edge_gateway_id' field
     */
    vdc?: string;
}

/**
 * A collection of values returned by getNsxtIpsecVpnTunnel.
 */
export interface GetNsxtIpsecVpnTunnelResult {
    readonly description: string;
    readonly edgeGatewayId: string;
    readonly enabled: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ikeFailReason: string;
    readonly ikeServiceStatus: string;
    readonly localIpAddress: string;
    readonly localNetworks: string[];
    readonly logging: boolean;
    readonly name: string;
    readonly org?: string;
    readonly preSharedKey: string;
    readonly remoteIpAddress: string;
    readonly remoteNetworks: string[];
    readonly securityProfile: string;
    readonly securityProfileCustomizations: outputs.GetNsxtIpsecVpnTunnelSecurityProfileCustomization[];
    readonly status: string;
    /**
     * @deprecated Edge Gateway will be looked up based on 'edge_gateway_id' field
     */
    readonly vdc?: string;
}

export function getNsxtIpsecVpnTunnelOutput(args: GetNsxtIpsecVpnTunnelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetNsxtIpsecVpnTunnelResult> {
    return pulumi.output(args).apply(a => getNsxtIpsecVpnTunnel(a, opts))
}

/**
 * A collection of arguments for invoking getNsxtIpsecVpnTunnel.
 */
export interface GetNsxtIpsecVpnTunnelOutputArgs {
    edgeGatewayId: pulumi.Input<string>;
    name: pulumi.Input<string>;
    org?: pulumi.Input<string>;
    /**
     * @deprecated Edge Gateway will be looked up based on 'edge_gateway_id' field
     */
    vdc?: pulumi.Input<string>;
}
