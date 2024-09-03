#!/bin/bash

generate_keys() {
    # Check if we need to generate new keys
    if [ "$SATP_PRIVATE_KEY" = "default_private_key" ] || [ "$SATP_PUBLIC_KEY" = "default_public_key" ]; then
        echo "Generating new key pair..."
        # Using OpenSSL to generate a 256-bit private key and derive public key
        PRIVATE_KEY=$(openssl ecparam -name secp256k1 -genkey -noout | openssl ec -text -noout | grep priv -A 3 | tail -n +2 | tr -d '\n[:space:]:' | sed 's/^00//')
        PUBLIC_KEY=$(openssl ecparam -name secp256k1 -genkey -noout | openssl ec -text -noout | grep pub -A 5 | tail -n +2 | tr -d '\n[:space:]:' | sed 's/^04//')

        export SATP_PRIVATE_KEY=$PRIVATE_KEY
        export SATP_PUBLIC_KEY=$PUBLIC_KEY
    else
        echo "Using existing key pair from environment variables"
    fi

    echo "SATP_PRIVATE_KEY: $SATP_PRIVATE_KEY"
    echo "SATP_PUBLIC_KEY: $SATP_PUBLIC_KEY"
}

run_gateway() {
    export NVM_DIR="$HOME/.nvm"
    [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh"
    nvm use 18.18.2

    echo "Starting SATP Hermes Gateway CLI..."
    node --max-old-space-size=2048 -r ts-node/register /usr/src/app/cactus/packages/cactus-plugin-satp-hermes/src/main/typescript/plugin-satp-hermes-gateway-cli.ts
    echo "SATP Hermes Gateway CLI execution completed."
}

main() {
    generate_keys
    run_gateway
}

main