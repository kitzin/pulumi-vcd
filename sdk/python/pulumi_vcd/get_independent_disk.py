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
    'GetIndependentDiskResult',
    'AwaitableGetIndependentDiskResult',
    'get_independent_disk',
    'get_independent_disk_output',
]

@pulumi.output_type
class GetIndependentDiskResult:
    """
    A collection of values returned by getIndependentDisk.
    """
    def __init__(__self__, attached_vm_ids=None, bus_sub_type=None, bus_type=None, datastore_name=None, description=None, encrypted=None, id=None, iops=None, is_attached=None, metadata=None, metadata_entries=None, name=None, org=None, owner_name=None, sharing_type=None, size_in_mb=None, storage_profile=None, uuid=None, vdc=None):
        if attached_vm_ids and not isinstance(attached_vm_ids, list):
            raise TypeError("Expected argument 'attached_vm_ids' to be a list")
        pulumi.set(__self__, "attached_vm_ids", attached_vm_ids)
        if bus_sub_type and not isinstance(bus_sub_type, str):
            raise TypeError("Expected argument 'bus_sub_type' to be a str")
        pulumi.set(__self__, "bus_sub_type", bus_sub_type)
        if bus_type and not isinstance(bus_type, str):
            raise TypeError("Expected argument 'bus_type' to be a str")
        pulumi.set(__self__, "bus_type", bus_type)
        if datastore_name and not isinstance(datastore_name, str):
            raise TypeError("Expected argument 'datastore_name' to be a str")
        pulumi.set(__self__, "datastore_name", datastore_name)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if encrypted and not isinstance(encrypted, bool):
            raise TypeError("Expected argument 'encrypted' to be a bool")
        pulumi.set(__self__, "encrypted", encrypted)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if iops and not isinstance(iops, int):
            raise TypeError("Expected argument 'iops' to be a int")
        pulumi.set(__self__, "iops", iops)
        if is_attached and not isinstance(is_attached, bool):
            raise TypeError("Expected argument 'is_attached' to be a bool")
        pulumi.set(__self__, "is_attached", is_attached)
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
        if owner_name and not isinstance(owner_name, str):
            raise TypeError("Expected argument 'owner_name' to be a str")
        pulumi.set(__self__, "owner_name", owner_name)
        if sharing_type and not isinstance(sharing_type, str):
            raise TypeError("Expected argument 'sharing_type' to be a str")
        pulumi.set(__self__, "sharing_type", sharing_type)
        if size_in_mb and not isinstance(size_in_mb, int):
            raise TypeError("Expected argument 'size_in_mb' to be a int")
        pulumi.set(__self__, "size_in_mb", size_in_mb)
        if storage_profile and not isinstance(storage_profile, str):
            raise TypeError("Expected argument 'storage_profile' to be a str")
        pulumi.set(__self__, "storage_profile", storage_profile)
        if uuid and not isinstance(uuid, str):
            raise TypeError("Expected argument 'uuid' to be a str")
        pulumi.set(__self__, "uuid", uuid)
        if vdc and not isinstance(vdc, str):
            raise TypeError("Expected argument 'vdc' to be a str")
        pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter(name="attachedVmIds")
    def attached_vm_ids(self) -> Sequence[str]:
        return pulumi.get(self, "attached_vm_ids")

    @property
    @pulumi.getter(name="busSubType")
    def bus_sub_type(self) -> str:
        return pulumi.get(self, "bus_sub_type")

    @property
    @pulumi.getter(name="busType")
    def bus_type(self) -> str:
        return pulumi.get(self, "bus_type")

    @property
    @pulumi.getter(name="datastoreName")
    def datastore_name(self) -> str:
        return pulumi.get(self, "datastore_name")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def encrypted(self) -> bool:
        return pulumi.get(self, "encrypted")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def iops(self) -> int:
        return pulumi.get(self, "iops")

    @property
    @pulumi.getter(name="isAttached")
    def is_attached(self) -> bool:
        return pulumi.get(self, "is_attached")

    @property
    @pulumi.getter
    def metadata(self) -> Mapping[str, Any]:
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter(name="metadataEntries")
    def metadata_entries(self) -> Sequence['outputs.GetIndependentDiskMetadataEntryResult']:
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
    @pulumi.getter(name="ownerName")
    def owner_name(self) -> str:
        return pulumi.get(self, "owner_name")

    @property
    @pulumi.getter(name="sharingType")
    def sharing_type(self) -> str:
        return pulumi.get(self, "sharing_type")

    @property
    @pulumi.getter(name="sizeInMb")
    def size_in_mb(self) -> int:
        return pulumi.get(self, "size_in_mb")

    @property
    @pulumi.getter(name="storageProfile")
    def storage_profile(self) -> str:
        return pulumi.get(self, "storage_profile")

    @property
    @pulumi.getter
    def uuid(self) -> str:
        return pulumi.get(self, "uuid")

    @property
    @pulumi.getter
    def vdc(self) -> Optional[str]:
        return pulumi.get(self, "vdc")


class AwaitableGetIndependentDiskResult(GetIndependentDiskResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetIndependentDiskResult(
            attached_vm_ids=self.attached_vm_ids,
            bus_sub_type=self.bus_sub_type,
            bus_type=self.bus_type,
            datastore_name=self.datastore_name,
            description=self.description,
            encrypted=self.encrypted,
            id=self.id,
            iops=self.iops,
            is_attached=self.is_attached,
            metadata=self.metadata,
            metadata_entries=self.metadata_entries,
            name=self.name,
            org=self.org,
            owner_name=self.owner_name,
            sharing_type=self.sharing_type,
            size_in_mb=self.size_in_mb,
            storage_profile=self.storage_profile,
            uuid=self.uuid,
            vdc=self.vdc)


def get_independent_disk(id: Optional[str] = None,
                         name: Optional[str] = None,
                         org: Optional[str] = None,
                         vdc: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetIndependentDiskResult:
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()
    __args__['id'] = id
    __args__['name'] = name
    __args__['org'] = org
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getIndependentDisk:getIndependentDisk', __args__, opts=opts, typ=GetIndependentDiskResult).value

    return AwaitableGetIndependentDiskResult(
        attached_vm_ids=__ret__.attached_vm_ids,
        bus_sub_type=__ret__.bus_sub_type,
        bus_type=__ret__.bus_type,
        datastore_name=__ret__.datastore_name,
        description=__ret__.description,
        encrypted=__ret__.encrypted,
        id=__ret__.id,
        iops=__ret__.iops,
        is_attached=__ret__.is_attached,
        metadata=__ret__.metadata,
        metadata_entries=__ret__.metadata_entries,
        name=__ret__.name,
        org=__ret__.org,
        owner_name=__ret__.owner_name,
        sharing_type=__ret__.sharing_type,
        size_in_mb=__ret__.size_in_mb,
        storage_profile=__ret__.storage_profile,
        uuid=__ret__.uuid,
        vdc=__ret__.vdc)


@_utilities.lift_output_func(get_independent_disk)
def get_independent_disk_output(id: Optional[pulumi.Input[Optional[str]]] = None,
                                name: Optional[pulumi.Input[Optional[str]]] = None,
                                org: Optional[pulumi.Input[Optional[str]]] = None,
                                vdc: Optional[pulumi.Input[Optional[str]]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetIndependentDiskResult]:
    """
    Use this data source to access information about an existing resource.
    """
    ...
