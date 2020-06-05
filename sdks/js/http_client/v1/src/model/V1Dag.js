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
 *
 */

import ApiClient from '../ApiClient';
import V1Component from './V1Component';
import V1Environment from './V1Environment';
import V1Operation from './V1Operation';

/**
 * The V1Dag model module.
 * @module model/V1Dag
 * @version 1.0.94
 */
class V1Dag {
    /**
     * Constructs a new <code>V1Dag</code>.
     * @alias module:model/V1Dag
     */
    constructor() { 
        
        V1Dag.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1Dag</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Dag} obj Optional instance to populate.
     * @return {module:model/V1Dag} The populated <code>V1Dag</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Dag();

            if (data.hasOwnProperty('kind')) {
                obj['kind'] = ApiClient.convertToType(data['kind'], 'String');
            }
            if (data.hasOwnProperty('operations')) {
                obj['operations'] = ApiClient.convertToType(data['operations'], [V1Operation]);
            }
            if (data.hasOwnProperty('components')) {
                obj['components'] = ApiClient.convertToType(data['components'], [V1Component]);
            }
            if (data.hasOwnProperty('concurrency')) {
                obj['concurrency'] = ApiClient.convertToType(data['concurrency'], 'Number');
            }
            if (data.hasOwnProperty('early_stopping')) {
                obj['early_stopping'] = ApiClient.convertToType(data['early_stopping'], [Object]);
            }
            if (data.hasOwnProperty('environment')) {
                obj['environment'] = V1Environment.constructFromObject(data['environment']);
            }
            if (data.hasOwnProperty('connections')) {
                obj['connections'] = ApiClient.convertToType(data['connections'], ['String']);
            }
            if (data.hasOwnProperty('volumes')) {
                obj['volumes'] = ApiClient.convertToType(data['volumes'], [Object]);
            }
        }
        return obj;
    }


}

/**
 * @member {String} kind
 * @default 'dag'
 */
V1Dag.prototype['kind'] = 'dag';

/**
 * @member {Array.<module:model/V1Operation>} operations
 */
V1Dag.prototype['operations'] = undefined;

/**
 * @member {Array.<module:model/V1Component>} components
 */
V1Dag.prototype['components'] = undefined;

/**
 * @member {Number} concurrency
 */
V1Dag.prototype['concurrency'] = undefined;

/**
 * @member {Array.<Object>} early_stopping
 */
V1Dag.prototype['early_stopping'] = undefined;

/**
 * @member {module:model/V1Environment} environment
 */
V1Dag.prototype['environment'] = undefined;

/**
 * @member {Array.<String>} connections
 */
V1Dag.prototype['connections'] = undefined;

/**
 * Volumes is a list of volumes that can be mounted.
 * @member {Array.<Object>} volumes
 */
V1Dag.prototype['volumes'] = undefined;






export default V1Dag;

