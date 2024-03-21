import { Proof } from "./proof";

export class TransactionProof {
  private creator: Proof;
  private endorsers?: Proof[] = [];

  constructor(creator: Proof) {
    this.creator = creator;
  }

  public addEndorser(endorser: Proof) {
    this.endorsers?.push(endorser);
  }

  public printEndorsement(): string {
    return "Endorsement: \n " + this.creator + " \n " + this.endorsers;
  }

  public getEndorsementJson(): string {
    const proof = {
      creator: this.creator,
      endorsers: this.endorsers,
    };
    // JSON.stringify(Array.from(endors.entries()));
    return JSON.stringify(proof);
  }

  public getCreator(): Proof {
    return this.creator;
  }

  public getEndorsements(): Proof[] | undefined {
    return this.endorsers;
  }
}
