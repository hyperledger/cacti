/**
 * Hyperledger Cactus Example - Carbon Accounting App
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
 * @param accountFound Indicates whether the account was found or not.
 * @param spenderFound Indicates whether the spender was found or not.
 */

data class DaoTokenGetAllowanceNotFoundResponse (

    /* Indicates whether the account was found or not. */
    @Json(name = "accountFound")
    val accountFound: kotlin.Boolean,

    /* Indicates whether the spender was found or not. */
    @Json(name = "spenderFound")
    val spenderFound: kotlin.Boolean

)

