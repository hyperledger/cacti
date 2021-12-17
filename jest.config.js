module.exports = {
  preset: "ts-jest",
  testEnvironment: "node",
  maxWorkers: 1,
  maxConcurrency: 1,
  setupFilesAfterEnv: ["jest-extended"],
  testTimeout: 60 * 60 * 1000,
  testMatch: [
    `**/cactus-*/src/test/typescript/{unit,integration,benchmark}/**/*.test.ts`,
  ],
  // Ignore the tests that are still using tap/tape for as their test runner
  testPathIgnorePatterns: [
    `./packages/cactus-plugin-keychain-aws-sm/src/test/typescript/integration/openapi/openapi-validation.test.ts`,
    `./packages/cactus-plugin-odap-hermes/src/test/typescript/integration/odap/transfer-complete.test.ts`,
    `./packages/cactus-plugin-odap-hermes/src/test/typescript/integration/odap/odap-api-call.test.ts`,
    `./packages/cactus-plugin-odap-hermes/src/test/typescript/integration/odap/odap-api-call-with-ledger-connector.test.ts`,
    `./packages/cactus-plugin-odap-hermes/src/test/typescript/integration/odap/lock-evidence.test.ts`,
    `./packages/cactus-plugin-odap-hermes/src/test/typescript/integration/odap/commit-preparation.test.ts`,
    `./packages/cactus-plugin-odap-hermes/src/test/typescript/integration/odap/commit-final.test.ts`,
    `./packages/cactus-plugin-ledger-connector-fabric/src/test/typescript/integration/fabric-v2-2-x/deploy-lock-asset.test.ts`,
    `./packages/cactus-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-ledger-connector-besu/lock-contract.test.ts`,
    `./packages/cactus-plugin-ledger-connector-fabric/src/test/typescript/integration/fabric-v2-2-x/run-transaction-with-ws-ids.test.ts`,
    `./packages/cactus-plugin-ledger-connector-corda/src/test/typescript/integration/deploy-cordapp-jars-to-nodes-v4.8-express.test.ts`,
    `./packages/cactus-plugin-ledger-connector-quorum/src/test/typescript/integration/plugin-ledger-connector-quorum/deploy-contract/v2.3.0-invoke-contract-json-object-endpoints.test.ts`,
    `./packages/cactus-plugin-ledger-connector-quorum/src/test/typescript/integration/plugin-ledger-connector-quorum/deploy-contract/v21.4.1-invoke-contract-json-object-endpoints.test.ts`,
    `./packages/cactus-plugin-ledger-connector-quorum/src/test/typescript/integration/plugin-ledger-connector-quorum/deploy-contract/v21.4.1-deploy-contract-from-json-json-object-endpoints.test.ts`,
    `./extensions/cactus-plugin-htlc-coordinator-besu/src/test/typescript/integration/plugin-htlc-coordinator/counterparty-htlc-endpoint.test.ts`,
    `./extensions/cactus-plugin-htlc-coordinator-besu/src/test/typescript/integration/plugin-htlc-coordinator/own-htlc-endpoint.test.ts`,
    `./extensions/cactus-plugin-htlc-coordinator-besu/src/test/typescript/integration/plugin-htlc-coordinator/withdraw-counterparty-endpoint.test.ts`,
    `./extensions/cactus-plugin-htlc-coordinator-besu/src/test/typescript/integration/plugin-htlc-coordinator/refund.test.ts`,
    `./packages/cactus-plugin-ledger-connector-quorum/src/test/typescript/integration/plugin-ledger-connector-quorum/deploy-contract/v2.3.0-deploy-contract-from-json-json-object-endpoints.test.ts`,
    `./packages/cactus-plugin-keychain-memory-wasm/src/test/typescript/unit/plugin-keychain-memory-wasm.test.ts`,
    `./packages/cactus-plugin-ledger-connector-quorum/src/test/typescript/integration/plugin-ledger-connector-quorum/deploy-contract/openapi/openapi-validation-no-keychain.test.ts`,
    `./packages/cactus-plugin-ledger-connector-fabric/src/test/typescript/integration/fabric-v2-2-x/add-orgs.test.ts`,
    `./packages/cactus-plugin-ledger-connector-fabric/src/test/typescript/integration/fabric-v2-2-x/obtain-profiles.test.ts`,
    `./packages/cactus-plugin-keychain-google-sm/src/test/typescript/integration/openapi/openapi-validation.test.ts`,
    `./packages/cactus-plugin-keychain-aws-sm/src/test/typescript/integration/plugin-factory-keychain.test.ts`,
    `./packages/cactus-plugin-keychain-aws-sm/src/test/typescript/integration/plugin-factory-keychain.test.ts`,
    `./packages/cactus-plugin-keychain-azure-kv/src/test/typescript/integration/plugin-keychain-azure-kv.test.ts`,
    `./packages/cactus-plugin-ledger-connector-quorum/src/test/typescript/integration/plugin-ledger-connector-quorum/deploy-contract/v21.4.1-invoke-contract.test.ts`,
    `./packages/cactus-plugin-ledger-connector-quorum/src/test/typescript/integration/plugin-ledger-connector-quorum/deploy-contract/v2.3.0-invoke-contract-json-object.test.ts`,
    `./packages/cactus-plugin-ledger-connector-quorum/src/test/typescript/integration/plugin-ledger-connector-quorum/deploy-contract/v2.3.0-deploy-contract-from-json.test.ts`,
    `./packages/cactus-plugin-ledger-connector-quorum/src/test/typescript/integration/plugin-ledger-connector-quorum/deploy-contract/v21.4.1-deploy-contract-from-json-json-object.test.ts`,
    `./packages/cactus-plugin-ledger-connector-quorum/src/test/typescript/integration/plugin-ledger-connector-quorum/deploy-contract/v2.3.0-deploy-contract-from-json-json-object.test.ts`,
    `./packages/cactus-plugin-ledger-connector-quorum/src/test/typescript/integration/plugin-ledger-connector-quorum/deploy-contract/v2.3.0-invoke-contract.test.ts`,
    `./packages/cactus-plugin-ledger-connector-quorum/src/test/typescript/integration/plugin-ledger-connector-quorum/deploy-contract/openapi/openapi-validation.test.ts`,
    `./packages/cactus-plugin-ledger-connector-quorum/src/test/typescript/integration/plugin-ledger-connector-quorum/deploy-contract/v21.4.1-invoke-contract-json-object.test.ts`,
    `./packages/cactus-test-plugin-htlc-eth-besu-erc20/src/test/typescript/integration/plugin-htlc-eth-besu-erc20/get-status-endpoint.test.ts`,
    `./packages/cactus-test-plugin-htlc-eth-besu-erc20/src/test/typescript/integration/plugin-htlc-eth-besu-erc20/initialize-endpoint.test.ts`,
    `./packages/cactus-test-plugin-htlc-eth-besu-erc20/src/test/typescript/integration/plugin-htlc-eth-besu-erc20/withdraw-endpoint.test.ts`,
    `./packages/cactus-test-plugin-htlc-eth-besu-erc20/src/test/typescript/integration/plugin-htlc-eth-besu-erc20/get-single-status-endpoint.test.ts`,
    `./packages/cactus-test-plugin-htlc-eth-besu-erc20/src/test/typescript/integration/plugin-htlc-eth-besu-erc20/get-single-status-endpoint-invalid.test.ts`,
    `./packages/cactus-test-plugin-htlc-eth-besu-erc20/src/test/typescript/integration/plugin-htlc-eth-besu-erc20/openapi/openapi-validation.test.ts`,
    `./packages/cactus-test-plugin-htlc-eth-besu-erc20/src/test/typescript/integration/plugin-htlc-eth-besu-erc20/refund-endpoint.test.ts`,
    `./packages/cactus-plugin-consortium-manual/src/test/typescript/unit/consortium/get-node-jws-endpoint-v1.test.ts`,
    `./packages/cactus-test-cmd-api-server/src/test/typescript/integration/remote-plugin-imports.test.ts`,
    `./packages/cactus-test-cmd-api-server/src/test/typescript/integration/plugin-import-with-npm-install.test.ts`,
    `./packages/cactus-test-cmd-api-server/src/test/typescript/integration/plugin-import-with-npm-install-version-selection.test.ts`,
    `./packages/cactus-test-cmd-api-server/src/test/typescript/integration/runtime-plugin-imports.test.ts`,
    `./packages/cactus-cmd-socketio-server/`,
    `./packages/cactus-plugin-ledger-connector-corda/src/test/typescript/integration/openapi/openapi-validation.test.ts`,
    `./packages/cactus-plugin-ledger-connector-corda/src/test/typescript/integration/deploy-cordapp-jars-to-nodes-v4.7.test.ts`,
    `./packages/cactus-plugin-ledger-connector-corda/src/test/typescript/integration/jvm-kotlin-spring-server-v4.7.test.ts`,
    `./packages/cactus-plugin-ledger-connector-corda/src/test/typescript/integration/jvm-kotlin-spring-server-v4.8.test.ts`,
    `./packages/cactus-plugin-ledger-connector-corda/src/test/typescript/integration/deploy-cordapp-jars-to-nodes.test.ts`,
    `./packages/cactus-plugin-ledger-connector-corda/src/test/typescript/integration/deploy-cordapp-jars-to-nodes-v4.8.test.ts`,
    `./packages/cactus-plugin-ledger-connector-corda/src/test/typescript/integration/jvm-kotlin-spring-server.test.ts`,
    `./packages/cactus-plugin-keychain-google-sm/src/test/typescript/integration/plugin-factory-keychain.test.ts`,
    `./packages/cactus-plugin-keychain-google-sm/src/test/typescript/integration/plugin-keychain-google-sm.test.ts`,
    `./packages/cactus-test-plugin-consortium-manual/src/test/typescript/integration/plugin-consortium-manual/get-consortium-jws-endpoint.test.ts`,
    `./packages/cactus-test-plugin-consortium-manual/src/test/typescript/integration/plugin-consortium-manual/openapi/openapi-validation.test.ts`,
    `./packages/cactus-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-ledger-connector-besu/deploy-contract/private-deploy-contract-from-json-cactus.test.ts`,
    `./packages/cactus-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-ledger-connector-besu/deploy-contract/v21-invoke-contract.test.ts`,
    `./packages/cactus-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-ledger-connector-besu/deploy-contract/v21-get-block.test.ts`,
    `./packages/cactus-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-ledger-connector-besu/deploy-contract/deploy-contract-from-json.test.ts`,
    `./packages/cactus-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-ledger-connector-besu/deploy-contract/v21-get-balance.test.ts`,
    `./packages/cactus-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-ledger-connector-besu/deploy-contract/v21-deploy-contract-from-json.test.ts`,
    `./packages/cactus-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-ledger-connector-besu/deploy-contract/v21-get-record-locator.test.ts`,
    `./packages/cactus-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-ledger-connector-besu/deploy-contract/get-balance.test.ts`,
    `./packages/cactus-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-ledger-connector-besu/deploy-contract/invoke-contract.test.ts`,
    `./packages/cactus-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-ledger-connector-besu/deploy-contract/private-deploy-contract-from-json-web3-eea.test.ts`,
    `./packages/cactus-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-ledger-connector-besu/deploy-contract/v21-get-past-logs.test.ts`,
    `./packages/cactus-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-ledger-connector-besu/deploy-contract/get-past-logs.test.ts`,
    `./packages/cactus-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-ledger-connector-besu/deploy-contract/get-record-locator.test.ts`,
    `./packages/cactus-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-ledger-connector-besu/deploy-contract/get-block.test.ts`,
    `./packages/cactus-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-ledger-connector-besu/v21-besu-get-transaction.test.ts`,
    `./packages/cactus-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-ledger-connector-besu/openapi/openapi-validation.test.ts`,
    `./packages/cactus-plugin-ledger-connector-iroha/src/test/typescript/integration/iroha-iroha-transfer-example.test.ts`,
    `./packages/cactus-plugin-ledger-connector-iroha/src/test/typescript/integration/run-transaction-endpoint-v1.test.ts`,
    `./packages/cactus-plugin-ledger-connector-iroha/src/test/typescript/integration/openapi/openapi-validation.test.ts`,
    `./packages/cactus-plugin-ledger-connector-iroha/src/test/typescript/unit/iroha-test-ledger-parameters.test.ts`,
    `./packages/cactus-plugin-ledger-connector-iroha/src/test/typescript/unit/postgres-test-container-parameters.test.ts`,
    `./packages/cactus-plugin-ledger-connector-xdai/src/test/typescript/integration/deploy-contract-from-json-xdai-json-object.test.ts`,
    `./packages/cactus-plugin-ledger-connector-xdai/src/test/typescript/integration/invoke-contract-xdai-json-object.test.ts`,
    `./packages/cactus-plugin-ledger-connector-xdai/src/test/typescript/integration/openapi/openapi-validation.test.ts`,
    `./packages/cactus-plugin-ledger-connector-xdai/src/test/typescript/integration/openapi/openapi-validation-no-keychain.test.ts`,
    `./packages/cactus-common/src/test/typescript/unit/servers.test.ts`,
    `./packages/cactus-common/src/test/typescript/unit/key-converter.test.ts`,
    `./packages/cactus-common/src/test/typescript/unit/logging/logger.test.ts`,
    `./packages/cactus-common/src/test/typescript/unit/js-object-signer.test.ts`,
    `./packages/cactus-plugin-keychain-aws-sm/src/test/typescript/integration/plugin-keychain-aws-sm.test.ts`,
    `./packages/cactus-test-api-client/src/test/typescript/integration/api-client-routing-node-to-node.test.ts`,
    `./packages/cactus-plugin-keychain-memory/src/test/typescript/unit/plugin-keychain-memory.test.ts`,
    `./packages/cactus-api-client/src/test/typescript/integration/default-consortium-provider.test.ts`,
    `./packages/cactus-core/src/test/typescript/unit/plugin-registry.test.ts`,
    `./packages/cactus-test-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-validator-besu/get-balance-endpoint.test.ts`,
    `./packages/cactus-test-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-validator-besu/v21-get-past-logs-endpoint.test.ts`,
    `./packages/cactus-test-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-validator-besu/v21-get-balance-endpoint.test.ts`,
    `./packages/cactus-test-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-validator-besu/v21-get-block-endpoint.test.ts`,
    `./packages/cactus-test-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-validator-besu/get-past-logs-endpoint.test.ts`,
    `./packages/cactus-test-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-validator-besu/get-transaction-endpoint.test.ts`,
    `./packages/cactus-test-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-validator-besu/get-block-endpoint.test.ts`,
    `./packages/cactus-test-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-validator-besu/v21-get-transaction-endpoint.test.ts`,
    `./packages/cactus-test-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-validator-besu/v21-sign-transaction-endpoint.test.ts`,
    `./packages/cactus-test-plugin-ledger-connector-besu/src/test/typescript/integration/plugin-validator-besu/sign-transaction-endpoint.test.ts`,
    `./packages/cactus-plugin-ledger-connector-fabric/src/test/typescript/integration/identity-client.test.ts`,
    `./packages/cactus-plugin-ledger-connector-fabric/src/test/typescript/integration/fabric-v1-4-x/run-transaction-endpoint-v1.test.ts`,
    `./packages/cactus-plugin-ledger-connector-fabric/src/test/typescript/integration/fabric-v1-4-x/deploy-cc-from-golang-source.test.ts`,
    `./packages/cactus-plugin-ledger-connector-fabric/src/test/typescript/integration/fabric-v2-2-x/deploy-cc-from-javascript-source.test.ts`,
    `./packages/cactus-plugin-ledger-connector-fabric/src/test/typescript/integration/fabric-v2-2-x/run-transaction-with-identities.test.ts`,
    `./packages/cactus-plugin-ledger-connector-fabric/src/test/typescript/integration/fabric-v2-2-x/deploy-cc-from-golang-source.test.ts`,
    `./packages/cactus-plugin-ledger-connector-fabric/src/test/typescript/integration/fabric-v2-2-x/deploy-cc-from-typescript-source.test.ts`,
    `./packages/cactus-plugin-ledger-connector-fabric/src/test/typescript/integration/openapi/openapi-validation-go.test.ts`,
    `./packages/cactus-plugin-ledger-connector-fabric/src/test/typescript/integration/openapi/openapi-validation.test.ts`,
    `./packages/cactus-plugin-ledger-connector-fabric/src/test/typescript/unit/identity-internal-crypto-utils.test.ts`,
    `./packages/cactus-plugin-keychain-vault/src/test/typescript/integration/cactus-keychain-vault-server.test.ts`,
    `./packages/cactus-plugin-keychain-vault/src/test/typescript/integration/plugin-keychain-vault.test.ts`,
    `./packages/cactus-plugin-keychain-vault/src/test/typescript/integration/openapi/openapi-validation.test.ts`,
    `./packages/cactus-test-tooling/src/test/typescript/integration/postgres/postgres-test-container/constructor-validates-options.test.ts`,
    `./packages/cactus-test-tooling/src/test/typescript/integration/quorum/quorum-test-ledger/constructor-validates-options.test.ts`,
    `./packages/cactus-test-tooling/src/test/typescript/integration/fabric/fabric-test-ledger-v1/constructor-validates-options.test.ts`,
    `./packages/cactus-test-tooling/src/test/typescript/integration/substrate/substrate-test-ledger-constructor.test.ts`,
    `./packages/cactus-test-tooling/src/test/typescript/integration/iroha/iroha-test-ledger/constructor-validates-options.test.ts`,
    `./packages/cactus-test-tooling/src/test/typescript/integration/rustc-container/rustc-container-target-bundler.test.ts`,
    `./packages/cactus-test-tooling/src/test/typescript/integration/common/containers.test.ts`,
    `./packages/cactus-test-tooling/src/test/typescript/integration/besu/besu-test-ledger/constructor-validates-options.test.ts`,
    `./packages/cactus-test-plugin-htlc-eth-besu/src/test/typescript/integration/plugin-htlc-eth-besu/get-status-endpoint-invalid.test.ts`,
    `./packages/cactus-test-plugin-htlc-eth-besu/src/test/typescript/integration/plugin-htlc-eth-besu/new-contract-endpoint.test.ts`,
    `./packages/cactus-test-plugin-htlc-eth-besu/src/test/typescript/integration/plugin-htlc-eth-besu/get-status-endpoint.test.ts`,
    `./packages/cactus-test-plugin-htlc-eth-besu/src/test/typescript/integration/plugin-htlc-eth-besu/initialize-endpoint.test.ts`,
    `./packages/cactus-test-plugin-htlc-eth-besu/src/test/typescript/integration/plugin-htlc-eth-besu/new-contract-endpoint-invalid.test.ts`,
    `./packages/cactus-test-plugin-htlc-eth-besu/src/test/typescript/integration/plugin-htlc-eth-besu/withdraw-endpoint.test.ts`,
    `./packages/cactus-test-plugin-htlc-eth-besu/src/test/typescript/integration/plugin-htlc-eth-besu/get-single-status-endpoint.test.ts`,
    `./packages/cactus-test-plugin-htlc-eth-besu/src/test/typescript/integration/plugin-htlc-eth-besu/initialize-endpoint-invalid.test.ts`,
    `./packages/cactus-test-plugin-htlc-eth-besu/src/test/typescript/integration/plugin-htlc-eth-besu/get-single-status-endpoint-invalid.test.ts`,
    `./packages/cactus-test-plugin-htlc-eth-besu/src/test/typescript/integration/plugin-htlc-eth-besu/refund-endpoint-invalid.test.ts`,
    `./packages/cactus-test-plugin-htlc-eth-besu/src/test/typescript/integration/plugin-htlc-eth-besu/withdraw-endpoint-invalid.test.ts`,
    `./packages/cactus-test-plugin-htlc-eth-besu/src/test/typescript/integration/plugin-htlc-eth-besu/openapi/openapi-validation.test.ts`,
    `./packages/cactus-test-plugin-htlc-eth-besu/src/test/typescript/integration/plugin-htlc-eth-besu/refund-endpoint.test.ts`,
    `./packages/cactus-cmd-api-server/src/test/typescript/benchmark/artillery-api-benchmark.test.ts`,
    `./packages/cactus-cmd-api-server/src/test/typescript/integration/jwt-socketio-endpoint-authorization.test.ts`,
    `./packages/cactus-cmd-api-server/src/test/typescript/integration/jwt-endpoint-authz-scope-enforcement.test.ts`,
    `./packages/cactus-cmd-api-server/src/test/typescript/integration/remote-plugin-imports.test.ts`,
    `./packages/cactus-cmd-api-server/src/test/typescript/integration/jwt-unprotected-endpoint-authz.test.ts`,
    `./packages/cactus-cmd-api-server/src/test/typescript/integration/jwt-endpoint-authorization.test.ts`,
    `./packages/cactus-cmd-api-server/src/test/typescript/integration/jwt-unprotected-endpoint-authz-ops-confirm.test.ts`,
    `./packages/cactus-cmd-api-server/src/test/typescript/integration/plugin-import-from-github.test.ts`,
    `./packages/cactus-cmd-api-server/src/test/typescript/integration/plugin-import-without-install.test.ts`,
    `./packages/cactus-cmd-api-server/src/test/typescript/unit/plugins/install-basic-plugin-keychain-memory.test.ts`,
    `./packages/cactus-cmd-api-server/src/test/typescript/unit/plugins/install-basic-plugin-ledger-connector-quorum-0-7-0.test.ts`,
    `./packages/cactus-cmd-api-server/src/test/typescript/unit/plugins/install-basic-plugin-ledger-connector-fabric-0-7-0.test.ts`,
    `./packages/cactus-cmd-api-server/src/test/typescript/unit/plugins/install-basic-plugin-consortium-manual.test.ts`,
    `./packages/cactus-cmd-api-server/src/test/typescript/unit/config/config-service-example-config-validity.test.ts`,
    `./packages/cactus-cmd-api-server/src/test/typescript/unit/config/self-signed-certificate-generator/generates-working-certificates.test.ts`,
    `./packages/cactus-cmd-api-server/src/test/typescript/unit/config/self-signed-certificate-generator/certificates-work-for-mutual-tls.test.ts`,
    `./packages/cactus-cmd-api-server/src/test/typescript/unit/grpc-proto-gen-ts-client-healthcheck.test.ts`,
    `./packages/cactus-cmd-api-server/src/test/typescript/unit/grpc-js-proto-loader-client-healthcheck.test.ts`,
    `./packages/cactus-cmd-api-server/src/test/typescript/unit/grpc-proto-gen-ts-client-m-tls-enabled.test.ts`,
    `./extensions/cactus-plugin-object-store-ipfs/src/test/typescript/integration/plugin-object-store-ipfs.test.ts`,
    `./extensions/cactus-plugin-object-store-ipfs/src/test/typescript/unit/plugin-object-store-ipfs.test.ts`,
    `./examples/cactus-example-carbon-accounting-backend/src/test/typescript/integration/admin-enroll-v1-endpoint.test.ts`,
    `./examples/cactus-example-supply-chain-backend/src/test/typescript/integration/supply-chain-backend-api-calls.test.ts`,
    `./examples/cactus-example-supply-chain-backend/src/test/typescript/integration/supply-chain-cli-via-npm-script.test.ts`,
  ],
};
