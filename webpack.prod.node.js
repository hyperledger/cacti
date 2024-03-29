const packageDir = process.cwd();
const pkg = require(`${packageDir}/package.json`);
const TerserPlugin = require("terser-webpack-plugin");

const packageNameNoScope = pkg.name.substring(pkg.name.lastIndexOf("/") + 1);
const libraryName = `${packageNameNoScope}`;

/** @type {import("webpack").Configuration} */
module.exports = {
  entry: {
    [pkg.mainMinified]: `${packageDir}/src/main/typescript/index.ts`,
  },
  target: "node",
  mode: "production",
  devtool: "source-map",
  module: {
    rules: [
      {
        test: /\.(js|ts)$/,
        use: ["shebang-loader"],
      },
      {
        test: /\.ts$/,
        exclude: /node_modules/,
        use: [
          {
            loader: "ts-loader",
            options: {
              transpileOnly: true,
              configFile: "tsconfig.json",
            },
          },
        ],
      },
      {
        test: /\.(js|ts)$/,
        enforce: "pre",
        use: [
          {
            loader: "source-map-loader",
          },
        ],
      },
    ],
  },
  resolve: {
    extensions: [".ts", ".js"],
  },
  plugins: [],
  optimization: {
    minimizer: [new TerserPlugin()],
  },
  output: {
    filename: "[name]",
    path: packageDir,
    libraryTarget: "umd",
    library: libraryName,
    umdNamedDefine: true,
    globalObject: "this",
  },
  externals: {
    "swarm-js": "swarm-js",
    "node-ssh": "node-ssh",
    npm: "npm",
    "fabric-client": "fabric-client",
    "fabric-ca-client": "fabric-ca-client",
    "@azure/identity": "@azure/identity",
  },
};
