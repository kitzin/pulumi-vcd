// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class NetworkRouted extends pulumi.CustomResource {
    /**
     * Get an existing NetworkRouted resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkRoutedState, opts?: pulumi.CustomResourceOptions): NetworkRouted {
        return new NetworkRouted(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/networkRouted:NetworkRouted';

    /**
     * Returns true if the given object is an instance of NetworkRouted.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkRouted {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkRouted.__pulumiType;
    }

    /**
     * Optional description for the network
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A range of IPs to issue to virtual machines that don't have a static IP
     */
    public readonly dhcpPools!: pulumi.Output<outputs.NetworkRoutedDhcpPool[] | undefined>;
    /**
     * First DNS server to use
     */
    public readonly dns1!: pulumi.Output<string | undefined>;
    /**
     * Second DNS server to use
     */
    public readonly dns2!: pulumi.Output<string | undefined>;
    /**
     * A FQDN for the virtual machines on this network
     */
    public readonly dnsSuffix!: pulumi.Output<string | undefined>;
    /**
     * The name of the edge gateway
     */
    public readonly edgeGateway!: pulumi.Output<string>;
    /**
     * The gateway of this network
     */
    public readonly gateway!: pulumi.Output<string | undefined>;
    /**
     * Network Hypertext Reference
     */
    public /*out*/ readonly href!: pulumi.Output<string>;
    /**
     * Which interface to use (one of `internal`, `subinterface`, `distributed`)
     */
    public readonly interfaceType!: pulumi.Output<string | undefined>;
    /**
     * Key value map of metadata to assign to this network. Key and value can be any string
     *
     * @deprecated Use metadata_entry instead
     */
    public readonly metadata!: pulumi.Output<{[key: string]: any}>;
    /**
     * Metadata entries for the given Network
     */
    public readonly metadataEntries!: pulumi.Output<outputs.NetworkRoutedMetadataEntry[]>;
    /**
     * A unique name for the network
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The netmask for the new network
     */
    public readonly netmask!: pulumi.Output<string | undefined>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * Defines if this network is shared between multiple VDCs in the Org
     */
    public readonly shared!: pulumi.Output<boolean | undefined>;
    /**
     * A range of IPs permitted to be used as static IPs for virtual machines
     */
    public readonly staticIpPools!: pulumi.Output<outputs.NetworkRoutedStaticIpPool[] | undefined>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    public readonly vdc!: pulumi.Output<string | undefined>;

    /**
     * Create a NetworkRouted resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkRoutedArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkRoutedArgs | NetworkRoutedState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkRoutedState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dhcpPools"] = state ? state.dhcpPools : undefined;
            resourceInputs["dns1"] = state ? state.dns1 : undefined;
            resourceInputs["dns2"] = state ? state.dns2 : undefined;
            resourceInputs["dnsSuffix"] = state ? state.dnsSuffix : undefined;
            resourceInputs["edgeGateway"] = state ? state.edgeGateway : undefined;
            resourceInputs["gateway"] = state ? state.gateway : undefined;
            resourceInputs["href"] = state ? state.href : undefined;
            resourceInputs["interfaceType"] = state ? state.interfaceType : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["metadataEntries"] = state ? state.metadataEntries : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["netmask"] = state ? state.netmask : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["shared"] = state ? state.shared : undefined;
            resourceInputs["staticIpPools"] = state ? state.staticIpPools : undefined;
            resourceInputs["vdc"] = state ? state.vdc : undefined;
        } else {
            const args = argsOrState as NetworkRoutedArgs | undefined;
            if ((!args || args.edgeGateway === undefined) && !opts.urn) {
                throw new Error("Missing required property 'edgeGateway'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dhcpPools"] = args ? args.dhcpPools : undefined;
            resourceInputs["dns1"] = args ? args.dns1 : undefined;
            resourceInputs["dns2"] = args ? args.dns2 : undefined;
            resourceInputs["dnsSuffix"] = args ? args.dnsSuffix : undefined;
            resourceInputs["edgeGateway"] = args ? args.edgeGateway : undefined;
            resourceInputs["gateway"] = args ? args.gateway : undefined;
            resourceInputs["interfaceType"] = args ? args.interfaceType : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["metadataEntries"] = args ? args.metadataEntries : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["netmask"] = args ? args.netmask : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["shared"] = args ? args.shared : undefined;
            resourceInputs["staticIpPools"] = args ? args.staticIpPools : undefined;
            resourceInputs["vdc"] = args ? args.vdc : undefined;
            resourceInputs["href"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetworkRouted.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkRouted resources.
 */
