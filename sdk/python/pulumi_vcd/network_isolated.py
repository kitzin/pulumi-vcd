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

__all__ = ['NetworkIsolatedArgs', 'NetworkIsolated']

@pulumi.input_type
class NetworkIsolatedArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 dhcp_pools: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkIsolatedDhcpPoolArgs']]]] = None,
                 dns1: Optional[pulumi.Input[str]] = None,
                 dns2: Optional[pulumi.Input[str]] = None,
                 dns_suffix: Optional[pulumi.Input[str]] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 metadata_entries: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkIsolatedMetadataEntryArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 netmask: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 shared: Optional[pulumi.Input[bool]] = None,
                 static_ip_pools: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkIsolatedStaticIpPoolArgs']]]] = None,
                 vdc: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NetworkIsolated resource.
        :param pulumi.Input[str] description: Optional description for the network
        :param pulumi.Input[Sequence[pulumi.Input['NetworkIsolatedDhcpPoolArgs']]] dhcp_pools: A range of IPs to issue to virtual machines that don't have a static IP
        :param pulumi.Input[str] dns1: First DNS server to use
        :param pulumi.Input[str] dns2: Second DNS server to use
        :param pulumi.Input[str] dns_suffix: A FQDN for the virtual machines on this network
        :param pulumi.Input[str] gateway: The gateway for this network
        :param pulumi.Input[Mapping[str, Any]] metadata: Key value map of metadata to assign to this network. Key and value can be any string
        :param pulumi.Input[Sequence[pulumi.Input['NetworkIsolatedMetadataEntryArgs']]] metadata_entries: Metadata entries for the given Network
        :param pulumi.Input[str] name: A unique name for this network
        :param pulumi.Input[str] netmask: The netmask for the new network
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[bool] shared: Defines if this network is shared between multiple VDCs in the Org
        :param pulumi.Input[Sequence[pulumi.Input['NetworkIsolatedStaticIpPoolArgs']]] static_ip_pools: A range of IPs permitted to be used as static IPs for virtual machines
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dhcp_pools is not None:
            pulumi.set(__self__, "dhcp_pools", dhcp_pools)
        if dns1 is not None:
            pulumi.set(__self__, "dns1", dns1)
        if dns2 is not None:
            pulumi.set(__self__, "dns2", dns2)
        if dns_suffix is not None:
            pulumi.set(__self__, "dns_suffix", dns_suffix)
        if gateway is not None:
            pulumi.set(__self__, "gateway", gateway)
        if metadata is not None:
            warnings.warn("""Use metadata_entry instead""", DeprecationWarning)
            pulumi.log.warn("""metadata is deprecated: Use metadata_entry instead""")
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if metadata_entries is not None:
            pulumi.set(__self__, "metadata_entries", metadata_entries)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if netmask is not None:
            pulumi.set(__self__, "netmask", netmask)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if shared is not None:
            pulumi.set(__self__, "shared", shared)
        if static_ip_pools is not None:
            pulumi.set(__self__, "static_ip_pools", static_ip_pools)
        if vdc is not None:
            pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Optional description for the network
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dhcpPools")
    def dhcp_pools(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkIsolatedDhcpPoolArgs']]]]:
        """
        A range of IPs to issue to virtual machines that don't have a static IP
        """
        return pulumi.get(self, "dhcp_pools")

    @dhcp_pools.setter
    def dhcp_pools(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkIsolatedDhcpPoolArgs']]]]):
        pulumi.set(self, "dhcp_pools", value)

    @property
    @pulumi.getter
    def dns1(self) -> Optional[pulumi.Input[str]]:
        """
        First DNS server to use
        """
        return pulumi.get(self, "dns1")

    @dns1.setter
    def dns1(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns1", value)

    @property
    @pulumi.getter
    def dns2(self) -> Optional[pulumi.Input[str]]:
        """
        Second DNS server to use
        """
        return pulumi.get(self, "dns2")

    @dns2.setter
    def dns2(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns2", value)

    @property
    @pulumi.getter(name="dnsSuffix")
    def dns_suffix(self) -> Optional[pulumi.Input[str]]:
        """
        A FQDN for the virtual machines on this network
        """
        return pulumi.get(self, "dns_suffix")

    @dns_suffix.setter
    def dns_suffix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_suffix", value)

    @property
    @pulumi.getter
    def gateway(self) -> Optional[pulumi.Input[str]]:
        """
        The gateway for this network
        """
        return pulumi.get(self, "gateway")

    @gateway.setter
    def gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Key value map of metadata to assign to this network. Key and value can be any string
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter(name="metadataEntries")
    def metadata_entries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkIsolatedMetadataEntryArgs']]]]:
        """
        Metadata entries for the given Network
        """
        return pulumi.get(self, "metadata_entries")

    @metadata_entries.setter
    def metadata_entries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkIsolatedMetadataEntryArgs']]]]):
        pulumi.set(self, "metadata_entries", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for this network
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def netmask(self) -> Optional[pulumi.Input[str]]:
        """
        The netmask for the new network
        """
        return pulumi.get(self, "netmask")

    @netmask.setter
    def netmask(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "netmask", value)

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
    def shared(self) -> Optional[pulumi.Input[bool]]:
        """
        Defines if this network is shared between multiple VDCs in the Org
        """
        return pulumi.get(self, "shared")

    @shared.setter
    def shared(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "shared", value)

    @property
    @pulumi.getter(name="staticIpPools")
    def static_ip_pools(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkIsolatedStaticIpPoolArgs']]]]:
        """
        A range of IPs permitted to be used as static IPs for virtual machines
        """
        return pulumi.get(self, "static_ip_pools")

    @static_ip_pools.setter
    def static_ip_pools(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkIsolatedStaticIpPoolArgs']]]]):
        pulumi.set(self, "static_ip_pools", value)

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
class _NetworkIsolatedState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 dhcp_pools: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkIsolatedDhcpPoolArgs']]]] = None,
                 dns1: Optional[pulumi.Input[str]] = None,
                 dns2: Optional[pulumi.Input[str]] = None,
                 dns_suffix: Optional[pulumi.Input[str]] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 href: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 metadata_entries: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkIsolatedMetadataEntryArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 netmask: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 shared: Optional[pulumi.Input[bool]] = None,
                 static_ip_pools: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkIsolatedStaticIpPoolArgs']]]] = None,
                 vdc: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NetworkIsolated resources.
        :param pulumi.Input[str] description: Optional description for the network
        :param pulumi.Input[Sequence[pulumi.Input['NetworkIsolatedDhcpPoolArgs']]] dhcp_pools: A range of IPs to issue to virtual machines that don't have a static IP
        :param pulumi.Input[str] dns1: First DNS server to use
        :param pulumi.Input[str] dns2: Second DNS server to use
        :param pulumi.Input[str] dns_suffix: A FQDN for the virtual machines on this network
        :param pulumi.Input[str] gateway: The gateway for this network
        :param pulumi.Input[str] href: Network Hyper Reference
        :param pulumi.Input[Mapping[str, Any]] metadata: Key value map of metadata to assign to this network. Key and value can be any string
        :param pulumi.Input[Sequence[pulumi.Input['NetworkIsolatedMetadataEntryArgs']]] metadata_entries: Metadata entries for the given Network
        :param pulumi.Input[str] name: A unique name for this network
        :param pulumi.Input[str] netmask: The netmask for the new network
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[bool] shared: Defines if this network is shared between multiple VDCs in the Org
        :param pulumi.Input[Sequence[pulumi.Input['NetworkIsolatedStaticIpPoolArgs']]] static_ip_pools: A range of IPs permitted to be used as static IPs for virtual machines
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dhcp_pools is not None:
            pulumi.set(__self__, "dhcp_pools", dhcp_pools)
        if dns1 is not None:
            pulumi.set(__self__, "dns1", dns1)
        if dns2 is not None:
            pulumi.set(__self__, "dns2", dns2)
        if dns_suffix is not None:
            pulumi.set(__self__, "dns_suffix", dns_suffix)
        if gateway is not None:
            pulumi.set(__self__, "gateway", gateway)
        if href is not None:
            pulumi.set(__self__, "href", href)
        if metadata is not None:
            warnings.warn("""Use metadata_entry instead""", DeprecationWarning)
            pulumi.log.warn("""metadata is deprecated: Use metadata_entry instead""")
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if metadata_entries is not None:
            pulumi.set(__self__, "metadata_entries", metadata_entries)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if netmask is not None:
            pulumi.set(__self__, "netmask", netmask)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if shared is not None:
            pulumi.set(__self__, "shared", shared)
        if static_ip_pools is not None:
            pulumi.set(__self__, "static_ip_pools", static_ip_pools)
        if vdc is not None:
            pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Optional description for the network
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="dhcpPools")
    def dhcp_pools(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkIsolatedDhcpPoolArgs']]]]:
        """
        A range of IPs to issue to virtual machines that don't have a static IP
        """
        return pulumi.get(self, "dhcp_pools")

    @dhcp_pools.setter
    def dhcp_pools(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkIsolatedDhcpPoolArgs']]]]):
        pulumi.set(self, "dhcp_pools", value)

    @property
    @pulumi.getter
    def dns1(self) -> Optional[pulumi.Input[str]]:
        """
        First DNS server to use
        """
        return pulumi.get(self, "dns1")

    @dns1.setter
    def dns1(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns1", value)

    @property
    @pulumi.getter
    def dns2(self) -> Optional[pulumi.Input[str]]:
        """
        Second DNS server to use
        """
        return pulumi.get(self, "dns2")

    @dns2.setter
    def dns2(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns2", value)

    @property
    @pulumi.getter(name="dnsSuffix")
    def dns_suffix(self) -> Optional[pulumi.Input[str]]:
        """
        A FQDN for the virtual machines on this network
        """
        return pulumi.get(self, "dns_suffix")

    @dns_suffix.setter
    def dns_suffix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_suffix", value)

    @property
    @pulumi.getter
    def gateway(self) -> Optional[pulumi.Input[str]]:
        """
        The gateway for this network
        """
        return pulumi.get(self, "gateway")

    @gateway.setter
    def gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway", value)

    @property
    @pulumi.getter
    def href(self) -> Optional[pulumi.Input[str]]:
        """
        Network Hyper Reference
        """
        return pulumi.get(self, "href")

    @href.setter
    def href(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "href", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Key value map of metadata to assign to this network. Key and value can be any string
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter(name="metadataEntries")
    def metadata_entries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkIsolatedMetadataEntryArgs']]]]:
        """
        Metadata entries for the given Network
        """
        return pulumi.get(self, "metadata_entries")

    @metadata_entries.setter
    def metadata_entries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkIsolatedMetadataEntryArgs']]]]):
        pulumi.set(self, "metadata_entries", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for this network
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def netmask(self) -> Optional[pulumi.Input[str]]:
        """
        The netmask for the new network
        """
        return pulumi.get(self, "netmask")

    @netmask.setter
    def netmask(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "netmask", value)

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
    def shared(self) -> Optional[pulumi.Input[bool]]:
        """
        Defines if this network is shared between multiple VDCs in the Org
        """
        return pulumi.get(self, "shared")

    @shared.setter
    def shared(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "shared", value)

    @property
    @pulumi.getter(name="staticIpPools")
    def static_ip_pools(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkIsolatedStaticIpPoolArgs']]]]:
        """
        A range of IPs permitted to be used as static IPs for virtual machines
        """
        return pulumi.get(self, "static_ip_pools")

    @static_ip_pools.setter
    def static_ip_pools(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkIsolatedStaticIpPoolArgs']]]]):
        pulumi.set(self, "static_ip_pools", value)

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


class NetworkIsolated(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dhcp_pools: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkIsolatedDhcpPoolArgs']]]]] = None,
                 dns1: Optional[pulumi.Input[str]] = None,
                 dns2: Optional[pulumi.Input[str]] = None,
                 dns_suffix: Optional[pulumi.Input[str]] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 metadata_entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkIsolatedMetadataEntryArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 netmask: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 shared: Optional[pulumi.Input[bool]] = None,
                 static_ip_pools: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkIsolatedStaticIpPoolArgs']]]]] = None,
                 vdc: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a NetworkIsolated resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Optional description for the network
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkIsolatedDhcpPoolArgs']]]] dhcp_pools: A range of IPs to issue to virtual machines that don't have a static IP
        :param pulumi.Input[str] dns1: First DNS server to use
        :param pulumi.Input[str] dns2: Second DNS server to use
        :param pulumi.Input[str] dns_suffix: A FQDN for the virtual machines on this network
        :param pulumi.Input[str] gateway: The gateway for this network
        :param pulumi.Input[Mapping[str, Any]] metadata: Key value map of metadata to assign to this network. Key and value can be any string
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkIsolatedMetadataEntryArgs']]]] metadata_entries: Metadata entries for the given Network
        :param pulumi.Input[str] name: A unique name for this network
        :param pulumi.Input[str] netmask: The netmask for the new network
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[bool] shared: Defines if this network is shared between multiple VDCs in the Org
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkIsolatedStaticIpPoolArgs']]]] static_ip_pools: A range of IPs permitted to be used as static IPs for virtual machines
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[NetworkIsolatedArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a NetworkIsolated resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param NetworkIsolatedArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetworkIsolatedArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dhcp_pools: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkIsolatedDhcpPoolArgs']]]]] = None,
                 dns1: Optional[pulumi.Input[str]] = None,
                 dns2: Optional[pulumi.Input[str]] = None,
                 dns_suffix: Optional[pulumi.Input[str]] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 metadata_entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkIsolatedMetadataEntryArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 netmask: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 shared: Optional[pulumi.Input[bool]] = None,
                 static_ip_pools: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkIsolatedStaticIpPoolArgs']]]]] = None,
                 vdc: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NetworkIsolatedArgs.__new__(NetworkIsolatedArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["dhcp_pools"] = dhcp_pools
            __props__.__dict__["dns1"] = dns1
            __props__.__dict__["dns2"] = dns2
            __props__.__dict__["dns_suffix"] = dns_suffix
            __props__.__dict__["gateway"] = gateway
            if metadata is not None and not opts.urn:
                warnings.warn("""Use metadata_entry instead""", DeprecationWarning)
                pulumi.log.warn("""metadata is deprecated: Use metadata_entry instead""")
            __props__.__dict__["metadata"] = metadata
            __props__.__dict__["metadata_entries"] = metadata_entries
            __props__.__dict__["name"] = name
            __props__.__dict__["netmask"] = netmask
            __props__.__dict__["org"] = org
            __props__.__dict__["shared"] = shared
            __props__.__dict__["static_ip_pools"] = static_ip_pools
            __props__.__dict__["vdc"] = vdc
            __props__.__dict__["href"] = None
        super(NetworkIsolated, __self__).__init__(
            'vcd:index/networkIsolated:NetworkIsolated',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            dhcp_pools: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkIsolatedDhcpPoolArgs']]]]] = None,
            dns1: Optional[pulumi.Input[str]] = None,
            dns2: Optional[pulumi.Input[str]] = None,
            dns_suffix: Optional[pulumi.Input[str]] = None,
            gateway: Optional[pulumi.Input[str]] = None,
            href: Optional[pulumi.Input[str]] = None,
            metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            metadata_entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkIsolatedMetadataEntryArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            netmask: Optional[pulumi.Input[str]] = None,
            org: Optional[pulumi.Input[str]] = None,
            shared: Optional[pulumi.Input[bool]] = None,
            static_ip_pools: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkIsolatedStaticIpPoolArgs']]]]] = None,
            vdc: Optional[pulumi.Input[str]] = None) -> 'NetworkIsolated':
        """
        Get an existing NetworkIsolated resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Optional description for the network
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkIsolatedDhcpPoolArgs']]]] dhcp_pools: A range of IPs to issue to virtual machines that don't have a static IP
        :param pulumi.Input[str] dns1: First DNS server to use
        :param pulumi.Input[str] dns2: Second DNS server to use
        :param pulumi.Input[str] dns_suffix: A FQDN for the virtual machines on this network
        :param pulumi.Input[str] gateway: The gateway for this network
        :param pulumi.Input[str] href: Network Hyper Reference
        :param pulumi.Input[Mapping[str, Any]] metadata: Key value map of metadata to assign to this network. Key and value can be any string
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkIsolatedMetadataEntryArgs']]]] metadata_entries: Metadata entries for the given Network
        :param pulumi.Input[str] name: A unique name for this network
        :param pulumi.Input[str] netmask: The netmask for the new network
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[bool] shared: Defines if this network is shared between multiple VDCs in the Org
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkIsolatedStaticIpPoolArgs']]]] static_ip_pools: A range of IPs permitted to be used as static IPs for virtual machines
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NetworkIsolatedState.__new__(_NetworkIsolatedState)

        __props__.__dict__["description"] = description
        __props__.__dict__["dhcp_pools"] = dhcp_pools
        __props__.__dict__["dns1"] = dns1
        __props__.__dict__["dns2"] = dns2
        __props__.__dict__["dns_suffix"] = dns_suffix
        __props__.__dict__["gateway"] = gateway
        __props__.__dict__["href"] = href
        __props__.__dict__["metadata"] = metadata
        __props__.__dict__["metadata_entries"] = metadata_entries
        __props__.__dict__["name"] = name
        __props__.__dict__["netmask"] = netmask
        __props__.__dict__["org"] = org
        __props__.__dict__["shared"] = shared
        __props__.__dict__["static_ip_pools"] = static_ip_pools
        __props__.__dict__["vdc"] = vdc
        return NetworkIsolated(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Optional description for the network
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="dhcpPools")
    def dhcp_pools(self) -> pulumi.Output[Optional[Sequence['outputs.NetworkIsolatedDhcpPool']]]:
        """
        A range of IPs to issue to virtual machines that don't have a static IP
        """
        return pulumi.get(self, "dhcp_pools")

    @property
    @pulumi.getter
    def dns1(self) -> pulumi.Output[Optional[str]]:
        """
        First DNS server to use
        """
        return pulumi.get(self, "dns1")

    @property
    @pulumi.getter
    def dns2(self) -> pulumi.Output[Optional[str]]:
        """
        Second DNS server to use
        """
        return pulumi.get(self, "dns2")

    @property
    @pulumi.getter(name="dnsSuffix")
    def dns_suffix(self) -> pulumi.Output[Optional[str]]:
        """
        A FQDN for the virtual machines on this network
        """
        return pulumi.get(self, "dns_suffix")

    @property
    @pulumi.getter
    def gateway(self) -> pulumi.Output[Optional[str]]:
        """
        The gateway for this network
        """
        return pulumi.get(self, "gateway")

    @property
    @pulumi.getter
    def href(self) -> pulumi.Output[str]:
        """
        Network Hyper Reference
        """
        return pulumi.get(self, "href")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        Key value map of metadata to assign to this network. Key and value can be any string
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter(name="metadataEntries")
    def metadata_entries(self) -> pulumi.Output[Sequence['outputs.NetworkIsolatedMetadataEntry']]:
        """
        Metadata entries for the given Network
        """
        return pulumi.get(self, "metadata_entries")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A unique name for this network
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def netmask(self) -> pulumi.Output[Optional[str]]:
        """
        The netmask for the new network
        """
        return pulumi.get(self, "netmask")

    @property
    @pulumi.getter
    def org(self) -> pulumi.Output[Optional[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        different organizations
        """
        return pulumi.get(self, "org")

    @property
    @pulumi.getter
    def shared(self) -> pulumi.Output[Optional[bool]]:
        """
        Defines if this network is shared between multiple VDCs in the Org
        """
        return pulumi.get(self, "shared")

    @property
    @pulumi.getter(name="staticIpPools")
    def static_ip_pools(self) -> pulumi.Output[Optional[Sequence['outputs.NetworkIsolatedStaticIpPool']]]:
        """
        A range of IPs permitted to be used as static IPs for virtual machines
        """
        return pulumi.get(self, "static_ip_pools")

    @property
    @pulumi.getter
    def vdc(self) -> pulumi.Output[Optional[str]]:
        """
        The name of VDC to use, optional if defined at provider level
        """
        return pulumi.get(self, "vdc")

