/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 3.15.6
 * source: common/ack.proto
 * git: https://github.com/thesayyn/protoc-gen-ts
 * buymeacoffee: https://www.buymeacoffee.com/thesayyn
 *  */
import * as pb_1 from "google-protobuf";
export namespace common.ack {
    export class Ack extends pb_1.Message {
        constructor(data?: any[] | {
            status?: Ack.STATUS;
            request_id?: string;
            message?: string;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], []);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("status" in data && data.status != undefined) {
                    this.status = data.status;
                }
                if ("request_id" in data && data.request_id != undefined) {
                    this.request_id = data.request_id;
                }
                if ("message" in data && data.message != undefined) {
                    this.message = data.message;
                }
            }
        }
        get status() {
            return pb_1.Message.getField(this, 2) as Ack.STATUS;
        }
        set status(value: Ack.STATUS) {
            pb_1.Message.setField(this, 2, value);
        }
        get request_id() {
            return pb_1.Message.getField(this, 3) as string;
        }
        set request_id(value: string) {
            pb_1.Message.setField(this, 3, value);
        }
        get message() {
            return pb_1.Message.getField(this, 4) as string;
        }
        set message(value: string) {
            pb_1.Message.setField(this, 4, value);
        }
        toObject() {
            const data: {
                status?: Ack.STATUS;
                request_id?: string;
                message?: string;
            } = {};
            if (this.status != null) {
                data.status = this.status;
            }
            if (this.request_id != null) {
                data.request_id = this.request_id;
            }
            if (this.message != null) {
                data.message = this.message;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.status !== undefined)
                writer.writeEnum(2, this.status);
            if (typeof this.request_id === "string" && this.request_id.length)
                writer.writeString(3, this.request_id);
            if (typeof this.message === "string" && this.message.length)
                writer.writeString(4, this.message);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): Ack {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new Ack();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 2:
                        message.status = reader.readEnum();
                        break;
                    case 3:
                        message.request_id = reader.readString();
                        break;
                    case 4:
                        message.message = reader.readString();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): Ack {
            return Ack.deserialize(bytes);
        }
    }
    export namespace Ack {
        export enum STATUS {
            OK = 0,
            ERROR = 1
        }
    }
}
