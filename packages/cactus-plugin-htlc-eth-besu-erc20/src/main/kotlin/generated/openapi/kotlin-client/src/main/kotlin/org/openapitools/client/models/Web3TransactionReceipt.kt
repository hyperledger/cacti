/**
 *
 * Please note:
 * This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * Do not edit this file manually.
 *
 */

@file:Suppress(
    "ArrayInDataClass",
    "EnumEntryName",
    "RemoveRedundantQualifierName",
    "UnusedImport"
)

package org.openapitools.client.models


import com.squareup.moshi.Json
import com.squareup.moshi.JsonClass

/**
 * 
 *
 * @param status 
 * @param transactionHash 
 * @param transactionIndex 
 * @param blockHash 
 * @param blockNumber 
 * @param gasUsed 
 * @param from 
 * @param to 
 * @param contractAddress 
 */


data class Web3TransactionReceipt (

    @Json(name = "status")
    val status: kotlin.Boolean,

    @Json(name = "transactionHash")
    val transactionHash: kotlin.String,

    @Json(name = "transactionIndex")
    val transactionIndex: java.math.BigDecimal,

    @Json(name = "blockHash")
    val blockHash: kotlin.String,

    @Json(name = "blockNumber")
    val blockNumber: java.math.BigDecimal,

    @Json(name = "gasUsed")
    val gasUsed: java.math.BigDecimal,

    @Json(name = "from")
    val from: kotlin.String,

    @Json(name = "to")
    val to: kotlin.String,

    @Json(name = "contractAddress")
    val contractAddress: kotlin.String? = null

) : kotlin.collections.HashMap<String, kotlin.Any>()

