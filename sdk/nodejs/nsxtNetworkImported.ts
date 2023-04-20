// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class NsxtNetworkImported extends pulumi.CustomResource {
    /**
     * Get an existing NsxtNetworkImported resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NsxtNetworkImportedState, opts?: pulumi.CustomResourceOptions): NsxtNetworkImported {
        return new NsxtNetworkImported(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/nsxtNetworkImported:NsxtNetworkImported';

    /**
     * Returns true if the given object is an instance of NsxtNetworkImported.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NsxtNetworkImported {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NsxtNetworkImported.__pulumiType;
    }

    /**
     * Network description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * DNS server 1
     */
    public readonly dns1!: pulumi.Output<string | undefined>;
    /**
     * DNS server 1
     */
    public readonly dns2!: pulumi.Output<string | undefined>;
    /**
     * DNS suffix
     */
    public readonly dnsSuffix!: pulumi.Output<string | undefined>;
    /**
     * Gateway IP address
     */
    public readonly gateway!: pulumi.Output<string>;
    /**
     * Network name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * ID of existing NSX-T Logical Switch
     */
    public /*out*/ readonly nsxtLogicalSwitchId!: pulumi.Output<string>;
    /**
     * Name of existing NSX-T Logical Switch
     */
    public readonly nsxtLogicalSwitchName!: pulumi.Output<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * ID of VDC or VDC Group
     */
    public readonly ownerId!: pulumi.Output<string>;
    /**
     * Network prefix
     */
    public readonly prefixLength!: pulumi.Output<number>;
    /**
     * IP ranges used for static pool allocation in the network
     */
    public readonly staticIpPools!: pulumi.Output<outputs.NsxtNetworkImportedStaticIpPool[] | undefined>;
    /**
     * The name of VDC to use, optional if defined at provider level
     *
     * @deprecated This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
     */
    public readonly vdc!: pulumi.Output<string>;

    /**
     * Create a NsxtNetworkImported resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NsxtNetworkImportedArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NsxtNetworkImportedArgs | NsxtNetworkImportedState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NsxtNetworkImportedState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dns1"] = state ? state.dns1 : undefined;
            resourceInputs["dns2"] = state ? state.dns2 : undefined;
            resourceInputs["dnsSuffix"] = state ? state.dnsSuffix : undefined;
            resourceInputs["gateway"] = state ? state.gateway : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nsxtLogicalSwitchId"] = state ? state.nsxtLogicalSwitchId : undefined;
            resourceInputs["nsxtLogicalSwitchName"] = state ? state.nsxtLogicalSwitchName : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["ownerId"] = state ? state.ownerId : undefined;
            resourceInputs["prefixLength"] = state ? state.prefixLength : undefined;
            resourceInputs["staticIpPools"] = state ? state.staticIpPools : undefined;
            resourceInputs["vdc"] = state ? state.vdc : undefined;
        } else {
            const args = argsOrState as NsxtNetworkImportedArgs | undefined;
            if ((!args || args.gateway === undefined) && !opts.urn) {
                throw new Error("Missing required property 'gateway'");
            }
            if ((!args || args.nsxtLogicalSwitchName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nsxtLogicalSwitchName'");
            }
            if ((!args || args.prefixLength === undefined) && !opts.urn) {
                throw new Error("Missing required property 'prefixLength'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dns1"] = args ? args.dns1 : undefined;
            resourceInputs["dns2"] = args ? args.dns2 : undefined;
            resourceInputs["dnsSuffix"] = args ? args.dnsSuffix : undefined;
            resourceInputs["gateway"] = args ? args.gateway : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nsxtLogicalSwitchName"] = args ? args.nsxtLogicalSwitchName : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["ownerId"] = args ? args.ownerId : undefined;
            resourceInputs["prefixLength"] = args ? args.prefixLength : undefined;
            resourceInputs["staticIpPools"] = args ? args.staticIpPools : undefined;
            resourceInputs["vdc"] = args ? args.vdc : undefined;
            resourceInputs["nsxtLogicalSwitchId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NsxtNetworkImported.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NsxtNetworkImported resources.
 */
export interface NsxtNetworkImportedState {
    /**
     * Network description
     */
    description?: pulumi.Input<string>;
    /**
     * DNS server 1
     */
    dns1?: pulumi.Input<string>;
    /**
     * DNS server 1
     */
    dns2?: pulumi.Input<string>;
    /**
     * DNS suffix
     */
    dnsSuffix?: pulumi.Input<string>;
    /**
     * Gateway IP address
     */
    gateway?: pulumi.Input<string>;
    /**
     * Network name
     */
    name?: pulumi.Input<string>;
    /**
     * ID of existing NSX-T Logical Switch
     */
    nsxtLogicalSwitchId?: pulumi.Input<string>;
    /**
     * Name of existing NSX-T Logical Switch
     */
    nsxtLogicalSwitchName?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * ID of VDC or VDC Group
     */
    ownerId?: pulumi.Input<string>;
    /**
     * Network prefix
     */
    prefixLength?: pulumi.Input<number>;
    /**
     * IP ranges used for static pool allocation in the network
     */
    staticIpPools?: pulumi.Input<pulumi.Input<inputs.NsxtNetworkImportedStaticIpPool>[]>;
    /**
     * The name of VDC to use, optional if defined at provider level
     *
     * @deprecated This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
     */
    vdc?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NsxtNetworkImported resource.
 */
export interface NsxtNetworkImportedArgs {
    /**
     * Network description
     */
    description?: pulumi.Input<string>;
    /**
     * DNS server 1
     */
    dns1?: pulumi.Input<string>;
    /**
     * DNS server 1
     */
    dns2?: pulumi.Input<string>;
    /**
     * DNS suffix
     */
    dnsSuffix?: pulumi.Input<string>;
    /**
     * Gateway IP address
     */
    gateway: pulumi.Input<string>;
    /**
     * Network name
     */
    name?: pulumi.Input<string>;
    /**
     * Name of existing NSX-T Logical Switch
     */
    nsxtLogicalSwitchName: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * ID of VDC or VDC Group
     */
    ownerId?: pulumi.Input<string>;
    /**
     * Network prefix
     */
    prefixLength: pulumi.Input<number>;
    /**
     * IP ranges used for static pool allocation in the network
     */
    staticIpPools?: pulumi.Input<pulumi.Input<inputs.NsxtNetworkImportedStaticIpPool>[]>;
    /**
     * The name of VDC to use, optional if defined at provider level
     *
     * @deprecated This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
     */
    vdc?: pulumi.Input<string>;
}
