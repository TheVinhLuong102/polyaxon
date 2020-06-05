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
 * The version of the OpenAPI document: 1.0.94
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    V1Environment,
    V1EnvironmentFromJSON,
    V1EnvironmentFromJSONTyped,
    V1EnvironmentToJSON,
    V1Init,
    V1InitFromJSON,
    V1InitFromJSONTyped,
    V1InitToJSON,
} from './';

/**
 * 
 * @export
 * @interface V1Job
 */
export interface V1Job {
    /**
     * 
     * @type {string}
     * @memberof V1Job
     */
    kind?: string;
    /**
     * 
     * @type {V1Environment}
     * @memberof V1Job
     */
    environment?: V1Environment;
    /**
     * 
     * @type {Array<string>}
     * @memberof V1Job
     */
    connections?: Array<string>;
    /**
     * Volumes is a list of volumes that can be mounted.
     * @type {Array<object>}
     * @memberof V1Job
     */
    volumes?: Array<object>;
    /**
     * 
     * @type {Array<V1Init>}
     * @memberof V1Job
     */
    init?: Array<V1Init>;
    /**
     * 
     * @type {Array<object>}
     * @memberof V1Job
     */
    sidecars?: Array<object>;
    /**
     * 
     * @type {object}
     * @memberof V1Job
     */
    container?: object;
}

export function V1JobFromJSON(json: any): V1Job {
    return V1JobFromJSONTyped(json, false);
}

export function V1JobFromJSONTyped(json: any, ignoreDiscriminator: boolean): V1Job {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'kind': !exists(json, 'kind') ? undefined : json['kind'],
        'environment': !exists(json, 'environment') ? undefined : V1EnvironmentFromJSON(json['environment']),
        'connections': !exists(json, 'connections') ? undefined : json['connections'],
        'volumes': !exists(json, 'volumes') ? undefined : json['volumes'],
        'init': !exists(json, 'init') ? undefined : ((json['init'] as Array<any>).map(V1InitFromJSON)),
        'sidecars': !exists(json, 'sidecars') ? undefined : json['sidecars'],
        'container': !exists(json, 'container') ? undefined : json['container'],
    };
}

export function V1JobToJSON(value?: V1Job | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'kind': value.kind,
        'environment': V1EnvironmentToJSON(value.environment),
        'connections': value.connections,
        'volumes': value.volumes,
        'init': value.init === undefined ? undefined : ((value.init as Array<any>).map(V1InitToJSON)),
        'sidecars': value.sidecars,
        'container': value.container,
    };
}


