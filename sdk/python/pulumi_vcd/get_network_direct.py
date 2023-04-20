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

__all__ = [
    'GetNetworkDirectResult',
    'AwaitableGetNetworkDirectResult',
    'get_network_direct',
    'get_network_direct_output',
]

@pulumi.output_type
class GetNetworkDirectResult:
    """
    A collection of values returned by getNetworkDirect.
    """
    def __init__(__self__, description=None, external_network=None, external_network_dns1=None, external_network_dns2=None, external_network_dns_suffix=None, external_network_gateway=None, external_network_netmask=None, filter=None, href=None, id=None, metadata=None, metadata_entries=None, name=None, org=None, shared=None, vdc=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if external_network and not isinstance(external_network, str):
            raise TypeError("Expected argument 'external_network' to be a str")
        pulumi.set(__self__, "external_network", external_network)
        if external_network_dns1 and not isinstance(external_network_dns1, str):
            raise TypeError("Expected argument 'external_network_dns1' to be a str")
        pulumi.set(__self__, "external_network_dns1", external_network_dns1)
        if external_network_dns2 and not isinstance(external_network_dns2, str):
            raise TypeError("Expected argument 'external_network_dns2' to be a str")
        pulumi.set(__self__, "external_network_dns2", external_network_dns2)
        if external_network_dns_suffix and not isinstance(external_network_dns_suffix, str):
            raise TypeError("Expected argument 'external_network_dns_suffix' to be a str")
        pulumi.set(__self__, "external_network_dns_suffix", external_network_dns_suffix)
        if external_network_gateway and not isinstance(external_network_gateway, str):
            raise TypeError("Expected argument 'external_network_gateway' to be a str")
        pulumi.set(__self__, "external_network_gateway", external_network_gateway)
        if external_network_netmask and not isinstance(external_network_netmask, str):
            raise TypeError("Expected argument 'external_network_netmask' to be a str")
        pulumi.set(__self__, "external_network_netmask", external_network_netmask)
        if filter and not isinstance(filter, dict):
            raise TypeError("Expected argument 'filter' to be a dict")
        pulumi.set(__self__, "filter", filter)
        if href and not isinstance(href, str):
            raise TypeError("Expected argument 'href' to be a str")
        pulumi.set(__self__, "href", href)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        if metadata is not None:
            warnings.warn("""Use metadata_entry instead""", DeprecationWarning)
            pulumi.log.warn("""metadata is deprecated: Use metadata_entry instead""")

        pulumi.set(__self__, "metadata", metadata)
        if metadata_entries and not isinstance(metadata_entries, list):
            raise TypeError("Expected argument 'metadata_entries' to be a list")
        pulumi.set(__self__, "metadata_entries", metadata_entries)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org and not isinstance(org, str):
            raise TypeError("Expected argument 'org' to be a str")
        pulumi.set(__self__, "org", org)
        if shared and not isinstance(shared, bool):
            raise TypeError("Expected argument 'shared' to be a bool")
        pulumi.set(__self__, "shared", shared)
        if vdc and not isinstance(vdc, str):
            raise TypeError("Expected argument 'vdc' to be a str")
        pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="externalNetwork")
    def external_network(self) -> str:
        return pulumi.get(self, "external_network")

    @property
    @pulumi.getter(name="externalNetworkDns1")
    def external_network_dns1(self) -> str:
        return pulumi.get(self, "external_network_dns1")

    @property
    @pulumi.getter(name="externalNetworkDns2")
    def external_network_dns2(self) -> str:
        return pulumi.get(self, "external_network_dns2")

    @property
    @pulumi.getter(name="externalNetworkDnsSuffix")
    def external_network_dns_suffix(self) -> str:
        return pulumi.get(self, "external_network_dns_suffix")

    @property
    @pulumi.getter(name="externalNetworkGateway")
    def external_network_gateway(self) -> str:
        return pulumi.get(self, "external_network_gateway")

    @property
    @pulumi.getter(name="externalNetworkNetmask")
    def external_network_netmask(self) -> str:
        return pulumi.get(self, "external_network_netmask")

    @property
    @pulumi.getter
    def filter(self) -> Optional['outputs.GetNetworkDirectFilterResult']:
        return pulumi.get(self, "filter")

    @property
    @pulumi.getter
    def href(self) -> str:
        return pulumi.get(self, "href")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def metadata(self) -> Mapping[str, Any]:
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter(name="metadataEntries")
    def metadata_entries(self) -> Sequence['outputs.GetNetworkDirectMetadataEntryResult']:
        return pulumi.get(self, "metadata_entries")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def org(self) -> Optional[str]:
        return pulumi.get(self, "org")

    @property
    @pulumi.getter
    def shared(self) -> bool:
        return pulumi.get(self, "shared")

    @property
    @pulumi.getter
    def vdc(self) -> Optional[str]:
        return pulumi.get(self, "vdc")


class AwaitableGetNetworkDirectResult(GetNetworkDirectResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNetworkDirectResult(
            description=self.description,
            external_network=self.external_network,
            external_network_dns1=self.external_network_dns1,
            external_network_dns2=self.external_network_dns2,
            external_network_dns_suffix=self.external_network_dns_suffix,
            external_network_gateway=self.external_network_gateway,
            external_network_netmask=self.external_network_netmask,
            filter=self.filter,
            href=self.href,
            id=self.id,
            metadata=self.metadata,
            metadata_entries=self.metadata_entries,
            name=self.name,
            org=self.org,
            shared=self.shared,
            vdc=self.vdc)


def get_network_direct(filter: Optional[pulumi.InputType['GetNetworkDirectFilterArgs']] = None,
                       name: Optional[str] = None,
                       org: Optional[str] = None,
                       vdc: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNetworkDirectResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['filter'] = filter
    __args__['name'] = name
    __args__['org'] = org
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getNetworkDirect:getNetworkDirect', __args__, opts=opts, typ=GetNetworkDirectResult).value

    return AwaitableGetNetworkDirectResult(
        description=__ret__.description,
        external_network=__ret__.external_network,
        external_network_dns1=__ret__.external_network_dns1,
        external_network_dns2=__ret__.external_network_dns2,
        external_network_dns_suffix=__ret__.external_network_dns_suffix,
        external_network_gateway=__ret__.external_network_gateway,
        external_network_netmask=__ret__.external_network_netmask,
        filter=__ret__.filter,
        href=__ret__.href,
        id=__ret__.id,
        metadata=__ret__.metadata,
        metadata_entries=__ret__.metadata_entries,
        name=__ret__.name,
        org=__ret__.org,
        shared=__ret__.shared,
        vdc=__ret__.vdc)


@_utilities.lift_output_func(get_network_direct)
def get_network_direct_output(filter: Optional[pulumi.Input[Optional[pulumi.InputType['GetNetworkDirectFilterArgs']]]] = None,
                              name: Optional[pulumi.Input[Optional[str]]] = None,
                              org: Optional[pulumi.Input[Optional[str]]] = None,
                              vdc: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNetworkDirectResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...