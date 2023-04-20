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
    'GetLbAppProfileResult',
    'AwaitableGetLbAppProfileResult',
    'get_lb_app_profile',
    'get_lb_app_profile_output',
]

@pulumi.output_type
class GetLbAppProfileResult:
    """
    A collection of values returned by getLbAppProfile.
    """
    def __init__(__self__, cookie_mode=None, cookie_name=None, edge_gateway=None, enable_pool_side_ssl=None, enable_ssl_passthrough=None, expiration=None, http_redirect_url=None, id=None, insert_x_forwarded_http_header=None, name=None, org=None, persistence_mechanism=None, type=None, vdc=None):
        if cookie_mode and not isinstance(cookie_mode, str):
            raise TypeError("Expected argument 'cookie_mode' to be a str")
        pulumi.set(__self__, "cookie_mode", cookie_mode)
        if cookie_name and not isinstance(cookie_name, str):
            raise TypeError("Expected argument 'cookie_name' to be a str")
        pulumi.set(__self__, "cookie_name", cookie_name)
        if edge_gateway and not isinstance(edge_gateway, str):
            raise TypeError("Expected argument 'edge_gateway' to be a str")
        pulumi.set(__self__, "edge_gateway", edge_gateway)
        if enable_pool_side_ssl and not isinstance(enable_pool_side_ssl, bool):
            raise TypeError("Expected argument 'enable_pool_side_ssl' to be a bool")
        pulumi.set(__self__, "enable_pool_side_ssl", enable_pool_side_ssl)
        if enable_ssl_passthrough and not isinstance(enable_ssl_passthrough, bool):
            raise TypeError("Expected argument 'enable_ssl_passthrough' to be a bool")
        pulumi.set(__self__, "enable_ssl_passthrough", enable_ssl_passthrough)
        if expiration and not isinstance(expiration, int):
            raise TypeError("Expected argument 'expiration' to be a int")
        pulumi.set(__self__, "expiration", expiration)
        if http_redirect_url and not isinstance(http_redirect_url, str):
            raise TypeError("Expected argument 'http_redirect_url' to be a str")
        pulumi.set(__self__, "http_redirect_url", http_redirect_url)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if insert_x_forwarded_http_header and not isinstance(insert_x_forwarded_http_header, bool):
            raise TypeError("Expected argument 'insert_x_forwarded_http_header' to be a bool")
        pulumi.set(__self__, "insert_x_forwarded_http_header", insert_x_forwarded_http_header)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org and not isinstance(org, str):
            raise TypeError("Expected argument 'org' to be a str")
        pulumi.set(__self__, "org", org)
        if persistence_mechanism and not isinstance(persistence_mechanism, str):
            raise TypeError("Expected argument 'persistence_mechanism' to be a str")
        pulumi.set(__self__, "persistence_mechanism", persistence_mechanism)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if vdc and not isinstance(vdc, str):
            raise TypeError("Expected argument 'vdc' to be a str")
        pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter(name="cookieMode")
    def cookie_mode(self) -> str:
        return pulumi.get(self, "cookie_mode")

    @property
    @pulumi.getter(name="cookieName")
    def cookie_name(self) -> str:
        return pulumi.get(self, "cookie_name")

    @property
    @pulumi.getter(name="edgeGateway")
    def edge_gateway(self) -> str:
        return pulumi.get(self, "edge_gateway")

    @property
    @pulumi.getter(name="enablePoolSideSsl")
    def enable_pool_side_ssl(self) -> bool:
        return pulumi.get(self, "enable_pool_side_ssl")

    @property
    @pulumi.getter(name="enableSslPassthrough")
    def enable_ssl_passthrough(self) -> bool:
        return pulumi.get(self, "enable_ssl_passthrough")

    @property
    @pulumi.getter
    def expiration(self) -> int:
        return pulumi.get(self, "expiration")

    @property
    @pulumi.getter(name="httpRedirectUrl")
    def http_redirect_url(self) -> str:
        return pulumi.get(self, "http_redirect_url")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="insertXForwardedHttpHeader")
    def insert_x_forwarded_http_header(self) -> bool:
        return pulumi.get(self, "insert_x_forwarded_http_header")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def org(self) -> Optional[str]:
        return pulumi.get(self, "org")

    @property
    @pulumi.getter(name="persistenceMechanism")
    def persistence_mechanism(self) -> str:
        return pulumi.get(self, "persistence_mechanism")

    @property
    @pulumi.getter
    def type(self) -> str:
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def vdc(self) -> Optional[str]:
        return pulumi.get(self, "vdc")


class AwaitableGetLbAppProfileResult(GetLbAppProfileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetLbAppProfileResult(
            cookie_mode=self.cookie_mode,
            cookie_name=self.cookie_name,
            edge_gateway=self.edge_gateway,
            enable_pool_side_ssl=self.enable_pool_side_ssl,
            enable_ssl_passthrough=self.enable_ssl_passthrough,
            expiration=self.expiration,
            http_redirect_url=self.http_redirect_url,
            id=self.id,
            insert_x_forwarded_http_header=self.insert_x_forwarded_http_header,
            name=self.name,
            org=self.org,
            persistence_mechanism=self.persistence_mechanism,
            type=self.type,
            vdc=self.vdc)


def get_lb_app_profile(edge_gateway: Optional[str] = None,
                       name: Optional[str] = None,
                       org: Optional[str] = None,
                       vdc: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetLbAppProfileResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['edgeGateway'] = edge_gateway
    __args__['name'] = name
    __args__['org'] = org
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getLbAppProfile:getLbAppProfile', __args__, opts=opts, typ=GetLbAppProfileResult).value

    return AwaitableGetLbAppProfileResult(
        cookie_mode=__ret__.cookie_mode,
        cookie_name=__ret__.cookie_name,
        edge_gateway=__ret__.edge_gateway,
        enable_pool_side_ssl=__ret__.enable_pool_side_ssl,
        enable_ssl_passthrough=__ret__.enable_ssl_passthrough,
        expiration=__ret__.expiration,
        http_redirect_url=__ret__.http_redirect_url,
        id=__ret__.id,
        insert_x_forwarded_http_header=__ret__.insert_x_forwarded_http_header,
        name=__ret__.name,
        org=__ret__.org,
        persistence_mechanism=__ret__.persistence_mechanism,
        type=__ret__.type,
        vdc=__ret__.vdc)


@_utilities.lift_output_func(get_lb_app_profile)
def get_lb_app_profile_output(edge_gateway: Optional[pulumi.Input[str]] = None,
                              name: Optional[pulumi.Input[str]] = None,
                              org: Optional[pulumi.Input[Optional[str]]] = None,
                              vdc: Optional[pulumi.Input[Optional[str]]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetLbAppProfileResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
