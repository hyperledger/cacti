export class Proof {
  private creator: string;
  private mspid?: string;
  private signature?: string;

  constructor(settings: {
    creator: string;
    mspid?: string;
    signature?: string;
  }) {
    this.creator = settings.creator;
    this.mspid = settings.mspid ? settings.mspid : "undefined";
    this.signature = settings.signature ? settings.signature : "undefined";
  }

  public printEndorsement(): string {
    return (
      "Endorsement: \n " +
      this.creator +
      " \n " +
      this.mspid +
      " \n " +
      this.signature
    );
  }

  public getEndorsementJson(): string {
    const proof = {
      creator: this.creator,
      mspid: this.mspid,
      signature: this.signature,
    };
    // JSON.stringify(Array.from(endors.entries()));
    return JSON.stringify(proof);
  }

  public getCreator(): string {
    return this.creator;
  }

  public getMspid(): string | undefined {
    return this.mspid;
  }

  public getSignature(): string | undefined {
    return this.signature;
  }
}
