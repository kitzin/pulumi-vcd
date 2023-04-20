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
    'GetLbVirtualServerResult',
    'AwaitableGetLbVirtualServerResult',
    'get_lb_virtual_server',
    'get_lb_virtual_server_output',
]

@pulumi.output_type
class GetLbVirtualServerResult:
    """
    A collection of values returned by getLbVirtualServer.
    """
    def __init__(__self__, app_profile_id=None, app_rule_ids=None, connection_limit=None, connection_rate_limit=None, description=None, edge_gateway=None, enable_acceleration=None, enabled=None, id=None, ip_address=None, name=None, org=None, port=None, protocol=None, server_pool_id=None, vdc=None):
        if app_profile_id and not isinstance(app_profile_id, str):
            raise TypeError("Expected argument 'app_profile_id' to be a str")
        pulumi.set(__self__, "app_profile_id", app_profile_id)
        if app_rule_ids and not isinstance(app_rule_ids, list):
            raise TypeError("Expected argument 'app_rule_ids' to be a list")
        pulumi.set(__self__, "app_rule_ids", app_rule_ids)
        if connection_limit and not isinstance(connection_limit, int):
            raise TypeError("Expected argument 'connection_limit' to be a int")
        pulumi.set(__self__, "connection_limit", connection_limit)
        if connection_rate_limit and not isinstance(connection_rate_limit, int):
            raise TypeError("Expected argument 'connection_rate_limit' to be a int")
        pulumi.set(__self__, "connection_rate_limit", connection_rate_limit)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if edge_gateway and not isinstance(edge_gateway, str):
            raise TypeError("Expected argument 'edge_gateway' to be a str")
        pulumi.set(__self__, "edge_gateway", edge_gateway)
        if enable_acceleration and not isinstance(enable_acceleration, bool):
            raise TypeError("Expected argument 'enable_acceleration' to be a bool")
        pulumi.set(__self__, "enable_acceleration", enable_acceleration)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_address and not isinstance(ip_address, str):
            raise TypeError("Expected argument 'ip_address' to be a str")
        pulumi.set(__self__, "ip_address", ip_address)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org and not isinstance(org, str):
            raise TypeError("Expected argument 'org' to be a str")
        pulumi.set(__self__, "org", org)
        if port and not isinstance(port, int):
            raise TypeError("Expected argument 'port' to be a int")
        pulumi.set(__self__, "port", port)
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        pulumi.set(__self__, "protocol", protocol)
        if server_pool_id and not isinstance(server_pool_id, str):
            raise TypeError("Expected argument 'server_pool_id' to be a str")
        pulumi.set(__self__, "server_pool_id", server_pool_id)
        if vdc and not isinstance(vdc, str):
            raise TypeError("Expected argument 'vdc' to be a str")
        pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter(name="appProfileId")
    def app_profile_id(self) -> str:
        return pulumi.get(self, "app_profile_id")

    @property
    @pulumi.getter(name="appRuleIds")
    def app_rule_ids(self) -> Sequence[str]:
        return pulumi.get(self, "app_rule_ids")

    @property
    @pulumi.getter(name="connectionLimit")
    def connection_limit(self) -> int:
        return pulumi.get(self, "connection_limit")

    @property
    @pulumi.getter(name="connectionRateLimit")
    def connection_rate_limit(self) -> int:
        return pulumi.get(self, "connection_rate_limit")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="edgeGateway")
    def edge_gateway(self) -> str:
        return pulumi.get(self, "edge_gateway")

    @property
    @pulumi.getter(name="enableAcceleration")
    def enable_acceleration(self) -> bool:
        return pulumi.get(self, "enable_acceleration")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> str:
        return pulumi.get(self, "ip_address")

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
    def port(self) -> int:
        return pulumi.get(self, "port")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="serverPoolId")
    def server_pool_id(self) -> str:
        return pulumi.get(self, "server_pool_id")

    @property
    @pulumi.getter
    def vdc(self) -> Optional[str]:
        return pulumi.get(self, "vdc")


class AwaitableGetLbVirtualServerResult(GetLbVirtualServerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLbVirtualServerResult(
            app_profile_id=self.app_profile_id,
            app_rule_ids=self.app_rule_ids,
            connection_limit=self.connection_limit,
            connection_rate_limit=self.connection_rate_limit,
            description=self.description,
            edge_gateway=self.edge_gateway,
            enable_acceleration=self.enable_acceleration,
            enabled=self.enabled,
            id=self.id,
            ip_address=self.ip_address,
            name=self.name,
            org=self.org,
            port=self.port,
            protocol=self.protocol,
            server_pool_id=self.server_pool_id,
            vdc=self.vdc)


def get_lb_virtual_server(edge_gateway: Optional[str] = None,
                          name: Optional[str] = None,
                          org: Optional[str] = None,
                          vdc: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLbVirtualServerResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['edgeGateway'] = edge_gateway
    __args__['name'] = name
    __args__['org'] = org
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getLbVirtualServer:getLbVirtualServer', __args__, opts=opts, typ=GetLbVirtualServerResult).value

    return AwaitableGetLbVirtualServerResult(
        app_profile_id=__ret__.app_profile_id,
        app_rule_ids=__ret__.app_rule_ids,
        connection_limit=__ret__.connection_limit,
        connection_rate_limit=__ret__.connection_rate_limit,
        description=__ret__.description,
        edge_gateway=__ret__.edge_gateway,
        enable_acceleration=__ret__.enable_acceleration,
        enabled=__ret__.enabled,
        id=__ret__.id,
        ip_address=__ret__.ip_address,
        name=__ret__.name,
        org=__ret__.org,
        port=__ret__.port,
        protocol=__ret__.protocol,
        server_pool_id=__ret__.server_pool_id,
        vdc=__ret__.vdc)


@_utilities.lift_output_func(get_lb_virtual_server)
def get_lb_virtual_server_output(edge_gateway: Optional[pulumi.Input[str]] = None,
                                 name: Optional[pulumi.Input[str]] = None,
                                 org: Optional[pulumi.Input[Optional[str]]] = None,
                                 vdc: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLbVirtualServerResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...