/**
 * Hyperledger Cactus Example - Supply Chain App
 *
 * Demonstrates how a business use case can be satisfied with Cactus when multiple distinct ledgers are involved.
 *
 * The version of the OpenAPI document: 0.2.0
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
 * @param callOutput 
 * @param transactionReceipt 
 */

data class InsertBookshelfResponse (

    @Json(name = "callOutput")
    val callOutput: kotlin.collections.Map<kotlin.String, kotlin.Any>? = null,

    @Json(name = "transactionReceipt")
    val transactionReceipt: kotlin.collections.Map<kotlin.String, kotlin.Any>? = null

)

