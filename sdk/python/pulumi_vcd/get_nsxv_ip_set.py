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
    'GetNsxvIpSetResult',
    'AwaitableGetNsxvIpSetResult',
    'get_nsxv_ip_set',
    'get_nsxv_ip_set_output',
]

@pulumi.output_type
class GetNsxvIpSetResult:
    """
    A collection of values returned by getNsxvIpSet.
    """
    def __init__(__self__, description=None, id=None, ip_addresses=None, is_inheritance_allowed=None, name=None, org=None, vdc=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_addresses and not isinstance(ip_addresses, list):
            raise TypeError("Expected argument 'ip_addresses' to be a list")
        pulumi.set(__self__, "ip_addresses", ip_addresses)
        if is_inheritance_allowed and not isinstance(is_inheritance_allowed, bool):
            raise TypeError("Expected argument 'is_inheritance_allowed' to be a bool")
        pulumi.set(__self__, "is_inheritance_allowed", is_inheritance_allowed)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org and not isinstance(org, str):
            raise TypeError("Expected argument 'org' to be a str")
        pulumi.set(__self__, "org", org)
        if vdc and not isinstance(vdc, str):
            raise TypeError("Expected argument 'vdc' to be a str")
        pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

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
    @pulumi.getter(name="isInheritanceAllowed")
    def is_inheritance_allowed(self) -> bool:
        return pulumi.get(self, "is_inheritance_allowed")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def org(self) -> Optional[str]:
        return pulumi.get(self, "org")

    @property
    @pulumi.getter
    def vdc(self) -> Optional[str]:
        return pulumi.get(self, "vdc")


class AwaitableGetNsxvIpSetResult(GetNsxvIpSetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNsxvIpSetResult(
            description=self.description,
            id=self.id,
            ip_addresses=self.ip_addresses,
            is_inheritance_allowed=self.is_inheritance_allowed,
            name=self.name,
            org=self.org,
            vdc=self.vdc)


def get_nsxv_ip_set(name: Optional[str] = None,
                    org: Optional[str] = None,
                    vdc: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNsxvIpSetResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['org'] = org
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getNsxvIpSet:getNsxvIpSet', __args__, opts=opts, typ=GetNsxvIpSetResult).value

    return AwaitableGetNsxvIpSetResult(
        description=__ret__.description,
        id=__ret__.id,
        ip_addresses=__ret__.ip_addresses,
        is_inheritance_allowed=__ret__.is_inheritance_allowed,
        name=__ret__.name,
        org=__ret__.org,
        vdc=__ret__.vdc)


@_utilities.lift_output_func(get_nsxv_ip_set)
def get_nsxv_ip_set_output(name: Optional[pulumi.Input[str]] = None,
                           org: Optional[pulumi.Input[Optional[str]]] = None,
                           vdc: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNsxvIpSetResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...