# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ExternalNetworkV2Args', 'ExternalNetworkV2']

@pulumi.input_type
class ExternalNetworkV2Args:
    def __init__(__self__, *,
                 ip_scopes: pulumi.Input[Sequence[pulumi.Input['ExternalNetworkV2IpScopeArgs']]],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nsxt_network: Optional[pulumi.Input['ExternalNetworkV2NsxtNetworkArgs']] = None,
                 vsphere_networks: Optional[pulumi.Input[Sequence[pulumi.Input['ExternalNetworkV2VsphereNetworkArgs']]]] = None):
        """
        The set of arguments for constructing a ExternalNetworkV2 resource.
        :param pulumi.Input[Sequence[pulumi.Input['ExternalNetworkV2IpScopeArgs']]] ip_scopes: A set of IP scopes for the network
        :param pulumi.Input[str] description: Network description
        :param pulumi.Input[str] name: Network name
        :param pulumi.Input['ExternalNetworkV2NsxtNetworkArgs'] nsxt_network: Reference to NSX-T Tier-0 router or segment and manager
        :param pulumi.Input[Sequence[pulumi.Input['ExternalNetworkV2VsphereNetworkArgs']]] vsphere_networks: A set of port groups that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server
               registered with the system.
        """
        pulumi.set(__self__, "ip_scopes", ip_scopes)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nsxt_network is not None:
            pulumi.set(__self__, "nsxt_network", nsxt_network)
        if vsphere_networks is not None:
            pulumi.set(__self__, "vsphere_networks", vsphere_networks)

    @property
    @pulumi.getter(name="ipScopes")
    def ip_scopes(self) -> pulumi.Input[Sequence[pulumi.Input['ExternalNetworkV2IpScopeArgs']]]:
        """
        A set of IP scopes for the network
        """
        return pulumi.get(self, "ip_scopes")

    @ip_scopes.setter
    def ip_scopes(self, value: pulumi.Input[Sequence[pulumi.Input['ExternalNetworkV2IpScopeArgs']]]):
        pulumi.set(self, "ip_scopes", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Network description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Network name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nsxtNetwork")
    def nsxt_network(self) -> Optional[pulumi.Input['ExternalNetworkV2NsxtNetworkArgs']]:
        """
        Reference to NSX-T Tier-0 router or segment and manager
        """
        return pulumi.get(self, "nsxt_network")

    @nsxt_network.setter
    def nsxt_network(self, value: Optional[pulumi.Input['ExternalNetworkV2NsxtNetworkArgs']]):
        pulumi.set(self, "nsxt_network", value)

    @property
    @pulumi.getter(name="vsphereNetworks")
    def vsphere_networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ExternalNetworkV2VsphereNetworkArgs']]]]:
        """
        A set of port groups that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server
        registered with the system.
        """
        return pulumi.get(self, "vsphere_networks")

    @vsphere_networks.setter
    def vsphere_networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ExternalNetworkV2VsphereNetworkArgs']]]]):
        pulumi.set(self, "vsphere_networks", value)


@pulumi.input_type
class _ExternalNetworkV2State:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_scopes: Optional[pulumi.Input[Sequence[pulumi.Input['ExternalNetworkV2IpScopeArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nsxt_network: Optional[pulumi.Input['ExternalNetworkV2NsxtNetworkArgs']] = None,
                 vsphere_networks: Optional[pulumi.Input[Sequence[pulumi.Input['ExternalNetworkV2VsphereNetworkArgs']]]] = None):
        """
        Input properties used for looking up and filtering ExternalNetworkV2 resources.
        :param pulumi.Input[str] description: Network description
        :param pulumi.Input[Sequence[pulumi.Input['ExternalNetworkV2IpScopeArgs']]] ip_scopes: A set of IP scopes for the network
        :param pulumi.Input[str] name: Network name
        :param pulumi.Input['ExternalNetworkV2NsxtNetworkArgs'] nsxt_network: Reference to NSX-T Tier-0 router or segment and manager
        :param pulumi.Input[Sequence[pulumi.Input['ExternalNetworkV2VsphereNetworkArgs']]] vsphere_networks: A set of port groups that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server
               registered with the system.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ip_scopes is not None:
            pulumi.set(__self__, "ip_scopes", ip_scopes)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if nsxt_network is not None:
            pulumi.set(__self__, "nsxt_network", nsxt_network)
        if vsphere_networks is not None:
            pulumi.set(__self__, "vsphere_networks", vsphere_networks)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Network description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ipScopes")
    def ip_scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ExternalNetworkV2IpScopeArgs']]]]:
        """
        A set of IP scopes for the network
        """
        return pulumi.get(self, "ip_scopes")

    @ip_scopes.setter
    def ip_scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ExternalNetworkV2IpScopeArgs']]]]):
        pulumi.set(self, "ip_scopes", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Network name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="nsxtNetwork")
    def nsxt_network(self) -> Optional[pulumi.Input['ExternalNetworkV2NsxtNetworkArgs']]:
        """
        Reference to NSX-T Tier-0 router or segment and manager
        """
        return pulumi.get(self, "nsxt_network")

    @nsxt_network.setter
    def nsxt_network(self, value: Optional[pulumi.Input['ExternalNetworkV2NsxtNetworkArgs']]):
        pulumi.set(self, "nsxt_network", value)

    @property
    @pulumi.getter(name="vsphereNetworks")
    def vsphere_networks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ExternalNetworkV2VsphereNetworkArgs']]]]:
        """
        A set of port groups that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server
        registered with the system.
        """
        return pulumi.get(self, "vsphere_networks")

    @vsphere_networks.setter
    def vsphere_networks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ExternalNetworkV2VsphereNetworkArgs']]]]):
        pulumi.set(self, "vsphere_networks", value)


class ExternalNetworkV2(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExternalNetworkV2IpScopeArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nsxt_network: Optional[pulumi.Input[pulumi.InputType['ExternalNetworkV2NsxtNetworkArgs']]] = None,
                 vsphere_networks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExternalNetworkV2VsphereNetworkArgs']]]]] = None,
                 __props__=None):
        """
        Create a ExternalNetworkV2 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Network description
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExternalNetworkV2IpScopeArgs']]]] ip_scopes: A set of IP scopes for the network
        :param pulumi.Input[str] name: Network name
        :param pulumi.Input[pulumi.InputType['ExternalNetworkV2NsxtNetworkArgs']] nsxt_network: Reference to NSX-T Tier-0 router or segment and manager
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExternalNetworkV2VsphereNetworkArgs']]]] vsphere_networks: A set of port groups that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server
               registered with the system.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ExternalNetworkV2Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ExternalNetworkV2 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ExternalNetworkV2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ExternalNetworkV2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExternalNetworkV2IpScopeArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 nsxt_network: Optional[pulumi.Input[pulumi.InputType['ExternalNetworkV2NsxtNetworkArgs']]] = None,
                 vsphere_networks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExternalNetworkV2VsphereNetworkArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ExternalNetworkV2Args.__new__(ExternalNetworkV2Args)

            __props__.__dict__["description"] = description
            if ip_scopes is None and not opts.urn:
                raise TypeError("Missing required property 'ip_scopes'")
            __props__.__dict__["ip_scopes"] = ip_scopes
            __props__.__dict__["name"] = name
            __props__.__dict__["nsxt_network"] = nsxt_network
            __props__.__dict__["vsphere_networks"] = vsphere_networks
        super(ExternalNetworkV2, __self__).__init__(
            'vcd:index/externalNetworkV2:ExternalNetworkV2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            ip_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExternalNetworkV2IpScopeArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            nsxt_network: Optional[pulumi.Input[pulumi.InputType['ExternalNetworkV2NsxtNetworkArgs']]] = None,
            vsphere_networks: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExternalNetworkV2VsphereNetworkArgs']]]]] = None) -> 'ExternalNetworkV2':
        """
        Get an existing ExternalNetworkV2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Network description
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExternalNetworkV2IpScopeArgs']]]] ip_scopes: A set of IP scopes for the network
        :param pulumi.Input[str] name: Network name
        :param pulumi.Input[pulumi.InputType['ExternalNetworkV2NsxtNetworkArgs']] nsxt_network: Reference to NSX-T Tier-0 router or segment and manager
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['ExternalNetworkV2VsphereNetworkArgs']]]] vsphere_networks: A set of port groups that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server
               registered with the system.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ExternalNetworkV2State.__new__(_ExternalNetworkV2State)

        __props__.__dict__["description"] = description
        __props__.__dict__["ip_scopes"] = ip_scopes
        __props__.__dict__["name"] = name
        __props__.__dict__["nsxt_network"] = nsxt_network
        __props__.__dict__["vsphere_networks"] = vsphere_networks
        return ExternalNetworkV2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Network description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ipScopes")
    def ip_scopes(self) -> pulumi.Output[Sequence['outputs.ExternalNetworkV2IpScope']]:
        """
        A set of IP scopes for the network
        """
        return pulumi.get(self, "ip_scopes")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Network name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nsxtNetwork")
    def nsxt_network(self) -> pulumi.Output[Optional['outputs.ExternalNetworkV2NsxtNetwork']]:
        """
        Reference to NSX-T Tier-0 router or segment and manager
        """
        return pulumi.get(self, "nsxt_network")

    @property
    @pulumi.getter(name="vsphereNetworks")
    def vsphere_networks(self) -> pulumi.Output[Optional[Sequence['outputs.ExternalNetworkV2VsphereNetwork']]]:
        """
        A set of port groups that back this network. Each referenced DV_PORTGROUP or NETWORK must exist on a vCenter server
        registered with the system.
        """
        return pulumi.get(self, "vsphere_networks")
