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

from marshmallow import fields, validate

from polyaxon.polyflow.run.job import JobSchema, V1Job
from polyaxon.polyflow.run.kinds import V1RunKind


class WatchDogSchema(JobSchema):
    kind = fields.Str(allow_none=True, validate=validate.Equal(V1RunKind.WATCHDOG))

    @staticmethod
    def schema_config():
        return V1WatchDog


class V1WatchDog(V1Job):
    SCHEMA = WatchDogSchema
    IDENTIFIER = V1RunKind.WATCHDOG
    REDUCED_ATTRIBUTES = [
        "kind",
        "container",
        "environment",
        "init",
        "sidecars",
        "connections",
        "volumes",
    ]