export interface NetworkRoutedState {
    /**
     * Optional description for the network
     */
    description?: pulumi.Input<string>;
    /**
     * A range of IPs to issue to virtual machines that don't have a static IP
     */
    dhcpPools?: pulumi.Input<pulumi.Input<inputs.NetworkRoutedDhcpPool>[]>;
    /**
     * First DNS server to use
     */
    dns1?: pulumi.Input<string>;
    /**
     * Second DNS server to use
     */
    dns2?: pulumi.Input<string>;
    /**
     * A FQDN for the virtual machines on this network
     */
    dnsSuffix?: pulumi.Input<string>;
    /**
     * The name of the edge gateway
     */
    edgeGateway?: pulumi.Input<string>;
    /**
     * The gateway of this network
     */
    gateway?: pulumi.Input<string>;
    /**
     * Network Hypertext Reference
     */
    href?: pulumi.Input<string>;
    /**
     * Which interface to use (one of `internal`, `subinterface`, `distributed`)
     */
    interfaceType?: pulumi.Input<string>;
    /**
     * Key value map of metadata to assign to this network. Key and value can be any string
     *
     * @deprecated Use metadata_entry instead
     */
    metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * Metadata entries for the given Network
     */
    metadataEntries?: pulumi.Input<pulumi.Input<inputs.NetworkRoutedMetadataEntry>[]>;
    /**
     * A unique name for the network
     */
    name?: pulumi.Input<string>;
    /**
     * The netmask for the new network
     */
    netmask?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * Defines if this network is shared between multiple VDCs in the Org
     */
    shared?: pulumi.Input<boolean>;
    /**
     * A range of IPs permitted to be used as static IPs for virtual machines
     */
    staticIpPools?: pulumi.Input<pulumi.Input<inputs.NetworkRoutedStaticIpPool>[]>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkRouted resource.
 */
export interface NetworkRoutedArgs {
    /**
     * Optional description for the network
     */
    description?: pulumi.Input<string>;
    /**
     * A range of IPs to issue to virtual machines that don't have a static IP
     */
    dhcpPools?: pulumi.Input<pulumi.Input<inputs.NetworkRoutedDhcpPool>[]>;
    /**
     * First DNS server to use
     */
    dns1?: pulumi.Input<string>;
    /**
     * Second DNS server to use
     */
    dns2?: pulumi.Input<string>;
    /**
     * A FQDN for the virtual machines on this network
     */
    dnsSuffix?: pulumi.Input<string>;
    /**
     * The name of the edge gateway
     */
    edgeGateway: pulumi.Input<string>;
    /**
     * The gateway of this network
     */
    gateway?: pulumi.Input<string>;
    /**
     * Which interface to use (one of `internal`, `subinterface`, `distributed`)
     */
    interfaceType?: pulumi.Input<string>;
    /**
     * Key value map of metadata to assign to this network. Key and value can be any string
     *
     * @deprecated Use metadata_entry instead
     */
    metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * Metadata entries for the given Network
     */
    metadataEntries?: pulumi.Input<pulumi.Input<inputs.NetworkRoutedMetadataEntry>[]>;
    /**
     * A unique name for the network
     */
    name?: pulumi.Input<string>;
    /**
     * The netmask for the new network
     */
    netmask?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * Defines if this network is shared between multiple VDCs in the Org
     */
    shared?: pulumi.Input<boolean>;
    /**
     * A range of IPs permitted to be used as static IPs for virtual machines
     */
    staticIpPools?: pulumi.Input<pulumi.Input<inputs.NetworkRoutedStaticIpPool>[]>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: pulumi.Input<string>;
}