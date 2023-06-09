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

__all__ = ['NetworkRoutedV2Args', 'NetworkRoutedV2']

@pulumi.input_type
class NetworkRoutedV2Args:
    def __init__(__self__, *,
                 edge_gateway_id: pulumi.Input[str],
                 gateway: pulumi.Input[str],
                 prefix_length: pulumi.Input[int],
                 description: Optional[pulumi.Input[str]] = None,
                 dns1: Optional[pulumi.Input[str]] = None,
                 dns2: Optional[pulumi.Input[str]] = None,
                 dns_suffix: Optional[pulumi.Input[str]] = None,
                 interface_type: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 metadata_entries: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkRoutedV2MetadataEntryArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 static_ip_pools: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkRoutedV2StaticIpPoolArgs']]]] = None,
                 vdc: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NetworkRoutedV2 resource.
        :param pulumi.Input[str] edge_gateway_id: Edge gateway ID in which Routed network should be located
        :param pulumi.Input[str] gateway: Gateway IP address
        :param pulumi.Input[int] prefix_length: Network prefix
        :param pulumi.Input[str] description: Network description
        :param pulumi.Input[str] dns1: DNS server 1
        :param pulumi.Input[str] dns2: DNS server 1
        :param pulumi.Input[str] dns_suffix: DNS suffix
        :param pulumi.Input[str] interface_type: Optional interface type (only for NSX-V networks). One of 'INTERNAL' (default), 'DISTRIBUTED', 'SUBINTERFACE'
        :param pulumi.Input[Mapping[str, Any]] metadata: Key value map of metadata to assign to this network. Key and value can be any string
        :param pulumi.Input[Sequence[pulumi.Input['NetworkRoutedV2MetadataEntryArgs']]] metadata_entries: Metadata entries for the given Network
        :param pulumi.Input[str] name: Network name
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[Sequence[pulumi.Input['NetworkRoutedV2StaticIpPoolArgs']]] static_ip_pools: IP ranges used for static pool allocation in the network
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        pulumi.set(__self__, "edge_gateway_id", edge_gateway_id)
        pulumi.set(__self__, "gateway", gateway)
        pulumi.set(__self__, "prefix_length", prefix_length)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dns1 is not None:
            pulumi.set(__self__, "dns1", dns1)
        if dns2 is not None:
            pulumi.set(__self__, "dns2", dns2)
        if dns_suffix is not None:
            pulumi.set(__self__, "dns_suffix", dns_suffix)
        if interface_type is not None:
            pulumi.set(__self__, "interface_type", interface_type)
        if metadata is not None:
            warnings.warn("""Use metadata_entry instead""", DeprecationWarning)
            pulumi.log.warn("""metadata is deprecated: Use metadata_entry instead""")
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if metadata_entries is not None:
            pulumi.set(__self__, "metadata_entries", metadata_entries)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if static_ip_pools is not None:
            pulumi.set(__self__, "static_ip_pools", static_ip_pools)
        if vdc is not None:
            warnings.warn("""'vdc' is deprecated and ineffective. Routed networks will inherit VDC setting from parent Edge Gateway""", DeprecationWarning)
            pulumi.log.warn("""vdc is deprecated: 'vdc' is deprecated and ineffective. Routed networks will inherit VDC setting from parent Edge Gateway""")
        if vdc is not None:
            pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter(name="edgeGatewayId")
    def edge_gateway_id(self) -> pulumi.Input[str]:
        """
        Edge gateway ID in which Routed network should be located
        """
        return pulumi.get(self, "edge_gateway_id")

    @edge_gateway_id.setter
    def edge_gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "edge_gateway_id", value)

    @property
    @pulumi.getter
    def gateway(self) -> pulumi.Input[str]:
        """
        Gateway IP address
        """
        return pulumi.get(self, "gateway")

    @gateway.setter
    def gateway(self, value: pulumi.Input[str]):
        pulumi.set(self, "gateway", value)

    @property
    @pulumi.getter(name="prefixLength")
    def prefix_length(self) -> pulumi.Input[int]:
        """
        Network prefix
        """
        return pulumi.get(self, "prefix_length")

    @prefix_length.setter
    def prefix_length(self, value: pulumi.Input[int]):
        pulumi.set(self, "prefix_length", value)

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
    def dns1(self) -> Optional[pulumi.Input[str]]:
        """
        DNS server 1
        """
        return pulumi.get(self, "dns1")

    @dns1.setter
    def dns1(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns1", value)

    @property
    @pulumi.getter
    def dns2(self) -> Optional[pulumi.Input[str]]:
        """
        DNS server 1
        """
        return pulumi.get(self, "dns2")

    @dns2.setter
    def dns2(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns2", value)

    @property
    @pulumi.getter(name="dnsSuffix")
    def dns_suffix(self) -> Optional[pulumi.Input[str]]:
        """
        DNS suffix
        """
        return pulumi.get(self, "dns_suffix")

    @dns_suffix.setter
    def dns_suffix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_suffix", value)

    @property
    @pulumi.getter(name="interfaceType")
    def interface_type(self) -> Optional[pulumi.Input[str]]:
        """
        Optional interface type (only for NSX-V networks). One of 'INTERNAL' (default), 'DISTRIBUTED', 'SUBINTERFACE'
        """
        return pulumi.get(self, "interface_type")

    @interface_type.setter
    def interface_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface_type", value)

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
    def metadata_entries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkRoutedV2MetadataEntryArgs']]]]:
        """
        Metadata entries for the given Network
        """
        return pulumi.get(self, "metadata_entries")

    @metadata_entries.setter
    def metadata_entries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkRoutedV2MetadataEntryArgs']]]]):
        pulumi.set(self, "metadata_entries", value)

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
    @pulumi.getter(name="staticIpPools")
    def static_ip_pools(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkRoutedV2StaticIpPoolArgs']]]]:
        """
        IP ranges used for static pool allocation in the network
        """
        return pulumi.get(self, "static_ip_pools")

    @static_ip_pools.setter
    def static_ip_pools(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkRoutedV2StaticIpPoolArgs']]]]):
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
class _NetworkRoutedV2State:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 dns1: Optional[pulumi.Input[str]] = None,
                 dns2: Optional[pulumi.Input[str]] = None,
                 dns_suffix: Optional[pulumi.Input[str]] = None,
                 edge_gateway_id: Optional[pulumi.Input[str]] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 interface_type: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 metadata_entries: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkRoutedV2MetadataEntryArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 owner_id: Optional[pulumi.Input[str]] = None,
                 prefix_length: Optional[pulumi.Input[int]] = None,
                 static_ip_pools: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkRoutedV2StaticIpPoolArgs']]]] = None,
                 vdc: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NetworkRoutedV2 resources.
        :param pulumi.Input[str] description: Network description
        :param pulumi.Input[str] dns1: DNS server 1
        :param pulumi.Input[str] dns2: DNS server 1
        :param pulumi.Input[str] dns_suffix: DNS suffix
        :param pulumi.Input[str] edge_gateway_id: Edge gateway ID in which Routed network should be located
        :param pulumi.Input[str] gateway: Gateway IP address
        :param pulumi.Input[str] interface_type: Optional interface type (only for NSX-V networks). One of 'INTERNAL' (default), 'DISTRIBUTED', 'SUBINTERFACE'
        :param pulumi.Input[Mapping[str, Any]] metadata: Key value map of metadata to assign to this network. Key and value can be any string
        :param pulumi.Input[Sequence[pulumi.Input['NetworkRoutedV2MetadataEntryArgs']]] metadata_entries: Metadata entries for the given Network
        :param pulumi.Input[str] name: Network name
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[str] owner_id: ID of VDC or VDC Group
        :param pulumi.Input[int] prefix_length: Network prefix
        :param pulumi.Input[Sequence[pulumi.Input['NetworkRoutedV2StaticIpPoolArgs']]] static_ip_pools: IP ranges used for static pool allocation in the network
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if dns1 is not None:
            pulumi.set(__self__, "dns1", dns1)
        if dns2 is not None:
            pulumi.set(__self__, "dns2", dns2)
        if dns_suffix is not None:
            pulumi.set(__self__, "dns_suffix", dns_suffix)
        if edge_gateway_id is not None:
            pulumi.set(__self__, "edge_gateway_id", edge_gateway_id)
        if gateway is not None:
            pulumi.set(__self__, "gateway", gateway)
        if interface_type is not None:
            pulumi.set(__self__, "interface_type", interface_type)
        if metadata is not None:
            warnings.warn("""Use metadata_entry instead""", DeprecationWarning)
            pulumi.log.warn("""metadata is deprecated: Use metadata_entry instead""")
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if metadata_entries is not None:
            pulumi.set(__self__, "metadata_entries", metadata_entries)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if owner_id is not None:
            pulumi.set(__self__, "owner_id", owner_id)
        if prefix_length is not None:
            pulumi.set(__self__, "prefix_length", prefix_length)
        if static_ip_pools is not None:
            pulumi.set(__self__, "static_ip_pools", static_ip_pools)
        if vdc is not None:
            warnings.warn("""'vdc' is deprecated and ineffective. Routed networks will inherit VDC setting from parent Edge Gateway""", DeprecationWarning)
            pulumi.log.warn("""vdc is deprecated: 'vdc' is deprecated and ineffective. Routed networks will inherit VDC setting from parent Edge Gateway""")
        if vdc is not None:
            pulumi.set(__self__, "vdc", vdc)

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
    def dns1(self) -> Optional[pulumi.Input[str]]:
        """
        DNS server 1
        """
        return pulumi.get(self, "dns1")

    @dns1.setter
    def dns1(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns1", value)

    @property
    @pulumi.getter
    def dns2(self) -> Optional[pulumi.Input[str]]:
        """
        DNS server 1
        """
        return pulumi.get(self, "dns2")

    @dns2.setter
    def dns2(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns2", value)

    @property
    @pulumi.getter(name="dnsSuffix")
    def dns_suffix(self) -> Optional[pulumi.Input[str]]:
        """
        DNS suffix
        """
        return pulumi.get(self, "dns_suffix")

    @dns_suffix.setter
    def dns_suffix(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dns_suffix", value)

    @property
    @pulumi.getter(name="edgeGatewayId")
    def edge_gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        Edge gateway ID in which Routed network should be located
        """
        return pulumi.get(self, "edge_gateway_id")

    @edge_gateway_id.setter
    def edge_gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "edge_gateway_id", value)

    @property
    @pulumi.getter
    def gateway(self) -> Optional[pulumi.Input[str]]:
        """
        Gateway IP address
        """
        return pulumi.get(self, "gateway")

    @gateway.setter
    def gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gateway", value)

    @property
    @pulumi.getter(name="interfaceType")
    def interface_type(self) -> Optional[pulumi.Input[str]]:
        """
        Optional interface type (only for NSX-V networks). One of 'INTERNAL' (default), 'DISTRIBUTED', 'SUBINTERFACE'
        """
        return pulumi.get(self, "interface_type")

    @interface_type.setter
    def interface_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "interface_type", value)

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
    def metadata_entries(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkRoutedV2MetadataEntryArgs']]]]:
        """
        Metadata entries for the given Network
        """
        return pulumi.get(self, "metadata_entries")

    @metadata_entries.setter
    def metadata_entries(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkRoutedV2MetadataEntryArgs']]]]):
        pulumi.set(self, "metadata_entries", value)

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
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of VDC or VDC Group
        """
        return pulumi.get(self, "owner_id")

    @owner_id.setter
    def owner_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner_id", value)

    @property
    @pulumi.getter(name="prefixLength")
    def prefix_length(self) -> Optional[pulumi.Input[int]]:
        """
        Network prefix
        """
        return pulumi.get(self, "prefix_length")

    @prefix_length.setter
    def prefix_length(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "prefix_length", value)

    @property
    @pulumi.getter(name="staticIpPools")
    def static_ip_pools(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NetworkRoutedV2StaticIpPoolArgs']]]]:
        """
        IP ranges used for static pool allocation in the network
        """
        return pulumi.get(self, "static_ip_pools")

    @static_ip_pools.setter
    def static_ip_pools(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NetworkRoutedV2StaticIpPoolArgs']]]]):
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


class NetworkRoutedV2(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dns1: Optional[pulumi.Input[str]] = None,
                 dns2: Optional[pulumi.Input[str]] = None,
                 dns_suffix: Optional[pulumi.Input[str]] = None,
                 edge_gateway_id: Optional[pulumi.Input[str]] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 interface_type: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 metadata_entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkRoutedV2MetadataEntryArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 prefix_length: Optional[pulumi.Input[int]] = None,
                 static_ip_pools: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkRoutedV2StaticIpPoolArgs']]]]] = None,
                 vdc: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a NetworkRoutedV2 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Network description
        :param pulumi.Input[str] dns1: DNS server 1
        :param pulumi.Input[str] dns2: DNS server 1
        :param pulumi.Input[str] dns_suffix: DNS suffix
        :param pulumi.Input[str] edge_gateway_id: Edge gateway ID in which Routed network should be located
        :param pulumi.Input[str] gateway: Gateway IP address
        :param pulumi.Input[str] interface_type: Optional interface type (only for NSX-V networks). One of 'INTERNAL' (default), 'DISTRIBUTED', 'SUBINTERFACE'
        :param pulumi.Input[Mapping[str, Any]] metadata: Key value map of metadata to assign to this network. Key and value can be any string
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkRoutedV2MetadataEntryArgs']]]] metadata_entries: Metadata entries for the given Network
        :param pulumi.Input[str] name: Network name
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[int] prefix_length: Network prefix
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkRoutedV2StaticIpPoolArgs']]]] static_ip_pools: IP ranges used for static pool allocation in the network
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NetworkRoutedV2Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a NetworkRoutedV2 resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param NetworkRoutedV2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NetworkRoutedV2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 dns1: Optional[pulumi.Input[str]] = None,
                 dns2: Optional[pulumi.Input[str]] = None,
                 dns_suffix: Optional[pulumi.Input[str]] = None,
                 edge_gateway_id: Optional[pulumi.Input[str]] = None,
                 gateway: Optional[pulumi.Input[str]] = None,
                 interface_type: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 metadata_entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkRoutedV2MetadataEntryArgs']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 prefix_length: Optional[pulumi.Input[int]] = None,
                 static_ip_pools: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkRoutedV2StaticIpPoolArgs']]]]] = None,
                 vdc: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NetworkRoutedV2Args.__new__(NetworkRoutedV2Args)

            __props__.__dict__["description"] = description
            __props__.__dict__["dns1"] = dns1
            __props__.__dict__["dns2"] = dns2
            __props__.__dict__["dns_suffix"] = dns_suffix
            if edge_gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'edge_gateway_id'")
            __props__.__dict__["edge_gateway_id"] = edge_gateway_id
            if gateway is None and not opts.urn:
                raise TypeError("Missing required property 'gateway'")
            __props__.__dict__["gateway"] = gateway
            __props__.__dict__["interface_type"] = interface_type
            if metadata is not None and not opts.urn:
                warnings.warn("""Use metadata_entry instead""", DeprecationWarning)
                pulumi.log.warn("""metadata is deprecated: Use metadata_entry instead""")
            __props__.__dict__["metadata"] = metadata
            __props__.__dict__["metadata_entries"] = metadata_entries
            __props__.__dict__["name"] = name
            __props__.__dict__["org"] = org
            if prefix_length is None and not opts.urn:
                raise TypeError("Missing required property 'prefix_length'")
            __props__.__dict__["prefix_length"] = prefix_length
            __props__.__dict__["static_ip_pools"] = static_ip_pools
            if vdc is not None and not opts.urn:
                warnings.warn("""'vdc' is deprecated and ineffective. Routed networks will inherit VDC setting from parent Edge Gateway""", DeprecationWarning)
                pulumi.log.warn("""vdc is deprecated: 'vdc' is deprecated and ineffective. Routed networks will inherit VDC setting from parent Edge Gateway""")
            __props__.__dict__["vdc"] = vdc
            __props__.__dict__["owner_id"] = None
        super(NetworkRoutedV2, __self__).__init__(
            'vcd:index/networkRoutedV2:NetworkRoutedV2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            dns1: Optional[pulumi.Input[str]] = None,
            dns2: Optional[pulumi.Input[str]] = None,
            dns_suffix: Optional[pulumi.Input[str]] = None,
            edge_gateway_id: Optional[pulumi.Input[str]] = None,
            gateway: Optional[pulumi.Input[str]] = None,
            interface_type: Optional[pulumi.Input[str]] = None,
            metadata: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            metadata_entries: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkRoutedV2MetadataEntryArgs']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org: Optional[pulumi.Input[str]] = None,
            owner_id: Optional[pulumi.Input[str]] = None,
            prefix_length: Optional[pulumi.Input[int]] = None,
            static_ip_pools: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkRoutedV2StaticIpPoolArgs']]]]] = None,
            vdc: Optional[pulumi.Input[str]] = None) -> 'NetworkRoutedV2':
        """
        Get an existing NetworkRoutedV2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Network description
        :param pulumi.Input[str] dns1: DNS server 1
        :param pulumi.Input[str] dns2: DNS server 1
        :param pulumi.Input[str] dns_suffix: DNS suffix
        :param pulumi.Input[str] edge_gateway_id: Edge gateway ID in which Routed network should be located
        :param pulumi.Input[str] gateway: Gateway IP address
        :param pulumi.Input[str] interface_type: Optional interface type (only for NSX-V networks). One of 'INTERNAL' (default), 'DISTRIBUTED', 'SUBINTERFACE'
        :param pulumi.Input[Mapping[str, Any]] metadata: Key value map of metadata to assign to this network. Key and value can be any string
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkRoutedV2MetadataEntryArgs']]]] metadata_entries: Metadata entries for the given Network
        :param pulumi.Input[str] name: Network name
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[str] owner_id: ID of VDC or VDC Group
        :param pulumi.Input[int] prefix_length: Network prefix
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['NetworkRoutedV2StaticIpPoolArgs']]]] static_ip_pools: IP ranges used for static pool allocation in the network
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NetworkRoutedV2State.__new__(_NetworkRoutedV2State)

        __props__.__dict__["description"] = description
        __props__.__dict__["dns1"] = dns1
        __props__.__dict__["dns2"] = dns2
        __props__.__dict__["dns_suffix"] = dns_suffix
        __props__.__dict__["edge_gateway_id"] = edge_gateway_id
        __props__.__dict__["gateway"] = gateway
        __props__.__dict__["interface_type"] = interface_type
        __props__.__dict__["metadata"] = metadata
        __props__.__dict__["metadata_entries"] = metadata_entries
        __props__.__dict__["name"] = name
        __props__.__dict__["org"] = org
        __props__.__dict__["owner_id"] = owner_id
        __props__.__dict__["prefix_length"] = prefix_length
        __props__.__dict__["static_ip_pools"] = static_ip_pools
        __props__.__dict__["vdc"] = vdc
        return NetworkRoutedV2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Network description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def dns1(self) -> pulumi.Output[Optional[str]]:
        """
        DNS server 1
        """
        return pulumi.get(self, "dns1")

    @property
    @pulumi.getter
    def dns2(self) -> pulumi.Output[Optional[str]]:
        """
        DNS server 1
        """
        return pulumi.get(self, "dns2")

    @property
    @pulumi.getter(name="dnsSuffix")
    def dns_suffix(self) -> pulumi.Output[Optional[str]]:
        """
        DNS suffix
        """
        return pulumi.get(self, "dns_suffix")

    @property
    @pulumi.getter(name="edgeGatewayId")
    def edge_gateway_id(self) -> pulumi.Output[str]:
        """
        Edge gateway ID in which Routed network should be located
        """
        return pulumi.get(self, "edge_gateway_id")

    @property
    @pulumi.getter
    def gateway(self) -> pulumi.Output[str]:
        """
        Gateway IP address
        """
        return pulumi.get(self, "gateway")

    @property
    @pulumi.getter(name="interfaceType")
    def interface_type(self) -> pulumi.Output[Optional[str]]:
        """
        Optional interface type (only for NSX-V networks). One of 'INTERNAL' (default), 'DISTRIBUTED', 'SUBINTERFACE'
        """
        return pulumi.get(self, "interface_type")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        Key value map of metadata to assign to this network. Key and value can be any string
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter(name="metadataEntries")
    def metadata_entries(self) -> pulumi.Output[Sequence['outputs.NetworkRoutedV2MetadataEntry']]:
        """
        Metadata entries for the given Network
        """
        return pulumi.get(self, "metadata_entries")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Network name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def org(self) -> pulumi.Output[Optional[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        different organizations
        """
        return pulumi.get(self, "org")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[str]:
        """
        ID of VDC or VDC Group
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter(name="prefixLength")
    def prefix_length(self) -> pulumi.Output[int]:
        """
        Network prefix
        """
        return pulumi.get(self, "prefix_length")

    @property
    @pulumi.getter(name="staticIpPools")
    def static_ip_pools(self) -> pulumi.Output[Optional[Sequence['outputs.NetworkRoutedV2StaticIpPool']]]:
        """
        IP ranges used for static pool allocation in the network
        """
        return pulumi.get(self, "static_ip_pools")

    @property
    @pulumi.getter
    def vdc(self) -> pulumi.Output[str]:
        """
        The name of VDC to use, optional if defined at provider level
        """
        return pulumi.get(self, "vdc")

