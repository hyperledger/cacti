import { TransactionResponse } from "../../../types/blockchain-interaction";

export abstract class NetworkBridge {
  network!: string;

  public networkName(): string {
    return this.network;
  }

  public abstract runTransaction(
    methodName: string,
    params: string[],
  ): Promise<TransactionResponse>;

  public abstract getReceipt(transactionId: string): Promise<string>;
}
