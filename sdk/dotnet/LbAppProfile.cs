// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/lbAppProfile:LbAppProfile")]
    public partial class LbAppProfile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The mode by which the cookie should be inserted. One of 'insert', 'prefix', or 'appsession'
        /// </summary>
        [Output("cookieMode")]
        public Output<string?> CookieMode { get; private set; } = null!;

        /// <summary>
        /// Used to uniquely identify the session the first time a client accesses the site. The load balancer refers to this cookie
        /// when connecting subsequent requests in the session, so that they all go to the same virtual server. Only applies for
        /// persistence_mechanism 'cookie'
        /// </summary>
        [Output("cookieName")]
        public Output<string?> CookieName { get; private set; } = null!;

        /// <summary>
        /// Edge gateway name in which the LB Application Profile is located
        /// </summary>
        [Output("edgeGateway")]
        public Output<string> EdgeGateway { get; private set; } = null!;

        /// <summary>
        /// Enable to define the certificate, CAs, or CRLs used to authenticate the load balancer from the server side
        /// </summary>
        [Output("enablePoolSideSsl")]
        public Output<bool?> EnablePoolSideSsl { get; private set; } = null!;

        /// <summary>
        /// Enable SSL authentication to be passed through to the virtual server. Otherwise SSL authentication takes place at the
        /// destination address.
        /// </summary>
        [Output("enableSslPassthrough")]
        public Output<bool?> EnableSslPassthrough { get; private set; } = null!;

        /// <summary>
        /// Length of time in seconds that persistence stays in effect
        /// </summary>
        [Output("expiration")]
        public Output<int?> Expiration { get; private set; } = null!;

        /// <summary>
        /// The URL to which traffic that arrives at the destination address should be redirected. Only applies for types 'http' and
        /// 'https'
        /// </summary>
        [Output("httpRedirectUrl")]
        public Output<string?> HttpRedirectUrl { get; private set; } = null!;

        /// <summary>
        /// Enables 'X-Forwarded-For' header for identifying the originating IP address of a client connecting to a Web server
        /// through the load balancer. Only applies for types HTTP and HTTPS
        /// </summary>
        [Output("insertXForwardedHttpHeader")]
        public Output<bool?> InsertXForwardedHttpHeader { get; private set; } = null!;

        /// <summary>
        /// Unique LB Application Profile name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        /// <summary>
        /// Persistence mechanism for the profile. One of 'cookie', 'ssl-sessionid', 'sourceip'
        /// </summary>
        [Output("persistenceMechanism")]
        public Output<string?> PersistenceMechanism { get; private set; } = null!;

        /// <summary>
        /// Protocol type used to send requests to the server. One of 'tcp', 'udp', 'http' org 'https'
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Output("vdc")]
        public Output<string?> Vdc { get; private set; } = null!;


        /// <summary>
        /// Create a LbAppProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LbAppProfile(string name, LbAppProfileArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/lbAppProfile:LbAppProfile", name, args ?? new LbAppProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LbAppProfile(string name, Input<string> id, LbAppProfileState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/lbAppProfile:LbAppProfile", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LbAppProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LbAppProfile Get(string name, Input<string> id, LbAppProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new LbAppProfile(name, id, state, options);
        }
    }

    public sealed class LbAppProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The mode by which the cookie should be inserted. One of 'insert', 'prefix', or 'appsession'
        /// </summary>
        [Input("cookieMode")]
        public Input<string>? CookieMode { get; set; }

        /// <summary>
        /// Used to uniquely identify the session the first time a client accesses the site. The load balancer refers to this cookie
        /// when connecting subsequent requests in the session, so that they all go to the same virtual server. Only applies for
        /// persistence_mechanism 'cookie'
        /// </summary>
        [Input("cookieName")]
        public Input<string>? CookieName { get; set; }

        /// <summary>
        /// Edge gateway name in which the LB Application Profile is located
        /// </summary>
        [Input("edgeGateway", required: true)]
        public Input<string> EdgeGateway { get; set; } = null!;

        /// <summary>
        /// Enable to define the certificate, CAs, or CRLs used to authenticate the load balancer from the server side
        /// </summary>
        [Input("enablePoolSideSsl")]
        public Input<bool>? EnablePoolSideSsl { get; set; }

        /// <summary>
        /// Enable SSL authentication to be passed through to the virtual server. Otherwise SSL authentication takes place at the
        /// destination address.
        /// </summary>
        [Input("enableSslPassthrough")]
        public Input<bool>? EnableSslPassthrough { get; set; }

        /// <summary>
        /// Length of time in seconds that persistence stays in effect
        /// </summary>
        [Input("expiration")]
        public Input<int>? Expiration { get; set; }

        /// <summary>
        /// The URL to which traffic that arrives at the destination address should be redirected. Only applies for types 'http' and
        /// 'https'
        /// </summary>
        [Input("httpRedirectUrl")]
        public Input<string>? HttpRedirectUrl { get; set; }

        /// <summary>
        /// Enables 'X-Forwarded-For' header for identifying the originating IP address of a client connecting to a Web server
        /// through the load balancer. Only applies for types HTTP and HTTPS
        /// </summary>
        [Input("insertXForwardedHttpHeader")]
        public Input<bool>? InsertXForwardedHttpHeader { get; set; }

        /// <summary>
        /// Unique LB Application Profile name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// Persistence mechanism for the profile. One of 'cookie', 'ssl-sessionid', 'sourceip'
        /// </summary>
        [Input("persistenceMechanism")]
        public Input<string>? PersistenceMechanism { get; set; }

        /// <summary>
        /// Protocol type used to send requests to the server. One of 'tcp', 'udp', 'http' org 'https'
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public LbAppProfileArgs()
        {
        }
        public static new LbAppProfileArgs Empty => new LbAppProfileArgs();
    }

    public sealed class LbAppProfileState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The mode by which the cookie should be inserted. One of 'insert', 'prefix', or 'appsession'
        /// </summary>
        [Input("cookieMode")]
        public Input<string>? CookieMode { get; set; }

        /// <summary>
        /// Used to uniquely identify the session the first time a client accesses the site. The load balancer refers to this cookie
        /// when connecting subsequent requests in the session, so that they all go to the same virtual server. Only applies for
        /// persistence_mechanism 'cookie'
        /// </summary>
        [Input("cookieName")]
        public Input<string>? CookieName { get; set; }

        /// <summary>
        /// Edge gateway name in which the LB Application Profile is located
        /// </summary>
        [Input("edgeGateway")]
        public Input<string>? EdgeGateway { get; set; }

        /// <summary>
        /// Enable to define the certificate, CAs, or CRLs used to authenticate the load balancer from the server side
        /// </summary>
        [Input("enablePoolSideSsl")]
        public Input<bool>? EnablePoolSideSsl { get; set; }

        /// <summary>
        /// Enable SSL authentication to be passed through to the virtual server. Otherwise SSL authentication takes place at the
        /// destination address.
        /// </summary>
        [Input("enableSslPassthrough")]
        public Input<bool>? EnableSslPassthrough { get; set; }

        /// <summary>
        /// Length of time in seconds that persistence stays in effect
        /// </summary>
        [Input("expiration")]
        public Input<int>? Expiration { get; set; }

        /// <summary>
        /// The URL to which traffic that arrives at the destination address should be redirected. Only applies for types 'http' and
        /// 'https'
        /// </summary>
        [Input("httpRedirectUrl")]
        public Input<string>? HttpRedirectUrl { get; set; }

        /// <summary>
        /// Enables 'X-Forwarded-For' header for identifying the originating IP address of a client connecting to a Web server
        /// through the load balancer. Only applies for types HTTP and HTTPS
        /// </summary>
        [Input("insertXForwardedHttpHeader")]
        public Input<bool>? InsertXForwardedHttpHeader { get; set; }

        /// <summary>
        /// Unique LB Application Profile name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// Persistence mechanism for the profile. One of 'cookie', 'ssl-sessionid', 'sourceip'
        /// </summary>
        [Input("persistenceMechanism")]
        public Input<string>? PersistenceMechanism { get; set; }

        /// <summary>
        /// Protocol type used to send requests to the server. One of 'tcp', 'udp', 'http' org 'https'
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public LbAppProfileState()
        {
        }
        public static new LbAppProfileState Empty => new LbAppProfileState();
    }
}