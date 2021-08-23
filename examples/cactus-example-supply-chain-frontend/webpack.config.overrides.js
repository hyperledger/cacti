/** @type {import("webpack").Configuration} */
module.exports = {
  externals: {
    "@fidm/x509": "@fidm/x509",
    "@grpc/grpc-js": "@grpc/grpc-js",
    "@grpc/proto-loader": "@grpc/proto-loader",
    busboy: "busboy",
    child_process: "child_process",
    "cross-spawn": "cross-spawn",
    decompress: "decompress",
    express: "express",
    "fabric-ca-client": "fabric-ca-client",
    "fabric-common": "fabric-common",
    "fabric-network": "fabric-network",
    "fd-slicer": "fd-slicer",
    fs: "fs",
    "node-ssh": "node-ssh",
    "prom-client": "prom-client",
    yargs: "yargs",
    winston: "winston",
  },
  resolve: {
    fallback: {
      constants: require.resolve("constants-browserify"),
      crypto: require.resolve("crypto-browserify"),
      buffer: require.resolve("buffer/"),
      http: require.resolve("stream-http"),
      https: require.resolve("https-browserify"),
      os: require.resolve("os-browserify/browser"),
      path: require.resolve("path-browserify"),
      stream: require.resolve("stream-browserify"),
      zlib: require.resolve("browserify-zlib"),
    },
  },
};
