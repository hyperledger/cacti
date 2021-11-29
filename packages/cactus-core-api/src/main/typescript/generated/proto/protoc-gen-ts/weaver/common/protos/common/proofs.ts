/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 3.15.6
 * source: common/proofs.proto
 * git: https://github.com/thesayyn/protoc-gen-ts
 * buymeacoffee: https://www.buymeacoffee.com/thesayyn
 *  */
import * as pb_1 from "google-protobuf";
export namespace common.proofs {
    export class Proof extends pb_1.Message {
        constructor(data?: any[] | {
            signature?: string;
            certificate?: string;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], []);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("signature" in data && data.signature != undefined) {
                    this.signature = data.signature;
                }
                if ("certificate" in data && data.certificate != undefined) {
                    this.certificate = data.certificate;
                }
            }
        }
        get signature() {
            return pb_1.Message.getField(this, 1) as string;
        }
        set signature(value: string) {
            pb_1.Message.setField(this, 1, value);
        }
        get certificate() {
            return pb_1.Message.getField(this, 2) as string;
        }
        set certificate(value: string) {
            pb_1.Message.setField(this, 2, value);
        }
        static fromObject(data: {
            signature?: string;
            certificate?: string;
        }) {
            const message = new Proof({});
            if (data.signature != null) {
                message.signature = data.signature;
            }
            if (data.certificate != null) {
                message.certificate = data.certificate;
            }
            return message;
        }
        toObject() {
            const data: {
                signature?: string;
                certificate?: string;
            } = {};
            if (this.signature != null) {
                data.signature = this.signature;
            }
            if (this.certificate != null) {
                data.certificate = this.certificate;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (typeof this.signature === "string" && this.signature.length)
                writer.writeString(1, this.signature);
            if (typeof this.certificate === "string" && this.certificate.length)
                writer.writeString(2, this.certificate);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): Proof {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new Proof();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.signature = reader.readString();
                        break;
                    case 2:
                        message.certificate = reader.readString();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): Proof {
            return Proof.deserialize(bytes);
        }
    }
    export class Proofs extends pb_1.Message {
        constructor(data?: any[] | {
            proofs?: Proof[];
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [1], []);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("proofs" in data && data.proofs != undefined) {
                    this.proofs = data.proofs;
                }
            }
        }
        get proofs() {
            return pb_1.Message.getRepeatedWrapperField(this, Proof, 1) as Proof[];
        }
        set proofs(value: Proof[]) {
            pb_1.Message.setRepeatedWrapperField(this, 1, value);
        }
        static fromObject(data: {
            proofs?: ReturnType<typeof Proof.prototype.toObject>[];
        }) {
            const message = new Proofs({});
            if (data.proofs != null) {
                message.proofs = data.proofs.map(item => Proof.fromObject(item));
            }
            return message;
        }
        toObject() {
            const data: {
                proofs?: ReturnType<typeof Proof.prototype.toObject>[];
            } = {};
            if (this.proofs != null) {
                data.proofs = this.proofs.map((item: Proof) => item.toObject());
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.proofs !== undefined)
                writer.writeRepeatedMessage(1, this.proofs, (item: Proof) => item.serialize(writer));
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): Proofs {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new Proofs();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        reader.readMessage(message.proofs, () => pb_1.Message.addToRepeatedWrapperField(message, 1, Proof.deserialize(reader), Proof));
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): Proofs {
            return Proofs.deserialize(bytes);
        }
    }
}
