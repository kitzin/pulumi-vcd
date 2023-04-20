// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class LbAppProfile extends pulumi.CustomResource {
    /**
     * Get an existing LbAppProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LbAppProfileState, opts?: pulumi.CustomResourceOptions): LbAppProfile {
        return new LbAppProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/lbAppProfile:LbAppProfile';

    /**
     * Returns true if the given object is an instance of LbAppProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LbAppProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LbAppProfile.__pulumiType;
    }

    /**
     * The mode by which the cookie should be inserted. One of 'insert', 'prefix', or 'appsession'
     */
    public readonly cookieMode!: pulumi.Output<string | undefined>;
    /**
     * Used to uniquely identify the session the first time a client accesses the site. The load balancer refers to this cookie
     * when connecting subsequent requests in the session, so that they all go to the same virtual server. Only applies for
     * persistence_mechanism 'cookie'
     */
    public readonly cookieName!: pulumi.Output<string | undefined>;
    /**
     * Edge gateway name in which the LB Application Profile is located
     */
    public readonly edgeGateway!: pulumi.Output<string>;
    /**
     * Enable to define the certificate, CAs, or CRLs used to authenticate the load balancer from the server side
     */
    public readonly enablePoolSideSsl!: pulumi.Output<boolean | undefined>;
    /**
     * Enable SSL authentication to be passed through to the virtual server. Otherwise SSL authentication takes place at the
     * destination address.
     */
    public readonly enableSslPassthrough!: pulumi.Output<boolean | undefined>;
    /**
     * Length of time in seconds that persistence stays in effect
     */
    public readonly expiration!: pulumi.Output<number | undefined>;
    /**
     * The URL to which traffic that arrives at the destination address should be redirected. Only applies for types 'http' and
     * 'https'
     */
    public readonly httpRedirectUrl!: pulumi.Output<string | undefined>;
    /**
     * Enables 'X-Forwarded-For' header for identifying the originating IP address of a client connecting to a Web server
     * through the load balancer. Only applies for types HTTP and HTTPS
     */
    public readonly insertXForwardedHttpHeader!: pulumi.Output<boolean | undefined>;
    /**
     * Unique LB Application Profile name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * Persistence mechanism for the profile. One of 'cookie', 'ssl-sessionid', 'sourceip'
     */
    public readonly persistenceMechanism!: pulumi.Output<string | undefined>;
    /**
     * Protocol type used to send requests to the server. One of 'tcp', 'udp', 'http' org 'https'
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    public readonly vdc!: pulumi.Output<string | undefined>;

    /**
     * Create a LbAppProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LbAppProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LbAppProfileArgs | LbAppProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LbAppProfileState | undefined;
            resourceInputs["cookieMode"] = state ? state.cookieMode : undefined;
            resourceInputs["cookieName"] = state ? state.cookieName : undefined;
            resourceInputs["edgeGateway"] = state ? state.edgeGateway : undefined;
            resourceInputs["enablePoolSideSsl"] = state ? state.enablePoolSideSsl : undefined;
            resourceInputs["enableSslPassthrough"] = state ? state.enableSslPassthrough : undefined;
            resourceInputs["expiration"] = state ? state.expiration : undefined;
            resourceInputs["httpRedirectUrl"] = state ? state.httpRedirectUrl : undefined;
            resourceInputs["insertXForwardedHttpHeader"] = state ? state.insertXForwardedHttpHeader : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["persistenceMechanism"] = state ? state.persistenceMechanism : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["vdc"] = state ? state.vdc : undefined;
        } else {
            const args = argsOrState as LbAppProfileArgs | undefined;
            if ((!args || args.edgeGateway === undefined) && !opts.urn) {
                throw new Error("Missing required property 'edgeGateway'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["cookieMode"] = args ? args.cookieMode : undefined;
            resourceInputs["cookieName"] = args ? args.cookieName : undefined;
            resourceInputs["edgeGateway"] = args ? args.edgeGateway : undefined;
            resourceInputs["enablePoolSideSsl"] = args ? args.enablePoolSideSsl : undefined;
            resourceInputs["enableSslPassthrough"] = args ? args.enableSslPassthrough : undefined;
            resourceInputs["expiration"] = args ? args.expiration : undefined;
            resourceInputs["httpRedirectUrl"] = args ? args.httpRedirectUrl : undefined;
            resourceInputs["insertXForwardedHttpHeader"] = args ? args.insertXForwardedHttpHeader : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["persistenceMechanism"] = args ? args.persistenceMechanism : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["vdc"] = args ? args.vdc : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LbAppProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LbAppProfile resources.
 */
