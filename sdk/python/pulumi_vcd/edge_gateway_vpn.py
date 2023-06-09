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

__all__ = ['EdgeGatewayVpnArgs', 'EdgeGatewayVpn']

@pulumi.input_type
class EdgeGatewayVpnArgs:
    def __init__(__self__, *,
                 edge_gateway: pulumi.Input[str],
                 encryption_protocol: pulumi.Input[str],
                 local_id: pulumi.Input[str],
                 local_ip_address: pulumi.Input[str],
                 mtu: pulumi.Input[int],
                 peer_id: pulumi.Input[str],
                 peer_ip_address: pulumi.Input[str],
                 shared_secret: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 local_subnets: Optional[pulumi.Input[Sequence[pulumi.Input['EdgeGatewayVpnLocalSubnetArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 peer_subnets: Optional[pulumi.Input[Sequence[pulumi.Input['EdgeGatewayVpnPeerSubnetArgs']]]] = None,
                 vdc: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EdgeGatewayVpn resource.
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        pulumi.set(__self__, "edge_gateway", edge_gateway)
        pulumi.set(__self__, "encryption_protocol", encryption_protocol)
        pulumi.set(__self__, "local_id", local_id)
        pulumi.set(__self__, "local_ip_address", local_ip_address)
        pulumi.set(__self__, "mtu", mtu)
        pulumi.set(__self__, "peer_id", peer_id)
        pulumi.set(__self__, "peer_ip_address", peer_ip_address)
        pulumi.set(__self__, "shared_secret", shared_secret)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if local_subnets is not None:
            pulumi.set(__self__, "local_subnets", local_subnets)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if peer_subnets is not None:
            pulumi.set(__self__, "peer_subnets", peer_subnets)
        if vdc is not None:
            pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter(name="edgeGateway")
    def edge_gateway(self) -> pulumi.Input[str]:
        return pulumi.get(self, "edge_gateway")

    @edge_gateway.setter
    def edge_gateway(self, value: pulumi.Input[str]):
        pulumi.set(self, "edge_gateway", value)

    @property
    @pulumi.getter(name="encryptionProtocol")
    def encryption_protocol(self) -> pulumi.Input[str]:
        return pulumi.get(self, "encryption_protocol")

    @encryption_protocol.setter
    def encryption_protocol(self, value: pulumi.Input[str]):
        pulumi.set(self, "encryption_protocol", value)

    @property
    @pulumi.getter(name="localId")
    def local_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "local_id")

    @local_id.setter
    def local_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "local_id", value)

    @property
    @pulumi.getter(name="localIpAddress")
    def local_ip_address(self) -> pulumi.Input[str]:
        return pulumi.get(self, "local_ip_address")

    @local_ip_address.setter
    def local_ip_address(self, value: pulumi.Input[str]):
        pulumi.set(self, "local_ip_address", value)

    @property
    @pulumi.getter
    def mtu(self) -> pulumi.Input[int]:
        return pulumi.get(self, "mtu")

    @mtu.setter
    def mtu(self, value: pulumi.Input[int]):
        pulumi.set(self, "mtu", value)

    @property
    @pulumi.getter(name="peerId")
    def peer_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "peer_id")

    @peer_id.setter
    def peer_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "peer_id", value)

    @property
    @pulumi.getter(name="peerIpAddress")
    def peer_ip_address(self) -> pulumi.Input[str]:
        return pulumi.get(self, "peer_ip_address")

    @peer_ip_address.setter
    def peer_ip_address(self, value: pulumi.Input[str]):
        pulumi.set(self, "peer_ip_address", value)

    @property
    @pulumi.getter(name="sharedSecret")
    def shared_secret(self) -> pulumi.Input[str]:
        return pulumi.get(self, "shared_secret")

    @shared_secret.setter
    def shared_secret(self, value: pulumi.Input[str]):
        pulumi.set(self, "shared_secret", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="localSubnets")
    def local_subnets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EdgeGatewayVpnLocalSubnetArgs']]]]:
        return pulumi.get(self, "local_subnets")

    @local_subnets.setter
    def local_subnets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EdgeGatewayVpnLocalSubnetArgs']]]]):
        pulumi.set(self, "local_subnets", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
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
    @pulumi.getter(name="peerSubnets")
    def peer_subnets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EdgeGatewayVpnPeerSubnetArgs']]]]:
        return pulumi.get(self, "peer_subnets")

    @peer_subnets.setter
    def peer_subnets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EdgeGatewayVpnPeerSubnetArgs']]]]):
        pulumi.set(self, "peer_subnets", value)

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
class _EdgeGatewayVpnState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 edge_gateway: Optional[pulumi.Input[str]] = None,
                 encryption_protocol: Optional[pulumi.Input[str]] = None,
                 local_id: Optional[pulumi.Input[str]] = None,
                 local_ip_address: Optional[pulumi.Input[str]] = None,
                 local_subnets: Optional[pulumi.Input[Sequence[pulumi.Input['EdgeGatewayVpnLocalSubnetArgs']]]] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 peer_id: Optional[pulumi.Input[str]] = None,
                 peer_ip_address: Optional[pulumi.Input[str]] = None,
                 peer_subnets: Optional[pulumi.Input[Sequence[pulumi.Input['EdgeGatewayVpnPeerSubnetArgs']]]] = None,
                 shared_secret: Optional[pulumi.Input[str]] = None,
                 vdc: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EdgeGatewayVpn resources.
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if edge_gateway is not None:
            pulumi.set(__self__, "edge_gateway", edge_gateway)
        if encryption_protocol is not None:
            pulumi.set(__self__, "encryption_protocol", encryption_protocol)
        if local_id is not None:
            pulumi.set(__self__, "local_id", local_id)
        if local_ip_address is not None:
            pulumi.set(__self__, "local_ip_address", local_ip_address)
        if local_subnets is not None:
            pulumi.set(__self__, "local_subnets", local_subnets)
        if mtu is not None:
            pulumi.set(__self__, "mtu", mtu)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if peer_id is not None:
            pulumi.set(__self__, "peer_id", peer_id)
        if peer_ip_address is not None:
            pulumi.set(__self__, "peer_ip_address", peer_ip_address)
        if peer_subnets is not None:
            pulumi.set(__self__, "peer_subnets", peer_subnets)
        if shared_secret is not None:
            pulumi.set(__self__, "shared_secret", shared_secret)
        if vdc is not None:
            pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="edgeGateway")
    def edge_gateway(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "edge_gateway")

    @edge_gateway.setter
    def edge_gateway(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "edge_gateway", value)

    @property
    @pulumi.getter(name="encryptionProtocol")
    def encryption_protocol(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "encryption_protocol")

    @encryption_protocol.setter
    def encryption_protocol(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encryption_protocol", value)

    @property
    @pulumi.getter(name="localId")
    def local_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "local_id")

    @local_id.setter
    def local_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_id", value)

    @property
    @pulumi.getter(name="localIpAddress")
    def local_ip_address(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "local_ip_address")

    @local_ip_address.setter
    def local_ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "local_ip_address", value)

    @property
    @pulumi.getter(name="localSubnets")
    def local_subnets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EdgeGatewayVpnLocalSubnetArgs']]]]:
        return pulumi.get(self, "local_subnets")

    @local_subnets.setter
    def local_subnets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EdgeGatewayVpnLocalSubnetArgs']]]]):
        pulumi.set(self, "local_subnets", value)

    @property
    @pulumi.getter
    def mtu(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "mtu")

    @mtu.setter
    def mtu(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "mtu", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
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
    @pulumi.getter(name="peerId")
    def peer_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "peer_id")

    @peer_id.setter
    def peer_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_id", value)

    @property
    @pulumi.getter(name="peerIpAddress")
    def peer_ip_address(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "peer_ip_address")

    @peer_ip_address.setter
    def peer_ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "peer_ip_address", value)

    @property
    @pulumi.getter(name="peerSubnets")
    def peer_subnets(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['EdgeGatewayVpnPeerSubnetArgs']]]]:
        return pulumi.get(self, "peer_subnets")

    @peer_subnets.setter
    def peer_subnets(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['EdgeGatewayVpnPeerSubnetArgs']]]]):
        pulumi.set(self, "peer_subnets", value)

    @property
    @pulumi.getter(name="sharedSecret")
    def shared_secret(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "shared_secret")

    @shared_secret.setter
    def shared_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "shared_secret", value)

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


