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
 * The version of the OpenAPI document: 1.3.3
 * Contact: contact@polyaxon.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 *
 */

import ApiClient from '../ApiClient';
import V1CronSchedule from './V1CronSchedule';
import V1DateTimeSchedule from './V1DateTimeSchedule';
import V1IntervalSchedule from './V1IntervalSchedule';

/**
 * The V1Schedule model module.
 * @module model/V1Schedule
 * @version 1.3.3
 */
class V1Schedule {
    /**
     * Constructs a new <code>V1Schedule</code>.
     * @alias module:model/V1Schedule
     */
    constructor() { 
        
        V1Schedule.initialize(this);
    }

    /**
     * Initializes the fields of this object.
     * This method is used by the constructors of any subclasses, in order to implement multiple inheritance (mix-ins).
     * Only for internal use.
     */
    static initialize(obj) { 
    }

    /**
     * Constructs a <code>V1Schedule</code> from a plain JavaScript object, optionally creating a new instance.
     * Copies all relevant properties from <code>data</code> to <code>obj</code> if supplied or a new instance if not.
     * @param {Object} data The plain JavaScript object bearing properties of interest.
     * @param {module:model/V1Schedule} obj Optional instance to populate.
     * @return {module:model/V1Schedule} The populated <code>V1Schedule</code> instance.
     */
    static constructFromObject(data, obj) {
        if (data) {
            obj = obj || new V1Schedule();

            if (data.hasOwnProperty('cron')) {
                obj['cron'] = V1CronSchedule.constructFromObject(data['cron']);
            }
            if (data.hasOwnProperty('datetime')) {
                obj['datetime'] = V1DateTimeSchedule.constructFromObject(data['datetime']);
            }
            if (data.hasOwnProperty('interval')) {
                obj['interval'] = V1IntervalSchedule.constructFromObject(data['interval']);
            }
        }
        return obj;
    }


}

/**
 * @member {module:model/V1CronSchedule} cron
 */
V1Schedule.prototype['cron'] = undefined;

/**
 * @member {module:model/V1DateTimeSchedule} datetime
 */
V1Schedule.prototype['datetime'] = undefined;

/**
 * @member {module:model/V1IntervalSchedule} interval
 */
V1Schedule.prototype['interval'] = undefined;






export default V1Schedule;

