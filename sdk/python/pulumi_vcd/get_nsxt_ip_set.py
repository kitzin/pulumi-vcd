# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetNsxtIpSetResult',
    'AwaitableGetNsxtIpSetResult',
    'get_nsxt_ip_set',
    'get_nsxt_ip_set_output',
]

@pulumi.output_type
class GetNsxtIpSetResult:
    """
    A collection of values returned by getNsxtIpSet.
    """
    def __init__(__self__, description=None, edge_gateway_id=None, id=None, ip_addresses=None, name=None, org=None, owner_id=None, vdc=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if edge_gateway_id and not isinstance(edge_gateway_id, str):
            raise TypeError("Expected argument 'edge_gateway_id' to be a str")
        pulumi.set(__self__, "edge_gateway_id", edge_gateway_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_addresses and not isinstance(ip_addresses, list):
            raise TypeError("Expected argument 'ip_addresses' to be a list")
        pulumi.set(__self__, "ip_addresses", ip_addresses)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org and not isinstance(org, str):
            raise TypeError("Expected argument 'org' to be a str")
        pulumi.set(__self__, "org", org)
        if owner_id and not isinstance(owner_id, str):
            raise TypeError("Expected argument 'owner_id' to be a str")
        pulumi.set(__self__, "owner_id", owner_id)
        if vdc and not isinstance(vdc, str):
            raise TypeError("Expected argument 'vdc' to be a str")
        if vdc is not None:
            warnings.warn("""Deprecated in favor of `edge_gateway_id`. IP Set will inherit VDC from parent Edge Gateway.""", DeprecationWarning)
            pulumi.log.warn("""vdc is deprecated: Deprecated in favor of `edge_gateway_id`. IP Set will inherit VDC from parent Edge Gateway.""")

        pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="edgeGatewayId")
    def edge_gateway_id(self) -> str:
        return pulumi.get(self, "edge_gateway_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Sequence[str]:
        return pulumi.get(self, "ip_addresses")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def org(self) -> Optional[str]:
        return pulumi.get(self, "org")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> str:
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter
    def vdc(self) -> Optional[str]:
        return pulumi.get(self, "vdc")


class AwaitableGetNsxtIpSetResult(GetNsxtIpSetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNsxtIpSetResult(
            description=self.description,
            edge_gateway_id=self.edge_gateway_id,
            id=self.id,
            ip_addresses=self.ip_addresses,
            name=self.name,
            org=self.org,
            owner_id=self.owner_id,
            vdc=self.vdc)


def get_nsxt_ip_set(edge_gateway_id: Optional[str] = None,
                    name: Optional[str] = None,
                    org: Optional[str] = None,
                    vdc: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNsxtIpSetResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['edgeGatewayId'] = edge_gateway_id
    __args__['name'] = name
    __args__['org'] = org
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getNsxtIpSet:getNsxtIpSet', __args__, opts=opts, typ=GetNsxtIpSetResult).value

    return AwaitableGetNsxtIpSetResult(
        description=__ret__.description,
        edge_gateway_id=__ret__.edge_gateway_id,
        id=__ret__.id,
        ip_addresses=__ret__.ip_addresses,
        name=__ret__.name,
        org=__ret__.org,
        owner_id=__ret__.owner_id,
        vdc=__ret__.vdc)


@_utilities.lift_output_func(get_nsxt_ip_set)
def get_nsxt_ip_set_output(edge_gateway_id: Optional[pulumi.Input[str]] = None,
                           name: Optional[pulumi.Input[str]] = None,
                           org: Optional[pulumi.Input[Optional[str]]] = None,
                           vdc: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNsxtIpSetResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