class EdgeGatewayVpn(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 edge_gateway: Optional[pulumi.Input[str]] = None,
                 encryption_protocol: Optional[pulumi.Input[str]] = None,
                 local_id: Optional[pulumi.Input[str]] = None,
                 local_ip_address: Optional[pulumi.Input[str]] = None,
                 local_subnets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EdgeGatewayVpnLocalSubnetArgs']]]]] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 peer_id: Optional[pulumi.Input[str]] = None,
                 peer_ip_address: Optional[pulumi.Input[str]] = None,
                 peer_subnets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EdgeGatewayVpnPeerSubnetArgs']]]]] = None,
                 shared_secret: Optional[pulumi.Input[str]] = None,
                 vdc: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a EdgeGatewayVpn resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EdgeGatewayVpnArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a EdgeGatewayVpn resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param EdgeGatewayVpnArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EdgeGatewayVpnArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 edge_gateway: Optional[pulumi.Input[str]] = None,
                 encryption_protocol: Optional[pulumi.Input[str]] = None,
                 local_id: Optional[pulumi.Input[str]] = None,
                 local_ip_address: Optional[pulumi.Input[str]] = None,
                 local_subnets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EdgeGatewayVpnLocalSubnetArgs']]]]] = None,
                 mtu: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 peer_id: Optional[pulumi.Input[str]] = None,
                 peer_ip_address: Optional[pulumi.Input[str]] = None,
                 peer_subnets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EdgeGatewayVpnPeerSubnetArgs']]]]] = None,
                 shared_secret: Optional[pulumi.Input[str]] = None,
                 vdc: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EdgeGatewayVpnArgs.__new__(EdgeGatewayVpnArgs)

            __props__.__dict__["description"] = description
            if edge_gateway is None and not opts.urn:
                raise TypeError("Missing required property 'edge_gateway'")
            __props__.__dict__["edge_gateway"] = edge_gateway
            if encryption_protocol is None and not opts.urn:
                raise TypeError("Missing required property 'encryption_protocol'")
            __props__.__dict__["encryption_protocol"] = encryption_protocol
            if local_id is None and not opts.urn:
                raise TypeError("Missing required property 'local_id'")
            __props__.__dict__["local_id"] = local_id
            if local_ip_address is None and not opts.urn:
                raise TypeError("Missing required property 'local_ip_address'")
            __props__.__dict__["local_ip_address"] = local_ip_address
            __props__.__dict__["local_subnets"] = local_subnets
            if mtu is None and not opts.urn:
                raise TypeError("Missing required property 'mtu'")
            __props__.__dict__["mtu"] = mtu
            __props__.__dict__["name"] = name
            __props__.__dict__["org"] = org
            if peer_id is None and not opts.urn:
                raise TypeError("Missing required property 'peer_id'")
            __props__.__dict__["peer_id"] = peer_id
            if peer_ip_address is None and not opts.urn:
                raise TypeError("Missing required property 'peer_ip_address'")
            __props__.__dict__["peer_ip_address"] = peer_ip_address
            __props__.__dict__["peer_subnets"] = peer_subnets
            if shared_secret is None and not opts.urn:
                raise TypeError("Missing required property 'shared_secret'")
            __props__.__dict__["shared_secret"] = shared_secret
            __props__.__dict__["vdc"] = vdc
        super(EdgeGatewayVpn, __self__).__init__(
            'vcd:index/edgeGatewayVpn:EdgeGatewayVpn',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            edge_gateway: Optional[pulumi.Input[str]] = None,
            encryption_protocol: Optional[pulumi.Input[str]] = None,
            local_id: Optional[pulumi.Input[str]] = None,
            local_ip_address: Optional[pulumi.Input[str]] = None,
            local_subnets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EdgeGatewayVpnLocalSubnetArgs']]]]] = None,
            mtu: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org: Optional[pulumi.Input[str]] = None,
            peer_id: Optional[pulumi.Input[str]] = None,
            peer_ip_address: Optional[pulumi.Input[str]] = None,
            peer_subnets: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['EdgeGatewayVpnPeerSubnetArgs']]]]] = None,
            shared_secret: Optional[pulumi.Input[str]] = None,
            vdc: Optional[pulumi.Input[str]] = None) -> 'EdgeGatewayVpn':
        """
        Get an existing EdgeGatewayVpn resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EdgeGatewayVpnState.__new__(_EdgeGatewayVpnState)

        __props__.__dict__["description"] = description
        __props__.__dict__["edge_gateway"] = edge_gateway
        __props__.__dict__["encryption_protocol"] = encryption_protocol
        __props__.__dict__["local_id"] = local_id
        __props__.__dict__["local_ip_address"] = local_ip_address
        __props__.__dict__["local_subnets"] = local_subnets
        __props__.__dict__["mtu"] = mtu
        __props__.__dict__["name"] = name
        __props__.__dict__["org"] = org
        __props__.__dict__["peer_id"] = peer_id
        __props__.__dict__["peer_ip_address"] = peer_ip_address
        __props__.__dict__["peer_subnets"] = peer_subnets
        __props__.__dict__["shared_secret"] = shared_secret
        __props__.__dict__["vdc"] = vdc
        return EdgeGatewayVpn(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="edgeGateway")
    def edge_gateway(self) -> pulumi.Output[str]:
        return pulumi.get(self, "edge_gateway")

    @property
    @pulumi.getter(name="encryptionProtocol")
    def encryption_protocol(self) -> pulumi.Output[str]:
        return pulumi.get(self, "encryption_protocol")

    @property
    @pulumi.getter(name="localId")
    def local_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "local_id")

    @property
    @pulumi.getter(name="localIpAddress")
    def local_ip_address(self) -> pulumi.Output[str]:
        return pulumi.get(self, "local_ip_address")

    @property
    @pulumi.getter(name="localSubnets")
    def local_subnets(self) -> pulumi.Output[Optional[Sequence['outputs.EdgeGatewayVpnLocalSubnet']]]:
        return pulumi.get(self, "local_subnets")

    @property
    @pulumi.getter
    def mtu(self) -> pulumi.Output[int]:
        return pulumi.get(self, "mtu")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
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
    @pulumi.getter(name="peerId")
    def peer_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "peer_id")

    @property
    @pulumi.getter(name="peerIpAddress")
    def peer_ip_address(self) -> pulumi.Output[str]:
        return pulumi.get(self, "peer_ip_address")

    @property
    @pulumi.getter(name="peerSubnets")
    def peer_subnets(self) -> pulumi.Output[Optional[Sequence['outputs.EdgeGatewayVpnPeerSubnet']]]:
        return pulumi.get(self, "peer_subnets")

    @property
    @pulumi.getter(name="sharedSecret")
    def shared_secret(self) -> pulumi.Output[str]:
        return pulumi.get(self, "shared_secret")

    @property
    @pulumi.getter
    def vdc(self) -> pulumi.Output[Optional[str]]:
        """
        The name of VDC to use, optional if defined at provider level
        """
        return pulumi.get(self, "vdc")

