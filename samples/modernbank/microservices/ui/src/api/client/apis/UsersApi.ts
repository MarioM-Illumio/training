// tslint:disable
/**
 * Modern Bank
 * This is the API exposed to customers of Modern Bank.
 *
 * OpenAPI spec version: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import {
    User,
    UserFromJSON,
    UserToJSON,
} from '../models';

export interface CreateUserRequest {
    body: User;
}

export interface GetUserByUserNameRequest {
    username: string;
}

/**
 * no description
 */
export class UsersApi extends runtime.BaseAPI {

    /**
     * Creates a new user
     * Create a new user identity for a customer
     */
    async createUserRaw(requestParameters: CreateUserRequest): Promise<runtime.ApiResponse<User>> {
        if (requestParameters.body === null || requestParameters.body === undefined) {
            throw new runtime.RequiredError('body','Required parameter requestParameters.body was null or undefined when calling createUser.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        headerParameters['Content-Type'] = 'application/json';

        const response = await this.request({
            path: `/users`,
            method: 'POST',
            headers: headerParameters,
            query: queryParameters,
            body: UserToJSON(requestParameters.body),
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => UserFromJSON(jsonValue));
    }

    /**
     * Creates a new user
     * Create a new user identity for a customer
     */
    async createUser(requestParameters: CreateUserRequest): Promise<User> {
        const response = await this.createUserRaw(requestParameters);
        return await response.value();
    }

    /**
     * Get user by user name
     */
    async getUserByUserNameRaw(requestParameters: GetUserByUserNameRequest): Promise<runtime.ApiResponse<User>> {
        if (requestParameters.username === null || requestParameters.username === undefined) {
            throw new runtime.RequiredError('username','Required parameter requestParameters.username was null or undefined when calling getUserByUserName.');
        }

        const queryParameters: runtime.HTTPQuery = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/users/{username}`.replace(`{${"username"}}`, encodeURIComponent(String(requestParameters.username))),
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        });

        return new runtime.JSONApiResponse(response, (jsonValue) => UserFromJSON(jsonValue));
    }

    /**
     * Get user by user name
     */
    async getUserByUserName(requestParameters: GetUserByUserNameRequest): Promise<User> {
        const response = await this.getUserByUserNameRaw(requestParameters);
        return await response.value();
    }

}