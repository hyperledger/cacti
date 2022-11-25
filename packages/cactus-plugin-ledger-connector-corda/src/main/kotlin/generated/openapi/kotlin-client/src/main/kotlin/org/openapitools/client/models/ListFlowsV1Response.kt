/**
 * Hyperledger Cactus Plugin - Connector Corda
 *
 * Can perform basic tasks on a Corda ledger
 *
 * The version of the OpenAPI document: 0.3.0
 * 
 *
 * Please note:
 * This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * Do not edit this file manually.
 */

@file:Suppress(
    "ArrayInDataClass",
    "EnumEntryName",
    "RemoveRedundantQualifierName",
    "UnusedImport"
)

package org.openapitools.client.models


import com.squareup.moshi.Json

/**
 * 
 *
 * @param flowNames An array of strings storing the names of the flows as returned by the flowList Corda RPC.
 */

data class ListFlowsV1Response (

    /* An array of strings storing the names of the flows as returned by the flowList Corda RPC. */
    @Json(name = "flowNames")
    val flowNames: kotlin.collections.List<kotlin.String> = arrayListOf()

)

