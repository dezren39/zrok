/* tslint:disable */
/* eslint-disable */
/**
 * zrok
 * zrok client access
 *
 * The version of the OpenAPI document: 0.3.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface AccessResponse
 */
export interface AccessResponse {
    /**
     * 
     * @type {string}
     * @memberof AccessResponse
     */
    frontendToken?: string;
    /**
     * 
     * @type {string}
     * @memberof AccessResponse
     */
    backendMode?: string;
}

/**
 * Check if a given object implements the AccessResponse interface.
 */
export function instanceOfAccessResponse(value: object): boolean {
    let isInstance = true;

    return isInstance;
}

export function AccessResponseFromJSON(json: any): AccessResponse {
    return AccessResponseFromJSONTyped(json, false);
}

export function AccessResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): AccessResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'frontendToken': !exists(json, 'frontendToken') ? undefined : json['frontendToken'],
        'backendMode': !exists(json, 'backendMode') ? undefined : json['backendMode'],
    };
}

export function AccessResponseToJSON(value?: AccessResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'frontendToken': value.frontendToken,
        'backendMode': value.backendMode,
    };
}

