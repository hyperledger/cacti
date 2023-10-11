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
 * @param chatName 
 * @param otherMember 
 * @param message 
 * @param numberOfRecords 
 */

data class StartFlowV5RequestRequestBody (

    @Json(name = "chatName")
    val chatName: kotlin.String? = null,

    @Json(name = "otherMember")
    val otherMember: kotlin.String? = null,

    @Json(name = "message")
    val message: kotlin.String? = null,

    @Json(name = "numberOfRecords")
    val numberOfRecords: kotlin.String? = null

)

