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
    'GetVappOrgNetworkResult',
    'AwaitableGetVappOrgNetworkResult',
    'get_vapp_org_network',
    'get_vapp_org_network_output',
]

@pulumi.output_type
class GetVappOrgNetworkResult:
    """
    A collection of values returned by getVappOrgNetwork.
    """
    def __init__(__self__, id=None, is_fenced=None, org=None, org_network_name=None, retain_ip_mac_enabled=None, vapp_name=None, vdc=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_fenced and not isinstance(is_fenced, bool):
            raise TypeError("Expected argument 'is_fenced' to be a bool")
        pulumi.set(__self__, "is_fenced", is_fenced)
        if org and not isinstance(org, str):
            raise TypeError("Expected argument 'org' to be a str")
        pulumi.set(__self__, "org", org)
        if org_network_name and not isinstance(org_network_name, str):
            raise TypeError("Expected argument 'org_network_name' to be a str")
        pulumi.set(__self__, "org_network_name", org_network_name)
        if retain_ip_mac_enabled and not isinstance(retain_ip_mac_enabled, bool):
            raise TypeError("Expected argument 'retain_ip_mac_enabled' to be a bool")
        pulumi.set(__self__, "retain_ip_mac_enabled", retain_ip_mac_enabled)
        if vapp_name and not isinstance(vapp_name, str):
            raise TypeError("Expected argument 'vapp_name' to be a str")
        pulumi.set(__self__, "vapp_name", vapp_name)
        if vdc and not isinstance(vdc, str):
            raise TypeError("Expected argument 'vdc' to be a str")
        pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isFenced")
    def is_fenced(self) -> bool:
        return pulumi.get(self, "is_fenced")

    @property
    @pulumi.getter
    def org(self) -> Optional[str]:
        return pulumi.get(self, "org")

    @property
    @pulumi.getter(name="orgNetworkName")
    def org_network_name(self) -> str:
        return pulumi.get(self, "org_network_name")

    @property
    @pulumi.getter(name="retainIpMacEnabled")
    def retain_ip_mac_enabled(self) -> bool:
        return pulumi.get(self, "retain_ip_mac_enabled")

    @property
    @pulumi.getter(name="vappName")
    def vapp_name(self) -> str:
        return pulumi.get(self, "vapp_name")

    @property
    @pulumi.getter
    def vdc(self) -> Optional[str]:
        return pulumi.get(self, "vdc")


class AwaitableGetVappOrgNetworkResult(GetVappOrgNetworkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVappOrgNetworkResult(
            id=self.id,
            is_fenced=self.is_fenced,
            org=self.org,
            org_network_name=self.org_network_name,
            retain_ip_mac_enabled=self.retain_ip_mac_enabled,
            vapp_name=self.vapp_name,
            vdc=self.vdc)


def get_vapp_org_network(org: Optional[str] = None,
                         org_network_name: Optional[str] = None,
                         vapp_name: Optional[str] = None,
                         vdc: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVappOrgNetworkResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['org'] = org
    __args__['orgNetworkName'] = org_network_name
    __args__['vappName'] = vapp_name
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getVappOrgNetwork:getVappOrgNetwork', __args__, opts=opts, typ=GetVappOrgNetworkResult).value

    return AwaitableGetVappOrgNetworkResult(
        id=__ret__.id,
        is_fenced=__ret__.is_fenced,
        org=__ret__.org,
        org_network_name=__ret__.org_network_name,
        retain_ip_mac_enabled=__ret__.retain_ip_mac_enabled,
        vapp_name=__ret__.vapp_name,
        vdc=__ret__.vdc)


@_utilities.lift_output_func(get_vapp_org_network)
def get_vapp_org_network_output(org: Optional[pulumi.Input[Optional[str]]] = None,
                                org_network_name: Optional[pulumi.Input[str]] = None,
                                vapp_name: Optional[pulumi.Input[str]] = None,
                                vdc: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVappOrgNetworkResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...