export interface LbAppProfileState {
    /**
     * The mode by which the cookie should be inserted. One of 'insert', 'prefix', or 'appsession'
     */
    cookieMode?: pulumi.Input<string>;
    /**
     * Used to uniquely identify the session the first time a client accesses the site. The load balancer refers to this cookie
     * when connecting subsequent requests in the session, so that they all go to the same virtual server. Only applies for
     * persistence_mechanism 'cookie'
     */
    cookieName?: pulumi.Input<string>;
    /**
     * Edge gateway name in which the LB Application Profile is located
     */
    edgeGateway?: pulumi.Input<string>;
    /**
     * Enable to define the certificate, CAs, or CRLs used to authenticate the load balancer from the server side
     */
    enablePoolSideSsl?: pulumi.Input<boolean>;
    /**
     * Enable SSL authentication to be passed through to the virtual server. Otherwise SSL authentication takes place at the
     * destination address.
     */
    enableSslPassthrough?: pulumi.Input<boolean>;
    /**
     * Length of time in seconds that persistence stays in effect
     */
    expiration?: pulumi.Input<number>;
    /**
     * The URL to which traffic that arrives at the destination address should be redirected. Only applies for types 'http' and
     * 'https'
     */
    httpRedirectUrl?: pulumi.Input<string>;
    /**
     * Enables 'X-Forwarded-For' header for identifying the originating IP address of a client connecting to a Web server
     * through the load balancer. Only applies for types HTTP and HTTPS
     */
    insertXForwardedHttpHeader?: pulumi.Input<boolean>;
    /**
     * Unique LB Application Profile name
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * Persistence mechanism for the profile. One of 'cookie', 'ssl-sessionid', 'sourceip'
     */
    persistenceMechanism?: pulumi.Input<string>;
    /**
     * Protocol type used to send requests to the server. One of 'tcp', 'udp', 'http' org 'https'
     */
    type?: pulumi.Input<string>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LbAppProfile resource.
 */
export interface LbAppProfileArgs {
    /**
     * The mode by which the cookie should be inserted. One of 'insert', 'prefix', or 'appsession'
     */
    cookieMode?: pulumi.Input<string>;
    /**
     * Used to uniquely identify the session the first time a client accesses the site. The load balancer refers to this cookie
     * when connecting subsequent requests in the session, so that they all go to the same virtual server. Only applies for
     * persistence_mechanism 'cookie'
     */
    cookieName?: pulumi.Input<string>;
    /**
     * Edge gateway name in which the LB Application Profile is located
     */
    edgeGateway: pulumi.Input<string>;
    /**
     * Enable to define the certificate, CAs, or CRLs used to authenticate the load balancer from the server side
     */
    enablePoolSideSsl?: pulumi.Input<boolean>;
    /**
     * Enable SSL authentication to be passed through to the virtual server. Otherwise SSL authentication takes place at the
     * destination address.
     */
    enableSslPassthrough?: pulumi.Input<boolean>;
    /**
     * Length of time in seconds that persistence stays in effect
     */
    expiration?: pulumi.Input<number>;
    /**
     * The URL to which traffic that arrives at the destination address should be redirected. Only applies for types 'http' and
     * 'https'
     */
    httpRedirectUrl?: pulumi.Input<string>;
    /**
     * Enables 'X-Forwarded-For' header for identifying the originating IP address of a client connecting to a Web server
     * through the load balancer. Only applies for types HTTP and HTTPS
     */
    insertXForwardedHttpHeader?: pulumi.Input<boolean>;
    /**
     * Unique LB Application Profile name
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * Persistence mechanism for the profile. One of 'cookie', 'ssl-sessionid', 'sourceip'
     */
    persistenceMechanism?: pulumi.Input<string>;
    /**
     * Protocol type used to send requests to the server. One of 'tcp', 'udp', 'http' org 'https'
     */
    type: pulumi.Input<string>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: pulumi.Input<string>;
}