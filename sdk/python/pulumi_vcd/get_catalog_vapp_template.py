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
    'GetCatalogVappTemplateResult',
    'AwaitableGetCatalogVappTemplateResult',
    'get_catalog_vapp_template',
    'get_catalog_vapp_template_output',
]

@pulumi.output_type
class GetCatalogVappTemplateResult:
    """
    A collection of values returned by getCatalogVappTemplate.
    """
    def __init__(__self__, catalog_id=None, created=None, description=None, filter=None, id=None, metadata=None, metadata_entries=None, name=None, org=None, vdc_id=None, vm_names=None):
        if catalog_id and not isinstance(catalog_id, str):
            raise TypeError("Expected argument 'catalog_id' to be a str")
        pulumi.set(__self__, "catalog_id", catalog_id)
        if created and not isinstance(created, str):
            raise TypeError("Expected argument 'created' to be a str")
        pulumi.set(__self__, "created", created)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if filter and not isinstance(filter, dict):
            raise TypeError("Expected argument 'filter' to be a dict")
        pulumi.set(__self__, "filter", filter)
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
        if vdc_id and not isinstance(vdc_id, str):
            raise TypeError("Expected argument 'vdc_id' to be a str")
        pulumi.set(__self__, "vdc_id", vdc_id)
        if vm_names and not isinstance(vm_names, list):
            raise TypeError("Expected argument 'vm_names' to be a list")
        pulumi.set(__self__, "vm_names", vm_names)

    @property
    @pulumi.getter(name="catalogId")
    def catalog_id(self) -> Optional[str]:
        return pulumi.get(self, "catalog_id")

    @property
    @pulumi.getter
    def created(self) -> str:
        return pulumi.get(self, "created")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def filter(self) -> Optional['outputs.GetCatalogVappTemplateFilterResult']:
        return pulumi.get(self, "filter")

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
    def metadata_entries(self) -> Sequence['outputs.GetCatalogVappTemplateMetadataEntryResult']:
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
    @pulumi.getter(name="vdcId")
    def vdc_id(self) -> Optional[str]:
        return pulumi.get(self, "vdc_id")

    @property
    @pulumi.getter(name="vmNames")
    def vm_names(self) -> Sequence[str]:
        return pulumi.get(self, "vm_names")


class AwaitableGetCatalogVappTemplateResult(GetCatalogVappTemplateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCatalogVappTemplateResult(
            catalog_id=self.catalog_id,
            created=self.created,
            description=self.description,
            filter=self.filter,
            id=self.id,
            metadata=self.metadata,
            metadata_entries=self.metadata_entries,
            name=self.name,
            org=self.org,
            vdc_id=self.vdc_id,
            vm_names=self.vm_names)


def get_catalog_vapp_template(catalog_id: Optional[str] = None,
                              filter: Optional[pulumi.InputType['GetCatalogVappTemplateFilterArgs']] = None,
                              name: Optional[str] = None,
                              org: Optional[str] = None,
                              vdc_id: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCatalogVappTemplateResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['catalogId'] = catalog_id
    __args__['filter'] = filter
    __args__['name'] = name
    __args__['org'] = org
    __args__['vdcId'] = vdc_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getCatalogVappTemplate:getCatalogVappTemplate', __args__, opts=opts, typ=GetCatalogVappTemplateResult).value

    return AwaitableGetCatalogVappTemplateResult(
        catalog_id=__ret__.catalog_id,
        created=__ret__.created,
        description=__ret__.description,
        filter=__ret__.filter,
        id=__ret__.id,
        metadata=__ret__.metadata,
        metadata_entries=__ret__.metadata_entries,
        name=__ret__.name,
        org=__ret__.org,
        vdc_id=__ret__.vdc_id,
        vm_names=__ret__.vm_names)


@_utilities.lift_output_func(get_catalog_vapp_template)
def get_catalog_vapp_template_output(catalog_id: Optional[pulumi.Input[Optional[str]]] = None,
                                     filter: Optional[pulumi.Input[Optional[pulumi.InputType['GetCatalogVappTemplateFilterArgs']]]] = None,
                                     name: Optional[pulumi.Input[Optional[str]]] = None,
                                     org: Optional[pulumi.Input[Optional[str]]] = None,
                                     vdc_id: Optional[pulumi.Input[Optional[str]]] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCatalogVappTemplateResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
