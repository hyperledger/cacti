/*
 * Hyperledger Cactus Plugin - Object Store - IPFS 
 * Contains/describes the Hyperledger Cactus Object Store IPFS plugin.
 *
 * The version of the OpenAPI document: v2.0.0-alpha.2
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package org.openapitools.client.api;

import org.openapitools.client.ApiException;
import org.openapitools.client.model.GetObjectRequestV1;
import org.openapitools.client.model.GetObjectResponseV1;
import org.openapitools.client.model.HasObjectRequestV1;
import org.openapitools.client.model.HasObjectResponseV1;
import org.openapitools.client.model.SetObjectRequestV1;
import org.openapitools.client.model.SetObjectResponseV1;
import org.junit.jupiter.api.Disabled;
import org.junit.jupiter.api.Test;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * API tests for DefaultApi
 */
@Disabled
public class DefaultApiTest {

    private final DefaultApi api = new DefaultApi();

    /**
     * Retrieves an object from the object store.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void getObjectV1Test() throws ApiException {
        GetObjectRequestV1 getObjectRequestV1 = null;
        GetObjectResponseV1 response = api.getObjectV1(getObjectRequestV1);
        // TODO: test validations
    }

    /**
     * Checks the presence of an object in the object store.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void hasObjectV1Test() throws ApiException {
        HasObjectRequestV1 hasObjectRequestV1 = null;
        HasObjectResponseV1 response = api.hasObjectV1(hasObjectRequestV1);
        // TODO: test validations
    }

    /**
     * Sets an object in the object store under the specified key.
     *
     * @throws ApiException if the Api call fails
     */
    @Test
    public void setObjectV1Test() throws ApiException {
        SetObjectRequestV1 setObjectRequestV1 = null;
        SetObjectResponseV1 response = api.setObjectV1(setObjectRequestV1);
        // TODO: test validations
    }

}
