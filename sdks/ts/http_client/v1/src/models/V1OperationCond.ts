// Copyright 2018-2020 Polyaxon, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/* tslint:disable */
/* eslint-disable */
/**
 * Polyaxon SDKs and REST API specification.
 * Polyaxon SDKs and REST API specification.
 *
 * The version of the OpenAPI document: 1.3.3
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    V1IoCond,
    V1IoCondFromJSON,
    V1IoCondFromJSONTyped,
    V1IoCondToJSON,
    V1StatusCond,
    V1StatusCondFromJSON,
    V1StatusCondFromJSONTyped,
    V1StatusCondToJSON,
} from './';

/**
 * 
 * @export
 * @interface V1OperationCond
 */
export interface V1OperationCond {
    /**
     * 
     * @type {V1IoCond}
     * @memberof V1OperationCond
     */
    io_conidtion?: V1IoCond;
    /**
     * 
     * @type {V1StatusCond}
     * @memberof V1OperationCond
     */
    status_condition?: V1StatusCond;
}

export function V1OperationCondFromJSON(json: any): V1OperationCond {
    return V1OperationCondFromJSONTyped(json, false);
}

export function V1OperationCondFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1OperationCond {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'io_conidtion': !exists(json, 'io_conidtion') ? undefined : V1IoCondFromJSON(json['io_conidtion']),
        'status_condition': !exists(json, 'status_condition') ? undefined : V1StatusCondFromJSON(json['status_condition']),
    };
}

export function V1OperationCondToJSON(value?: V1OperationCond | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'io_conidtion': V1IoCondToJSON(value.io_conidtion),
        'status_condition': V1StatusCondToJSON(value.status_condition),
    };
}


