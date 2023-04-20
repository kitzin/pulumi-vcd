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

__all__ = ['CatalogAccessControlArgs', 'CatalogAccessControl']

@pulumi.input_type
class CatalogAccessControlArgs:
    def __init__(__self__, *,
                 catalog_id: pulumi.Input[str],
                 shared_with_everyone: pulumi.Input[bool],
                 everyone_access_level: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 shared_withs: Optional[pulumi.Input[Sequence[pulumi.Input['CatalogAccessControlSharedWithArgs']]]] = None):
        """
        The set of arguments for constructing a CatalogAccessControl resource.
        :param pulumi.Input[str] catalog_id: The ID of Catalog to use
        :param pulumi.Input[bool] shared_with_everyone: Whether the Catalog is shared with everyone
        :param pulumi.Input[str] everyone_access_level: Access level when the Catalog is shared with everyone (only ReadOnly is available). Required when shared_with_everyone
               is set
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        """
        pulumi.set(__self__, "catalog_id", catalog_id)
        pulumi.set(__self__, "shared_with_everyone", shared_with_everyone)
        if everyone_access_level is not None:
            pulumi.set(__self__, "everyone_access_level", everyone_access_level)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if shared_withs is not None:
            pulumi.set(__self__, "shared_withs", shared_withs)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> pulumi.Input[str]:
        """
        The ID of Catalog to use
        """
        return pulumi.get(self, "catalog_id")

    @catalog_id.setter
    def catalog_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "catalog_id", value)

    @property
    @pulumi.getter(name="sharedWithEveryone")
    def shared_with_everyone(self) -> pulumi.Input[bool]:
        """
        Whether the Catalog is shared with everyone
        """
        return pulumi.get(self, "shared_with_everyone")

    @shared_with_everyone.setter
    def shared_with_everyone(self, value: pulumi.Input[bool]):
        pulumi.set(self, "shared_with_everyone", value)

    @property
    @pulumi.getter(name="everyoneAccessLevel")
    def everyone_access_level(self) -> Optional[pulumi.Input[str]]:
        """
        Access level when the Catalog is shared with everyone (only ReadOnly is available). Required when shared_with_everyone
        is set
        """
        return pulumi.get(self, "everyone_access_level")

    @everyone_access_level.setter
    def everyone_access_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "everyone_access_level", value)

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
    @pulumi.getter(name="sharedWiths")
    def shared_withs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CatalogAccessControlSharedWithArgs']]]]:
        return pulumi.get(self, "shared_withs")

    @shared_withs.setter
    def shared_withs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CatalogAccessControlSharedWithArgs']]]]):
        pulumi.set(self, "shared_withs", value)


@pulumi.input_type
class _CatalogAccessControlState:
    def __init__(__self__, *,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 everyone_access_level: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 shared_with_everyone: Optional[pulumi.Input[bool]] = None,
                 shared_withs: Optional[pulumi.Input[Sequence[pulumi.Input['CatalogAccessControlSharedWithArgs']]]] = None):
        """
        Input properties used for looking up and filtering CatalogAccessControl resources.
        :param pulumi.Input[str] catalog_id: The ID of Catalog to use
        :param pulumi.Input[str] everyone_access_level: Access level when the Catalog is shared with everyone (only ReadOnly is available). Required when shared_with_everyone
               is set
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[bool] shared_with_everyone: Whether the Catalog is shared with everyone
        """
        if catalog_id is not None:
            pulumi.set(__self__, "catalog_id", catalog_id)
        if everyone_access_level is not None:
            pulumi.set(__self__, "everyone_access_level", everyone_access_level)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if shared_with_everyone is not None:
            pulumi.set(__self__, "shared_with_everyone", shared_with_everyone)
        if shared_withs is not None:
            pulumi.set(__self__, "shared_withs", shared_withs)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of Catalog to use
        """
        return pulumi.get(self, "catalog_id")

    @catalog_id.setter
    def catalog_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "catalog_id", value)

    @property
    @pulumi.getter(name="everyoneAccessLevel")
    def everyone_access_level(self) -> Optional[pulumi.Input[str]]:
        """
        Access level when the Catalog is shared with everyone (only ReadOnly is available). Required when shared_with_everyone
        is set
        """
        return pulumi.get(self, "everyone_access_level")

    @everyone_access_level.setter
    def everyone_access_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "everyone_access_level", value)

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
    @pulumi.getter(name="sharedWithEveryone")
    def shared_with_everyone(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the Catalog is shared with everyone
        """
        return pulumi.get(self, "shared_with_everyone")

    @shared_with_everyone.setter
    def shared_with_everyone(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "shared_with_everyone", value)

    @property
    @pulumi.getter(name="sharedWiths")
    def shared_withs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CatalogAccessControlSharedWithArgs']]]]:
        return pulumi.get(self, "shared_withs")

    @shared_withs.setter
    def shared_withs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CatalogAccessControlSharedWithArgs']]]]):
        pulumi.set(self, "shared_withs", value)


