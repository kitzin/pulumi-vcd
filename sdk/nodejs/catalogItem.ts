// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class CatalogItem extends pulumi.CustomResource {
    /**
     * Get an existing CatalogItem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CatalogItemState, opts?: pulumi.CustomResourceOptions): CatalogItem {
        return new CatalogItem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/catalogItem:CatalogItem';

    /**
     * Returns true if the given object is an instance of CatalogItem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CatalogItem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CatalogItem.__pulumiType;
    }

    /**
     * catalog name where upload the OVA file
     */
    public readonly catalog!: pulumi.Output<string>;
    /**
     * Key and value pairs for catalog item metadata
     *
     * @deprecated Use metadata_entry instead
     */
    public readonly catalogItemMetadata!: pulumi.Output<{[key: string]: any}>;
    /**
     * Time stamp of when the item was created
     */
    public /*out*/ readonly created!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Key and value pairs for the metadata of the vApp template associated to this catalog item
     */
    public readonly metadata!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Metadata entries for the given Catalog Item
     */
    public readonly metadataEntries!: pulumi.Output<outputs.CatalogItemMetadataEntry[]>;
    /**
     * catalog item name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * Absolute or relative path to OVA
     */
    public readonly ovaPath!: pulumi.Output<string | undefined>;
    /**
     * URL of OVF file
     */
    public readonly ovfUrl!: pulumi.Output<string | undefined>;
    /**
     * shows upload progress in stdout
     */
    public readonly showUploadProgress!: pulumi.Output<boolean | undefined>;
    /**
     * size of upload file piece size in mega bytes
     */
    public readonly uploadPieceSize!: pulumi.Output<number | undefined>;

    /**
     * Create a CatalogItem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CatalogItemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CatalogItemArgs | CatalogItemState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CatalogItemState | undefined;
            resourceInputs["catalog"] = state ? state.catalog : undefined;
            resourceInputs["catalogItemMetadata"] = state ? state.catalogItemMetadata : undefined;
            resourceInputs["created"] = state ? state.created : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["metadataEntries"] = state ? state.metadataEntries : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["ovaPath"] = state ? state.ovaPath : undefined;
            resourceInputs["ovfUrl"] = state ? state.ovfUrl : undefined;
            resourceInputs["showUploadProgress"] = state ? state.showUploadProgress : undefined;
            resourceInputs["uploadPieceSize"] = state ? state.uploadPieceSize : undefined;
        } else {
            const args = argsOrState as CatalogItemArgs | undefined;
            if ((!args || args.catalog === undefined) && !opts.urn) {
                throw new Error("Missing required property 'catalog'");
            }
            resourceInputs["catalog"] = args ? args.catalog : undefined;
            resourceInputs["catalogItemMetadata"] = args ? args.catalogItemMetadata : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["metadataEntries"] = args ? args.metadataEntries : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["ovaPath"] = args ? args.ovaPath : undefined;
            resourceInputs["ovfUrl"] = args ? args.ovfUrl : undefined;
            resourceInputs["showUploadProgress"] = args ? args.showUploadProgress : undefined;
            resourceInputs["uploadPieceSize"] = args ? args.uploadPieceSize : undefined;
            resourceInputs["created"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CatalogItem.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CatalogItem resources.
 */
export interface CatalogItemState {
    /**
     * catalog name where upload the OVA file
     */
    catalog?: pulumi.Input<string>;
    /**
     * Key and value pairs for catalog item metadata
     *
     * @deprecated Use metadata_entry instead
     */
    catalogItemMetadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * Time stamp of when the item was created
     */
    created?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    /**
     * Key and value pairs for the metadata of the vApp template associated to this catalog item
     */
    metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * Metadata entries for the given Catalog Item
     */
    metadataEntries?: pulumi.Input<pulumi.Input<inputs.CatalogItemMetadataEntry>[]>;
    /**
     * catalog item name
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * Absolute or relative path to OVA
     */
    ovaPath?: pulumi.Input<string>;
    /**
     * URL of OVF file
     */
    ovfUrl?: pulumi.Input<string>;
    /**
     * shows upload progress in stdout
     */
    showUploadProgress?: pulumi.Input<boolean>;
    /**
     * size of upload file piece size in mega bytes
     */
    uploadPieceSize?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a CatalogItem resource.
 */
export interface CatalogItemArgs {
    /**
     * catalog name where upload the OVA file
     */
    catalog: pulumi.Input<string>;
    /**
     * Key and value pairs for catalog item metadata
     *
     * @deprecated Use metadata_entry instead
     */
    catalogItemMetadata?: pulumi.Input<{[key: string]: any}>;
    description?: pulumi.Input<string>;
    /**
     * Key and value pairs for the metadata of the vApp template associated to this catalog item
     */
    metadata?: pulumi.Input<{[key: string]: any}>;
    /**
     * Metadata entries for the given Catalog Item
     */
    metadataEntries?: pulumi.Input<pulumi.Input<inputs.CatalogItemMetadataEntry>[]>;
    /**
     * catalog item name
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * Absolute or relative path to OVA
     */
    ovaPath?: pulumi.Input<string>;
    /**
     * URL of OVF file
     */
    ovfUrl?: pulumi.Input<string>;
    /**
     * shows upload progress in stdout
     */
    showUploadProgress?: pulumi.Input<boolean>;
    /**
     * size of upload file piece size in mega bytes
     */
    uploadPieceSize?: pulumi.Input<number>;
}