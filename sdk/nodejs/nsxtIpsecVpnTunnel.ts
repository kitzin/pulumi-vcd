// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class NsxtIpsecVpnTunnel extends pulumi.CustomResource {
    /**
     * Get an existing NsxtIpsecVpnTunnel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NsxtIpsecVpnTunnelState, opts?: pulumi.CustomResourceOptions): NsxtIpsecVpnTunnel {
        return new NsxtIpsecVpnTunnel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/nsxtIpsecVpnTunnel:NsxtIpsecVpnTunnel';

    /**
     * Returns true if the given object is an instance of NsxtIpsecVpnTunnel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NsxtIpsecVpnTunnel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NsxtIpsecVpnTunnel.__pulumiType;
    }

    /**
     * Description IP Sec VPN Tunnel
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Edge gateway name in which IP Sec VPN configuration is located
     */
    public readonly edgeGatewayId!: pulumi.Output<string>;
    /**
     * Enables or disables this configuration (default true)
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Provides more details of failure if the IKE service is not UP
     */
    public /*out*/ readonly ikeFailReason!: pulumi.Output<string>;
    /**
     * Status for the actual IKE Session for the given tunnel
     */
    public /*out*/ readonly ikeServiceStatus!: pulumi.Output<string>;
    /**
     * IPv4 Address for the endpoint. This has to be a sub-allocated IP on the Edge Gateway.
     */
    public readonly localIpAddress!: pulumi.Output<string>;
    /**
     * Set of local networks in CIDR format. At least one value is required
     */
    public readonly localNetworks!: pulumi.Output<string[]>;
    /**
     * Sets whether logging for the tunnel is enabled or not. (default - false)
     */
    public readonly logging!: pulumi.Output<boolean | undefined>;
    /**
     * Name of IP Sec VPN Tunnel
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * Pre-Shared Key (PSK)
     */
    public readonly preSharedKey!: pulumi.Output<string>;
    /**
     * Public IPv4 Address of the remote device terminating the VPN connection
     */
    public readonly remoteIpAddress!: pulumi.Output<string>;
    /**
     * Set of remote networks in CIDR format. Leaving it empty is interpreted as 0.0.0.0/0
     */
    public readonly remoteNetworks!: pulumi.Output<string[] | undefined>;
    /**
     * Security type which is use for IPsec VPN Tunnel. It will be 'DEFAULT' if nothing is customized and 'CUSTOM' if some
     * changes are applied
     */
    public /*out*/ readonly securityProfile!: pulumi.Output<string>;
    /**
     * Security profile customization
     */
    public readonly securityProfileCustomization!: pulumi.Output<outputs.NsxtIpsecVpnTunnelSecurityProfileCustomization | undefined>;
    /**
     * Overall IPsec VPN Tunnel Status
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The name of VDC to use, optional if defined at provider level
     *
     * @deprecated Edge Gateway will be looked up based on 'edge_gateway_id' field
     */
    public readonly vdc!: pulumi.Output<string>;

    /**
     * Create a NsxtIpsecVpnTunnel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NsxtIpsecVpnTunnelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NsxtIpsecVpnTunnelArgs | NsxtIpsecVpnTunnelState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NsxtIpsecVpnTunnelState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["edgeGatewayId"] = state ? state.edgeGatewayId : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["ikeFailReason"] = state ? state.ikeFailReason : undefined;
            resourceInputs["ikeServiceStatus"] = state ? state.ikeServiceStatus : undefined;
            resourceInputs["localIpAddress"] = state ? state.localIpAddress : undefined;
            resourceInputs["localNetworks"] = state ? state.localNetworks : undefined;
            resourceInputs["logging"] = state ? state.logging : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["preSharedKey"] = state ? state.preSharedKey : undefined;
            resourceInputs["remoteIpAddress"] = state ? state.remoteIpAddress : undefined;
            resourceInputs["remoteNetworks"] = state ? state.remoteNetworks : undefined;
            resourceInputs["securityProfile"] = state ? state.securityProfile : undefined;
            resourceInputs["securityProfileCustomization"] = state ? state.securityProfileCustomization : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdc"] = state ? state.vdc : undefined;
        } else {
            const args = argsOrState as NsxtIpsecVpnTunnelArgs | undefined;
            if ((!args || args.edgeGatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'edgeGatewayId'");
            }
            if ((!args || args.localIpAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'localIpAddress'");
            }
            if ((!args || args.localNetworks === undefined) && !opts.urn) {
                throw new Error("Missing required property 'localNetworks'");
            }
            if ((!args || args.preSharedKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'preSharedKey'");
            }
            if ((!args || args.remoteIpAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'remoteIpAddress'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["edgeGatewayId"] = args ? args.edgeGatewayId : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["localIpAddress"] = args ? args.localIpAddress : undefined;
            resourceInputs["localNetworks"] = args ? args.localNetworks : undefined;
            resourceInputs["logging"] = args ? args.logging : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["preSharedKey"] = args ? args.preSharedKey : undefined;
            resourceInputs["remoteIpAddress"] = args ? args.remoteIpAddress : undefined;
            resourceInputs["remoteNetworks"] = args ? args.remoteNetworks : undefined;
            resourceInputs["securityProfileCustomization"] = args ? args.securityProfileCustomization : undefined;
            resourceInputs["vdc"] = args ? args.vdc : undefined;
            resourceInputs["ikeFailReason"] = undefined /*out*/;
            resourceInputs["ikeServiceStatus"] = undefined /*out*/;
            resourceInputs["securityProfile"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NsxtIpsecVpnTunnel.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NsxtIpsecVpnTunnel resources.
 */
export interface NsxtIpsecVpnTunnelState {
    /**
     * Description IP Sec VPN Tunnel
     */
    description?: pulumi.Input<string>;
    /**
     * Edge gateway name in which IP Sec VPN configuration is located
     */
    edgeGatewayId?: pulumi.Input<string>;
    /**
     * Enables or disables this configuration (default true)
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Provides more details of failure if the IKE service is not UP
     */
    ikeFailReason?: pulumi.Input<string>;
    /**
     * Status for the actual IKE Session for the given tunnel
     */
    ikeServiceStatus?: pulumi.Input<string>;
    /**
     * IPv4 Address for the endpoint. This has to be a sub-allocated IP on the Edge Gateway.
     */
    localIpAddress?: pulumi.Input<string>;
    /**
     * Set of local networks in CIDR format. At least one value is required
     */
    localNetworks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Sets whether logging for the tunnel is enabled or not. (default - false)
     */
    logging?: pulumi.Input<boolean>;
    /**
     * Name of IP Sec VPN Tunnel
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * Pre-Shared Key (PSK)
     */
    preSharedKey?: pulumi.Input<string>;
    /**
     * Public IPv4 Address of the remote device terminating the VPN connection
     */
    remoteIpAddress?: pulumi.Input<string>;
    /**
     * Set of remote networks in CIDR format. Leaving it empty is interpreted as 0.0.0.0/0
     */
    remoteNetworks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Security type which is use for IPsec VPN Tunnel. It will be 'DEFAULT' if nothing is customized and 'CUSTOM' if some
     * changes are applied
     */
    securityProfile?: pulumi.Input<string>;
    /**
     * Security profile customization
     */
    securityProfileCustomization?: pulumi.Input<inputs.NsxtIpsecVpnTunnelSecurityProfileCustomization>;
    /**
     * Overall IPsec VPN Tunnel Status
     */
    status?: pulumi.Input<string>;
    /**
     * The name of VDC to use, optional if defined at provider level
     *
     * @deprecated Edge Gateway will be looked up based on 'edge_gateway_id' field
     */
    vdc?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NsxtIpsecVpnTunnel resource.
 */
export interface NsxtIpsecVpnTunnelArgs {
    /**
     * Description IP Sec VPN Tunnel
     */
    description?: pulumi.Input<string>;
    /**
     * Edge gateway name in which IP Sec VPN configuration is located
     */
    edgeGatewayId: pulumi.Input<string>;
    /**
     * Enables or disables this configuration (default true)
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * IPv4 Address for the endpoint. This has to be a sub-allocated IP on the Edge Gateway.
     */
    localIpAddress: pulumi.Input<string>;
    /**
     * Set of local networks in CIDR format. At least one value is required
     */
    localNetworks: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Sets whether logging for the tunnel is enabled or not. (default - false)
     */
    logging?: pulumi.Input<boolean>;
    /**
     * Name of IP Sec VPN Tunnel
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * Pre-Shared Key (PSK)
     */
    preSharedKey: pulumi.Input<string>;
    /**
     * Public IPv4 Address of the remote device terminating the VPN connection
     */
    remoteIpAddress: pulumi.Input<string>;
    /**
     * Set of remote networks in CIDR format. Leaving it empty is interpreted as 0.0.0.0/0
     */
    remoteNetworks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Security profile customization
     */
    securityProfileCustomization?: pulumi.Input<inputs.NsxtIpsecVpnTunnelSecurityProfileCustomization>;
    /**
     * The name of VDC to use, optional if defined at provider level
     *
     * @deprecated Edge Gateway will be looked up based on 'edge_gateway_id' field
     */
    vdc?: pulumi.Input<string>;
}