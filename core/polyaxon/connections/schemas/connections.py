#!/usr/bin/python
#
# Copyright 2018-2020 Polyaxon, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import polyaxon_sdk

from marshmallow import INCLUDE, ValidationError, fields

from polyaxon.connections.kinds import V1ConnectionKind
from polyaxon.schemas.base import BaseCamelSchema, BaseConfig, BaseOneOfSchema


class BucketConnectionSchema(BaseCamelSchema):
    bucket = fields.Str(required=True)

    @staticmethod
    def schema_config():
        return V1BucketConnection


class V1BucketConnection(BaseConfig, polyaxon_sdk.V1BucketConnection):
    SCHEMA = BucketConnectionSchema
    IDENTIFIER = "bucket"

    def patch(self, schema: "V1BucketConnection"):
        self.bucket = schema.bucket or self.bucket


class ClaimConnectionSchema(BaseCamelSchema):
    volume_claim = fields.Str(required=True)
    mount_path = fields.Str(required=True)
    read_only = fields.Bool(allow_none=True)

    @staticmethod
    def schema_config():
        return V1ClaimConnection


class V1ClaimConnection(BaseConfig, polyaxon_sdk.V1ClaimConnection):
    SCHEMA = ClaimConnectionSchema
    IDENTIFIER = "volume_claim"

    def patch(self, schema: "V1ClaimConnection"):
        self.volume_claim = schema.volume_claim or self.volume_claim
        self.mount_path = schema.mount_path or self.mount_path
        self.read_only = schema.read_only or self.read_only


class HostPathConnectionSchema(BaseCamelSchema):
    host_path = fields.Str(required=True)
    mount_path = fields.Str(required=True)
    read_only = fields.Bool(allow_none=True)

    @staticmethod
    def schema_config():
        return V1HostPathConnection


class V1HostPathConnection(BaseConfig, polyaxon_sdk.V1HostPathConnection):
    SCHEMA = HostPathConnectionSchema
    IDENTIFIER = "host_path"

    def patch(self, schema: "V1HostPathConnection"):
        self.host_path = schema.host_path or self.host_path
        self.mount_path = schema.mount_path or self.mount_path
        self.read_only = schema.read_only or self.read_only


class HostConnectionSchema(BaseCamelSchema):
    url = fields.Str(required=True)
    insecure = fields.Bool(allow_none=True)

    @staticmethod
    def schema_config():
        return V1HostConnection


class V1HostConnection(BaseConfig, polyaxon_sdk.V1HostConnection):
    SCHEMA = HostConnectionSchema
    IDENTIFIER = "host"

    def patch(self, schema: "V1HostConnection"):
        self.url = schema.url or self.url
        self.insecure = schema.insecure or self.insecure


class GitConnectionSchema(BaseCamelSchema):
    url = fields.Str(allow_none=True)
    revision = fields.Str(allow_none=True)

    @staticmethod
    def schema_config():
        return V1GitConnection


class V1GitConnection(BaseConfig, polyaxon_sdk.V1GitConnection):
    SCHEMA = GitConnectionSchema
    IDENTIFIER = "git"
    REDUCED_ATTRIBUTES = ["url", "revision"]

    def get_name(self):
        if self.url:
            return self.url.split("/")[-1].split(".")[0]
        return None

    def patch(self, schema: "GitConnectionSchema"):
        self.url = schema.url or self.url
        self.revision = schema.revision or self.revision


class CustomConnectionSchema(BaseCamelSchema):
    class Meta(BaseCamelSchema.Meta):
        unknown = INCLUDE

    @staticmethod
    def schema_config():
        return V1CustomConnection


class V1CustomConnection(BaseConfig):
    UNKNOWN_BEHAVIOUR = INCLUDE
    IDENTIFIER = "custom"
    SCHEMA = CustomConnectionSchema

    def __init__(self, **kwargs):
        self._schema_keys = set([])
        for k, v in kwargs.items():
            self._schema_keys.add(k)
            self.__setattr__(k, v)

    @classmethod
    def obj_to_dict(
        cls,
        obj,
        humanize_values=False,
        unknown=None,
        include_kind=False,
        include_version=False,
    ):
        value = super().obj_to_dict(
            obj=obj,
            humanize_values=humanize_values,
            unknown=cls.UNKNOWN_BEHAVIOUR,
            include_kind=include_kind,
            include_version=include_version,
        )
        value.update({k: getattr(obj, k) for k in obj._schema_keys})
        return value

    @classmethod
    def from_dict(cls, value, unknown=None):
        return super().from_dict(value=value, unknown=cls.UNKNOWN_BEHAVIOUR)

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, V1CustomConnection):
            return False

        return self.to_dict() == other.to_dict()

    def patch_git(self, schema: "GitConnectionSchema"):
        if schema.url:
            if "url" not in self._schema_keys:
                self._schema_keys.add("url")
            setattr(self, "url", schema.url)
        if schema.revision:
            if "revision" not in self._schema_keys:
                self._schema_keys.add("revision")
            setattr(self, "revision", schema.url)


def validate_connection(kind, definition):
    if kind not in V1ConnectionKind.allowable_values:
        raise ValidationError("Connection with kind {} is not supported.".format(kind))

    if kind in V1ConnectionKind.BLOB_VALUES:
        V1BucketConnection.from_dict(definition)

    if kind == V1ConnectionKind.VOLUME_CLAIM:
        V1ClaimConnection.from_dict(definition)

    if kind == V1ConnectionKind.HOST_PATH:
        V1HostPathConnection.from_dict(definition)

    if kind == V1ConnectionKind.REGISTRY:
        V1HostConnection.from_dict(definition)

    if kind == V1ConnectionKind.GIT:
        V1GitConnection.from_dict(definition)


class ConnectionSchema(BaseOneOfSchema):
    TYPE_FIELD = "kind"
    TYPE_FIELD_REMOVE = True
    UNKNOWN_BEHAVIOUR = INCLUDE

    class Meta:
        unknown = INCLUDE

    SCHEMAS = {
        V1BucketConnection.IDENTIFIER: BucketConnectionSchema,
        V1ClaimConnection.IDENTIFIER: ClaimConnectionSchema,
        V1HostPathConnection.IDENTIFIER: HostPathConnectionSchema,
        V1HostConnection.IDENTIFIER: HostConnectionSchema,
        V1GitConnection.IDENTIFIER: GitConnectionSchema,
        V1CustomConnection.IDENTIFIER: CustomConnectionSchema,
    }
