// @generated by protoc-gen-connect-es v1.3.0 with parameter "target=ts,js_import_style=module"
// @generated from file cacti/satp/v02/stage_1.proto (package cacti.satp.v02, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { TransferCommenceRequestMessage, TransferCommenceResponseMessage, TransferProposalRequestMessage, TransferProposalResponseMessage } from "./stage_1_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service cacti.satp.v02.SatpStage1Service
 */
export const SatpStage1Service = {
  typeName: "cacti.satp.v02.SatpStage1Service",
  methods: {
    /**
     * @generated from rpc cacti.satp.v02.SatpStage1Service.TransferProposal
     */
    transferProposal: {
      name: "TransferProposal",
      I: TransferProposalRequestMessage,
      O: TransferProposalResponseMessage,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc cacti.satp.v02.SatpStage1Service.TransferCommence
     */
    transferCommence: {
      name: "TransferCommence",
      I: TransferCommenceRequestMessage,
      O: TransferCommenceResponseMessage,
      kind: MethodKind.Unary,
    },
  }
} as const;

