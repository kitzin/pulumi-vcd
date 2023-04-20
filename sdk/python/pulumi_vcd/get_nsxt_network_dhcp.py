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

__all__ = [
    'GetNsxtNetworkDhcpResult',
    'AwaitableGetNsxtNetworkDhcpResult',
    'get_nsxt_network_dhcp',
    'get_nsxt_network_dhcp_output',
]

@pulumi.output_type
class GetNsxtNetworkDhcpResult:
    """
    A collection of values returned by getNsxtNetworkDhcp.
    """
    def __init__(__self__, dns_servers=None, id=None, lease_time=None, listener_ip_address=None, mode=None, org=None, org_network_id=None, pools=None, vdc=None):
        if dns_servers and not isinstance(dns_servers, list):
            raise TypeError("Expected argument 'dns_servers' to be a list")
        pulumi.set(__self__, "dns_servers", dns_servers)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lease_time and not isinstance(lease_time, int):
            raise TypeError("Expected argument 'lease_time' to be a int")
        pulumi.set(__self__, "lease_time", lease_time)
        if listener_ip_address and not isinstance(listener_ip_address, str):
            raise TypeError("Expected argument 'listener_ip_address' to be a str")
        pulumi.set(__self__, "listener_ip_address", listener_ip_address)
        if mode and not isinstance(mode, str):
            raise TypeError("Expected argument 'mode' to be a str")
        pulumi.set(__self__, "mode", mode)
        if org and not isinstance(org, str):
            raise TypeError("Expected argument 'org' to be a str")
        pulumi.set(__self__, "org", org)
        if org_network_id and not isinstance(org_network_id, str):
            raise TypeError("Expected argument 'org_network_id' to be a str")
        pulumi.set(__self__, "org_network_id", org_network_id)
        if pools and not isinstance(pools, list):
            raise TypeError("Expected argument 'pools' to be a list")
        pulumi.set(__self__, "pools", pools)
        if vdc and not isinstance(vdc, str):
            raise TypeError("Expected argument 'vdc' to be a str")
        if vdc is not None:
            warnings.warn("""Org network will be looked up based on 'org_network_id' field""", DeprecationWarning)
            pulumi.log.warn("""vdc is deprecated: Org network will be looked up based on 'org_network_id' field""")

        pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter(name="dnsServers")
    def dns_servers(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "dns_servers")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="leaseTime")
    def lease_time(self) -> int:
        return pulumi.get(self, "lease_time")

    @property
    @pulumi.getter(name="listenerIpAddress")
    def listener_ip_address(self) -> str:
        return pulumi.get(self, "listener_ip_address")

    @property
    @pulumi.getter
    def mode(self) -> str:
        return pulumi.get(self, "mode")

    @property
    @pulumi.getter
    def org(self) -> Optional[str]:
        return pulumi.get(self, "org")

    @property
    @pulumi.getter(name="orgNetworkId")
    def org_network_id(self) -> str:
        return pulumi.get(self, "org_network_id")

    @property
    @pulumi.getter
    def pools(self) -> Sequence['outputs.GetNsxtNetworkDhcpPoolResult']:
        return pulumi.get(self, "pools")

    @property
    @pulumi.getter
    def vdc(self) -> Optional[str]:
        return pulumi.get(self, "vdc")


class AwaitableGetNsxtNetworkDhcpResult(GetNsxtNetworkDhcpResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNsxtNetworkDhcpResult(
            dns_servers=self.dns_servers,
            id=self.id,
            lease_time=self.lease_time,
            listener_ip_address=self.listener_ip_address,
            mode=self.mode,
            org=self.org,
            org_network_id=self.org_network_id,
            pools=self.pools,
            vdc=self.vdc)


def get_nsxt_network_dhcp(dns_servers: Optional[Sequence[str]] = None,
                          org: Optional[str] = None,
                          org_network_id: Optional[str] = None,
                          vdc: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNsxtNetworkDhcpResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['dnsServers'] = dns_servers
    __args__['org'] = org
    __args__['orgNetworkId'] = org_network_id
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getNsxtNetworkDhcp:getNsxtNetworkDhcp', __args__, opts=opts, typ=GetNsxtNetworkDhcpResult).value

    return AwaitableGetNsxtNetworkDhcpResult(
        dns_servers=__ret__.dns_servers,
        id=__ret__.id,
        lease_time=__ret__.lease_time,
        listener_ip_address=__ret__.listener_ip_address,
        mode=__ret__.mode,
        org=__ret__.org,
        org_network_id=__ret__.org_network_id,
        pools=__ret__.pools,
        vdc=__ret__.vdc)


@_utilities.lift_output_func(get_nsxt_network_dhcp)
def get_nsxt_network_dhcp_output(dns_servers: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                 org: Optional[pulumi.Input[Optional[str]]] = None,
                                 org_network_id: Optional[pulumi.Input[str]] = None,
                                 vdc: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNsxtNetworkDhcpResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
