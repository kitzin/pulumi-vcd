# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('vcd')


class _ExportableConfig(types.ModuleType):
    @property
    def allow_unverified_ssl(self) -> Optional[bool]:
        """
        If set, VCDClient will permit unverifiable SSL certificates.
        """
        return __config__.get_bool('allowUnverifiedSsl')

    @property
    def api_token(self) -> Optional[str]:
        """
        The API token used instead of username/password for VCD API operations. (Requires VCD 10.3.1+)
        """
        return __config__.get('apiToken')

    @property
    def auth_type(self) -> Optional[str]:
        """
        'integrated', 'saml_adfs', 'token', and 'api_token' are the only ones supported now. 'integrated' is default.
        """
        return __config__.get('authType')

    @property
    def import_separator(self) -> Optional[str]:
        """
        Defines the import separation string to be used with 'terraform import'
        """
        return __config__.get('importSeparator')

    @property
    def logging(self) -> Optional[bool]:
        """
        If set, it will enable logging of API requests and responses
        """
        return __config__.get_bool('logging')

    @property
    def logging_file(self) -> Optional[str]:
        """
        Defines the full name of the logging file for API calls (requires 'logging')
        """
        return __config__.get('loggingFile')

    @property
    def max_retry_timeout(self) -> Optional[int]:
        """
        Max num seconds to wait for successful response when operating on resources within vCloud (defaults to 60)
        """
        return __config__.get_int('maxRetryTimeout')

    @property
    def org(self) -> Optional[str]:
        """
        The VCD Org for API operations
        """
        return __config__.get('org')

    @property
    def password(self) -> Optional[str]:
        """
        The user password for VCD API operations.
        """
        return __config__.get('password')

    @property
    def saml_adfs_rpt_id(self) -> Optional[str]:
        """
        Allows to specify custom Relaying Party Trust Identifier for auth_type=saml_adfs
        """
        return __config__.get('samlAdfsRptId')

    @property
    def sysorg(self) -> Optional[str]:
        """
        The VCD Org for user authentication
        """
        return __config__.get('sysorg')

    @property
    def token(self) -> Optional[str]:
        """
        The token used instead of username/password for VCD API operations.
        """
        return __config__.get('token')

    @property
    def url(self) -> Optional[str]:
        """
        The VCD url for VCD API operations.
        """
        return __config__.get('url')

    @property
    def user(self) -> Optional[str]:
        """
        The user name for VCD API operations.
        """
        return __config__.get('user')

    @property
    def vdc(self) -> Optional[str]:
        """
        The VDC for API operations
        """
        return __config__.get('vdc')