class CatalogAccessControl(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 everyone_access_level: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 shared_with_everyone: Optional[pulumi.Input[bool]] = None,
                 shared_withs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CatalogAccessControlSharedWithArgs']]]]] = None,
                 __props__=None):
        """
        Create a CatalogAccessControl resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] catalog_id: The ID of Catalog to use
        :param pulumi.Input[str] everyone_access_level: Access level when the Catalog is shared with everyone (only ReadOnly is available). Required when shared_with_everyone
               is set
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[bool] shared_with_everyone: Whether the Catalog is shared with everyone
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CatalogAccessControlArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a CatalogAccessControl resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param CatalogAccessControlArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CatalogAccessControlArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog_id: Optional[pulumi.Input[str]] = None,
                 everyone_access_level: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 shared_with_everyone: Optional[pulumi.Input[bool]] = None,
                 shared_withs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CatalogAccessControlSharedWithArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CatalogAccessControlArgs.__new__(CatalogAccessControlArgs)

            if catalog_id is None and not opts.urn:
                raise TypeError("Missing required property 'catalog_id'")
            __props__.__dict__["catalog_id"] = catalog_id
            __props__.__dict__["everyone_access_level"] = everyone_access_level
            __props__.__dict__["org"] = org
            if shared_with_everyone is None and not opts.urn:
                raise TypeError("Missing required property 'shared_with_everyone'")
            __props__.__dict__["shared_with_everyone"] = shared_with_everyone
            __props__.__dict__["shared_withs"] = shared_withs
        super(CatalogAccessControl, __self__).__init__(
            'vcd:index/catalogAccessControl:CatalogAccessControl',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            catalog_id: Optional[pulumi.Input[str]] = None,
            everyone_access_level: Optional[pulumi.Input[str]] = None,
            org: Optional[pulumi.Input[str]] = None,
            shared_with_everyone: Optional[pulumi.Input[bool]] = None,
            shared_withs: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['CatalogAccessControlSharedWithArgs']]]]] = None) -> 'CatalogAccessControl':
        """
        Get an existing CatalogAccessControl resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] catalog_id: The ID of Catalog to use
        :param pulumi.Input[str] everyone_access_level: Access level when the Catalog is shared with everyone (only ReadOnly is available). Required when shared_with_everyone
               is set
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        :param pulumi.Input[bool] shared_with_everyone: Whether the Catalog is shared with everyone
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CatalogAccessControlState.__new__(_CatalogAccessControlState)

        __props__.__dict__["catalog_id"] = catalog_id
        __props__.__dict__["everyone_access_level"] = everyone_access_level
        __props__.__dict__["org"] = org
        __props__.__dict__["shared_with_everyone"] = shared_with_everyone
        __props__.__dict__["shared_withs"] = shared_withs
        return CatalogAccessControl(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> pulumi.Output[str]:
        """
        The ID of Catalog to use
        """
        return pulumi.get(self, "catalog_id")

    @property
    @pulumi.getter(name="everyoneAccessLevel")
    def everyone_access_level(self) -> pulumi.Output[Optional[str]]:
        """
        Access level when the Catalog is shared with everyone (only ReadOnly is available). Required when shared_with_everyone
        is set
        """
        return pulumi.get(self, "everyone_access_level")

    @property
    @pulumi.getter
    def org(self) -> pulumi.Output[Optional[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        different organizations
        """
        return pulumi.get(self, "org")

    @property
    @pulumi.getter(name="sharedWithEveryone")
    def shared_with_everyone(self) -> pulumi.Output[bool]:
        """
        Whether the Catalog is shared with everyone
        """
        return pulumi.get(self, "shared_with_everyone")

    @property
    @pulumi.getter(name="sharedWiths")
    def shared_withs(self) -> pulumi.Output[Optional[Sequence['outputs.CatalogAccessControlSharedWith']]]:
        return pulumi.get(self, "shared_withs")
