# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RoleArgs', 'Role']

@pulumi.input_type
class RoleArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 rights: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Role resource.
        :param pulumi.Input[str] description: Role description
        :param pulumi.Input[str] name: Name of Role.
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[Sequence[pulumi.Input[str]]] rights: Set of rights assigned to this role
        """
        pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if rights is not None:
            pulumi.set(__self__, "rights", rights)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        """
        Role description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of Role.
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
    @pulumi.getter
    def rights(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Set of rights assigned to this role
        """
        return pulumi.get(self, "rights")

    @rights.setter
    def rights(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "rights", value)


@pulumi.input_type
class _RoleState:
    def __init__(__self__, *,
                 bundle_key: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 read_only: Optional[pulumi.Input[bool]] = None,
                 rights: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering Role resources.
        :param pulumi.Input[str] bundle_key: Key used for internationalization
        :param pulumi.Input[str] description: Role description
        :param pulumi.Input[str] name: Name of Role.
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[bool] read_only: Whether this role is read-only
        :param pulumi.Input[Sequence[pulumi.Input[str]]] rights: Set of rights assigned to this role
        """
        if bundle_key is not None:
            pulumi.set(__self__, "bundle_key", bundle_key)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if read_only is not None:
            pulumi.set(__self__, "read_only", read_only)
        if rights is not None:
            pulumi.set(__self__, "rights", rights)

    @property
    @pulumi.getter(name="bundleKey")
    def bundle_key(self) -> Optional[pulumi.Input[str]]:
        """
        Key used for internationalization
        """
        return pulumi.get(self, "bundle_key")

    @bundle_key.setter
    def bundle_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bundle_key", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Role description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of Role.
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
    @pulumi.getter(name="readOnly")
    def read_only(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether this role is read-only
        """
        return pulumi.get(self, "read_only")

    @read_only.setter
    def read_only(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "read_only", value)

    @property
    @pulumi.getter
    def rights(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Set of rights assigned to this role
        """
        return pulumi.get(self, "rights")

    @rights.setter
    def rights(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "rights", value)


class Role(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 rights: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Create a Role resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Role description
        :param pulumi.Input[str] name: Name of Role.
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[Sequence[pulumi.Input[str]]] rights: Set of rights assigned to this role
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RoleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Role resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param RoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 rights: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RoleArgs.__new__(RoleArgs)

            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["org"] = org
            __props__.__dict__["rights"] = rights
            __props__.__dict__["bundle_key"] = None
            __props__.__dict__["read_only"] = None
        super(Role, __self__).__init__(
            'vcd:index/role:Role',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            bundle_key: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org: Optional[pulumi.Input[str]] = None,
            read_only: Optional[pulumi.Input[bool]] = None,
            rights: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'Role':
        """
        Get an existing Role resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] bundle_key: Key used for internationalization
        :param pulumi.Input[str] description: Role description
        :param pulumi.Input[str] name: Name of Role.
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[bool] read_only: Whether this role is read-only
        :param pulumi.Input[Sequence[pulumi.Input[str]]] rights: Set of rights assigned to this role
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RoleState.__new__(_RoleState)

        __props__.__dict__["bundle_key"] = bundle_key
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["org"] = org
        __props__.__dict__["read_only"] = read_only
        __props__.__dict__["rights"] = rights
        return Role(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="bundleKey")
    def bundle_key(self) -> pulumi.Output[str]:
        """
        Key used for internationalization
        """
        return pulumi.get(self, "bundle_key")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Role description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of Role.
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
    @pulumi.getter(name="readOnly")
    def read_only(self) -> pulumi.Output[bool]:
        """
        Whether this role is read-only
        """
        return pulumi.get(self, "read_only")

    @property
    @pulumi.getter
    def rights(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Set of rights assigned to this role
        """
        return pulumi.get(self, "rights")
