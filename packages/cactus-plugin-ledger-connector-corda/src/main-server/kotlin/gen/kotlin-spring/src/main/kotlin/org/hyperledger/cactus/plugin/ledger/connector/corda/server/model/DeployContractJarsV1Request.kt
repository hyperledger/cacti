package org.hyperledger.cactus.plugin.ledger.connector.corda.server.model

import java.util.Objects
import com.fasterxml.jackson.annotation.JsonProperty
import org.hyperledger.cactus.plugin.ledger.connector.corda.server.model.CordappDeploymentConfig
import org.hyperledger.cactus.plugin.ledger.connector.corda.server.model.JarFile
import jakarta.validation.constraints.DecimalMax
import jakarta.validation.constraints.DecimalMin
import jakarta.validation.constraints.Email
import jakarta.validation.constraints.Max
import jakarta.validation.constraints.Min
import jakarta.validation.constraints.NotNull
import jakarta.validation.constraints.Pattern
import jakarta.validation.constraints.Size
import jakarta.validation.Valid

/**
 * 
 * @param cordappDeploymentConfigs The list of deployment configurations pointing to the nodes where the provided cordapp jar files are to be deployed .
 * @param jarFiles 
 */
data class DeployContractJarsV1Request(

    @field:Valid
    @get:JsonProperty("cordappDeploymentConfigs", required = true) val cordappDeploymentConfigs: kotlin.collections.List<CordappDeploymentConfig> = arrayListOf(),

    @field:Valid
    @get:JsonProperty("jarFiles", required = true) val jarFiles: kotlin.collections.List<JarFile>
) {

}
