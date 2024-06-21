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
 * Response that corresponds to Fabric SDK 'filtered' EventType.
 *
 * @param filteredBlock Filtered commited block.
 */


data class WatchBlocksFilteredResponseV1 (

    /* Filtered commited block. */
    @Json(name = "filteredBlock")
    val filteredBlock: kotlin.Any?

)
