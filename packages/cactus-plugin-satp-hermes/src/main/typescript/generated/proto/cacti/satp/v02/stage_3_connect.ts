// @generated by protoc-gen-connect-es v1.3.0 with parameter "target=ts,js_import_style=module"
// @generated from file cacti/satp/v02/stage_3.proto (package cacti.satp.v02, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CommitFinalAcknowledgementReceiptResponseMessage, CommitFinalAssertionRequestMessage, CommitPreparationRequestMessage, CommitReadyResponseMessage, TransferCompleteRequestMessage, TransferCompleteResponseMessage } from "./stage_3_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service cacti.satp.v02.SatpStage3Service
 */
export const SatpStage3Service = {
  typeName: "cacti.satp.v02.SatpStage3Service",
  methods: {
    /**
     * @generated from rpc cacti.satp.v02.SatpStage3Service.CommitPreparation
     */
    commitPreparation: {
      name: "CommitPreparation",
      I: CommitPreparationRequestMessage,
      O: CommitReadyResponseMessage,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc cacti.satp.v02.SatpStage3Service.CommitFinalAssertion
     */
    commitFinalAssertion: {
      name: "CommitFinalAssertion",
      I: CommitFinalAssertionRequestMessage,
      O: CommitFinalAcknowledgementReceiptResponseMessage,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc cacti.satp.v02.SatpStage3Service.TransferComplete
     */
    transferComplete: {
      name: "TransferComplete",
      I: TransferCompleteRequestMessage,
      O: TransferCompleteResponseMessage,
      kind: MethodKind.Unary,
    },
  }
} as const;

