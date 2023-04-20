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

__all__ = ['NsxvDhcpRelayArgs', 'NsxvDhcpRelay']

@pulumi.input_type
class NsxvDhcpRelayArgs:
    def __init__(__self__, *,
                 edge_gateway: pulumi.Input[str],
                 relay_agents: pulumi.Input[Sequence[pulumi.Input['NsxvDhcpRelayRelayAgentArgs']]],
                 domain_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ip_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 vdc: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NsxvDhcpRelay resource.
        :param pulumi.Input[str] edge_gateway: Edge gateway name for DHCP relay settings
        :param pulumi.Input[Sequence[pulumi.Input[str]]] domain_names: A set of IP domain names of DHCP servers
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: A set of IP address of DHCP servers
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_sets: A set of IP set names which consist DHCP servers
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        pulumi.set(__self__, "edge_gateway", edge_gateway)
        pulumi.set(__self__, "relay_agents", relay_agents)
        if domain_names is not None:
            pulumi.set(__self__, "domain_names", domain_names)
        if ip_addresses is not None:
            pulumi.set(__self__, "ip_addresses", ip_addresses)
        if ip_sets is not None:
            pulumi.set(__self__, "ip_sets", ip_sets)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if vdc is not None:
            pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter(name="edgeGateway")
    def edge_gateway(self) -> pulumi.Input[str]:
        """
        Edge gateway name for DHCP relay settings
        """
        return pulumi.get(self, "edge_gateway")

    @edge_gateway.setter
    def edge_gateway(self, value: pulumi.Input[str]):
        pulumi.set(self, "edge_gateway", value)

    @property
    @pulumi.getter(name="relayAgents")
    def relay_agents(self) -> pulumi.Input[Sequence[pulumi.Input['NsxvDhcpRelayRelayAgentArgs']]]:
        return pulumi.get(self, "relay_agents")

    @relay_agents.setter
    def relay_agents(self, value: pulumi.Input[Sequence[pulumi.Input['NsxvDhcpRelayRelayAgentArgs']]]):
        pulumi.set(self, "relay_agents", value)

    @property
    @pulumi.getter(name="domainNames")
    def domain_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of IP domain names of DHCP servers
        """
        return pulumi.get(self, "domain_names")

    @domain_names.setter
    def domain_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "domain_names", value)

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of IP address of DHCP servers
        """
        return pulumi.get(self, "ip_addresses")

    @ip_addresses.setter
    def ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_addresses", value)

    @property
    @pulumi.getter(name="ipSets")
    def ip_sets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of IP set names which consist DHCP servers
        """
        return pulumi.get(self, "ip_sets")

    @ip_sets.setter
    def ip_sets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_sets", value)

    @property
    @pulumi.getter
    def org(self) -> Optional[pulumi.Input[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        different organizations
        """
        return pulumi.get(self, "org")

    @org.setter
    def org(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org", value)

    @property
    @pulumi.getter
    def vdc(self) -> Optional[pulumi.Input[str]]:
        """
        The name of VDC to use, optional if defined at provider level
        """
        return pulumi.get(self, "vdc")

    @vdc.setter
    def vdc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdc", value)


@pulumi.input_type
class _NsxvDhcpRelayState:
    def __init__(__self__, *,
                 domain_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 edge_gateway: Optional[pulumi.Input[str]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ip_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 relay_agents: Optional[pulumi.Input[Sequence[pulumi.Input['NsxvDhcpRelayRelayAgentArgs']]]] = None,
                 vdc: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NsxvDhcpRelay resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] domain_names: A set of IP domain names of DHCP servers
        :param pulumi.Input[str] edge_gateway: Edge gateway name for DHCP relay settings
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: A set of IP address of DHCP servers
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_sets: A set of IP set names which consist DHCP servers
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        if domain_names is not None:
            pulumi.set(__self__, "domain_names", domain_names)
        if edge_gateway is not None:
            pulumi.set(__self__, "edge_gateway", edge_gateway)
        if ip_addresses is not None:
            pulumi.set(__self__, "ip_addresses", ip_addresses)
        if ip_sets is not None:
            pulumi.set(__self__, "ip_sets", ip_sets)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if relay_agents is not None:
            pulumi.set(__self__, "relay_agents", relay_agents)
        if vdc is not None:
            pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter(name="domainNames")
    def domain_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of IP domain names of DHCP servers
        """
        return pulumi.get(self, "domain_names")

    @domain_names.setter
    def domain_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "domain_names", value)

    @property
    @pulumi.getter(name="edgeGateway")
    def edge_gateway(self) -> Optional[pulumi.Input[str]]:
        """
        Edge gateway name for DHCP relay settings
        """
        return pulumi.get(self, "edge_gateway")

    @edge_gateway.setter
    def edge_gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "edge_gateway", value)

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of IP address of DHCP servers
        """
        return pulumi.get(self, "ip_addresses")

    @ip_addresses.setter
    def ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_addresses", value)

    @property
    @pulumi.getter(name="ipSets")
    def ip_sets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of IP set names which consist DHCP servers
        """
        return pulumi.get(self, "ip_sets")

    @ip_sets.setter
    def ip_sets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_sets", value)

    @property
    @pulumi.getter
    def org(self) -> Optional[pulumi.Input[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        different organizations
        """
        return pulumi.get(self, "org")

    @org.setter
    def org(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org", value)

    @property
    @pulumi.getter(name="relayAgents")
    def relay_agents(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NsxvDhcpRelayRelayAgentArgs']]]]:
        return pulumi.get(self, "relay_agents")

    @relay_agents.setter
    def relay_agents(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NsxvDhcpRelayRelayAgentArgs']]]]):
        pulumi.set(self, "relay_agents", value)

    @property
    @pulumi.getter
    def vdc(self) -> Optional[pulumi.Input[str]]:
        """
        The name of VDC to use, optional if defined at provider level
        """
        return pulumi.get(self, "vdc")

    @vdc.setter
    def vdc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdc", value)


class NsxvDhcpRelay(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 edge_gateway: Optional[pulumi.Input[str]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ip_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 relay_agents: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NsxvDhcpRelayRelayAgentArgs']]]]] = None,
                 vdc: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a NsxvDhcpRelay resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] domain_names: A set of IP domain names of DHCP servers
        :param pulumi.Input[str] edge_gateway: Edge gateway name for DHCP relay settings
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: A set of IP address of DHCP servers
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_sets: A set of IP set names which consist DHCP servers
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NsxvDhcpRelayArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a NsxvDhcpRelay resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param NsxvDhcpRelayArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NsxvDhcpRelayArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 domain_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 edge_gateway: Optional[pulumi.Input[str]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ip_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 relay_agents: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NsxvDhcpRelayRelayAgentArgs']]]]] = None,
                 vdc: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NsxvDhcpRelayArgs.__new__(NsxvDhcpRelayArgs)

            __props__.__dict__["domain_names"] = domain_names
            if edge_gateway is None and not opts.urn:
                raise TypeError("Missing required property 'edge_gateway'")
            __props__.__dict__["edge_gateway"] = edge_gateway
            __props__.__dict__["ip_addresses"] = ip_addresses
            __props__.__dict__["ip_sets"] = ip_sets
            __props__.__dict__["org"] = org
            if relay_agents is None and not opts.urn:
                raise TypeError("Missing required property 'relay_agents'")
            __props__.__dict__["relay_agents"] = relay_agents
            __props__.__dict__["vdc"] = vdc
        super(NsxvDhcpRelay, __self__).__init__(
            'vcd:index/nsxvDhcpRelay:NsxvDhcpRelay',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            domain_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            edge_gateway: Optional[pulumi.Input[str]] = None,
            ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            ip_sets: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            org: Optional[pulumi.Input[str]] = None,
            relay_agents: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NsxvDhcpRelayRelayAgentArgs']]]]] = None,
            vdc: Optional[pulumi.Input[str]] = None) -> 'NsxvDhcpRelay':
        """
        Get an existing NsxvDhcpRelay resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] domain_names: A set of IP domain names of DHCP servers
        :param pulumi.Input[str] edge_gateway: Edge gateway name for DHCP relay settings
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: A set of IP address of DHCP servers
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_sets: A set of IP set names which consist DHCP servers
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NsxvDhcpRelayState.__new__(_NsxvDhcpRelayState)

        __props__.__dict__["domain_names"] = domain_names
        __props__.__dict__["edge_gateway"] = edge_gateway
        __props__.__dict__["ip_addresses"] = ip_addresses
        __props__.__dict__["ip_sets"] = ip_sets
        __props__.__dict__["org"] = org
        __props__.__dict__["relay_agents"] = relay_agents
        __props__.__dict__["vdc"] = vdc
        return NsxvDhcpRelay(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="domainNames")
    def domain_names(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A set of IP domain names of DHCP servers
        """
        return pulumi.get(self, "domain_names")

    @property
    @pulumi.getter(name="edgeGateway")
    def edge_gateway(self) -> pulumi.Output[str]:
        """
        Edge gateway name for DHCP relay settings
        """
        return pulumi.get(self, "edge_gateway")

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A set of IP address of DHCP servers
        """
        return pulumi.get(self, "ip_addresses")

    @property
    @pulumi.getter(name="ipSets")
    def ip_sets(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A set of IP set names which consist DHCP servers
        """
        return pulumi.get(self, "ip_sets")

    @property
    @pulumi.getter
    def org(self) -> pulumi.Output[Optional[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        different organizations
        """
        return pulumi.get(self, "org")

    @property
    @pulumi.getter(name="relayAgents")
    def relay_agents(self) -> pulumi.Output[Sequence['outputs.NsxvDhcpRelayRelayAgent']]:
        return pulumi.get(self, "relay_agents")

    @property
    @pulumi.getter
    def vdc(self) -> pulumi.Output[Optional[str]]:
        """
        The name of VDC to use, optional if defined at provider level
        """
        return pulumi.get(self, "vdc")

