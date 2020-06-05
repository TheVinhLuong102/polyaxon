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

/**
 * The V1IO model module.
 * @module model/V1IO
 * @version 1.0.94
 */
class V1IO {
    /**
     * Constructs a new <code>V1IO</code>.
     * @alias module:model/V1IO
     */
    constructor() { 
        
        V1IO.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1IO</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1IO} obj Optional instance to populate.
     * @return {module:model/V1IO} The populated <code>V1IO</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1IO();

            if (data.hasOwnProperty('name')) {
                obj['name'] = ApiClient.convertToType(data['name'], 'String');
            }
            if (data.hasOwnProperty('description')) {
                obj['description'] = ApiClient.convertToType(data['description'], 'String');
            }
            if (data.hasOwnProperty('iotype')) {
                obj['iotype'] = ApiClient.convertToType(data['iotype'], 'String');
            }
            if (data.hasOwnProperty('value')) {
                obj['value'] = ApiClient.convertToType(data['value'], Object);
            }
            if (data.hasOwnProperty('is_optional')) {
                obj['is_optional'] = ApiClient.convertToType(data['is_optional'], 'Boolean');
            }
            if (data.hasOwnProperty('is_list')) {
                obj['is_list'] = ApiClient.convertToType(data['is_list'], 'Boolean');
            }
            if (data.hasOwnProperty('is_flag')) {
                obj['is_flag'] = ApiClient.convertToType(data['is_flag'], 'Boolean');
            }
            if (data.hasOwnProperty('delay_validation')) {
                obj['delay_validation'] = ApiClient.convertToType(data['delay_validation'], 'Boolean');
            }
            if (data.hasOwnProperty('options')) {
                obj['options'] = ApiClient.convertToType(data['options'], [Object]);
            }
        }
        return obj;
    }


}

/**
 * @member {String} name
 */
V1IO.prototype['name'] = undefined;

/**
 * @member {String} description
 */
V1IO.prototype['description'] = undefined;

/**
 * @member {String} iotype
 */
V1IO.prototype['iotype'] = undefined;

/**
 * @member {Object} value
 */
V1IO.prototype['value'] = undefined;

/**
 * @member {Boolean} is_optional
 */
V1IO.prototype['is_optional'] = undefined;

/**
 * @member {Boolean} is_list
 */
V1IO.prototype['is_list'] = undefined;

/**
 * @member {Boolean} is_flag
 */
V1IO.prototype['is_flag'] = undefined;

/**
 * @member {Boolean} delay_validation
 */
V1IO.prototype['delay_validation'] = undefined;

/**
 * @member {Array.<Object>} options
 */
V1IO.prototype['options'] = undefined;






export default V1IO;

