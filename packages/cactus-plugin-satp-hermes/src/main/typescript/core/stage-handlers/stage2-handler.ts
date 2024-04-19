import { ConnectRouter, HandlerContext } from "@connectrpc/connect";
import { SatpStage2Service } from "../../generated/proto/cacti/satp/v02/stage_2_connect";
import { LockAssertionRequestMessage } from "../../generated/proto/cacti/satp/v02/stage_2_pb";
import { SATPGateway } from "../../gateway-refactor";
import { Stage2ServerService } from "../stage-services/server/stage2-server-service";
import { TimestampType, saveTimestamp } from "../session-utils";
import { MessageType } from "../../generated/proto/cacti/satp/v02/common/message_pb";

export const Stage2Handler = (gateway: SATPGateway, service: Stage2ServerService) =>
  (router: ConnectRouter) =>
    router.service(SatpStage2Service, {
      async lockAssertion(
        req: LockAssertionRequestMessage,
        context: HandlerContext,
      ) {
        console.log("Received LockAssertionRequest", req, context);
        const recvTimestamp: string = Date.now().toString();

        const sessionData = await service.checkLockAssertionRequestMessage(
          req,
          gateway,
        );

        saveTimestamp(
          sessionData,
          MessageType.LOCK_ASSERT,
          TimestampType.RECEIVED,
          recvTimestamp,
        );

        const message = await service.lockAssertionResponse(req, gateway);

        if (!message) {
          throw new Error("No message returned from lockAssertionResponse");
        }

        saveTimestamp(
          sessionData,
          MessageType.ASSERTION_RECEIPT,
          TimestampType.PROCESSED,
        );

        console.log("Returning response", message);

        return message;
      },
    });
