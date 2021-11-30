/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 3.15.6
 * source: common/asset_locks.proto
 * git: https://github.com/thesayyn/protoc-gen-ts
 * buymeacoffee: https://www.buymeacoffee.com/thesayyn
 *  */
import * as pb_1 from "google-protobuf";
export namespace common.asset_locks {
    export enum LockMechanism {
        HTLC = 0
    }
    export class AssetLock extends pb_1.Message {
        constructor(data?: any[] | {
            lockMechanism?: LockMechanism;
            lockInfo?: Uint8Array;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], []);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("lockMechanism" in data && data.lockMechanism != undefined) {
                    this.lockMechanism = data.lockMechanism;
                }
                if ("lockInfo" in data && data.lockInfo != undefined) {
                    this.lockInfo = data.lockInfo;
                }
            }
        }
        get lockMechanism() {
            return pb_1.Message.getField(this, 1) as LockMechanism;
        }
        set lockMechanism(value: LockMechanism) {
            pb_1.Message.setField(this, 1, value);
        }
        get lockInfo() {
            return pb_1.Message.getField(this, 2) as Uint8Array;
        }
        set lockInfo(value: Uint8Array) {
            pb_1.Message.setField(this, 2, value);
        }
        static fromObject(data: {
            lockMechanism?: LockMechanism;
            lockInfo?: Uint8Array;
        }) {
            const message = new AssetLock({});
            if (data.lockMechanism != null) {
                message.lockMechanism = data.lockMechanism;
            }
            if (data.lockInfo != null) {
                message.lockInfo = data.lockInfo;
            }
            return message;
        }
        toObject() {
            const data: {
                lockMechanism?: LockMechanism;
                lockInfo?: Uint8Array;
            } = {};
            if (this.lockMechanism != null) {
                data.lockMechanism = this.lockMechanism;
            }
            if (this.lockInfo != null) {
                data.lockInfo = this.lockInfo;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.lockMechanism !== undefined)
                writer.writeEnum(1, this.lockMechanism);
            if (this.lockInfo !== undefined)
                writer.writeBytes(2, this.lockInfo);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): AssetLock {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new AssetLock();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.lockMechanism = reader.readEnum();
                        break;
                    case 2:
                        message.lockInfo = reader.readBytes();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): AssetLock {
            return AssetLock.deserialize(bytes);
        }
    }
    export class AssetClaim extends pb_1.Message {
        constructor(data?: any[] | {
            lockMechanism?: LockMechanism;
            claimInfo?: Uint8Array;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], []);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("lockMechanism" in data && data.lockMechanism != undefined) {
                    this.lockMechanism = data.lockMechanism;
                }
                if ("claimInfo" in data && data.claimInfo != undefined) {
                    this.claimInfo = data.claimInfo;
                }
            }
        }
        get lockMechanism() {
            return pb_1.Message.getField(this, 1) as LockMechanism;
        }
        set lockMechanism(value: LockMechanism) {
            pb_1.Message.setField(this, 1, value);
        }
        get claimInfo() {
            return pb_1.Message.getField(this, 2) as Uint8Array;
        }
        set claimInfo(value: Uint8Array) {
            pb_1.Message.setField(this, 2, value);
        }
        static fromObject(data: {
            lockMechanism?: LockMechanism;
            claimInfo?: Uint8Array;
        }) {
            const message = new AssetClaim({});
            if (data.lockMechanism != null) {
                message.lockMechanism = data.lockMechanism;
            }
            if (data.claimInfo != null) {
                message.claimInfo = data.claimInfo;
            }
            return message;
        }
        toObject() {
            const data: {
                lockMechanism?: LockMechanism;
                claimInfo?: Uint8Array;
            } = {};
            if (this.lockMechanism != null) {
                data.lockMechanism = this.lockMechanism;
            }
            if (this.claimInfo != null) {
                data.claimInfo = this.claimInfo;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.lockMechanism !== undefined)
                writer.writeEnum(1, this.lockMechanism);
            if (this.claimInfo !== undefined)
                writer.writeBytes(2, this.claimInfo);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): AssetClaim {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new AssetClaim();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.lockMechanism = reader.readEnum();
                        break;
                    case 2:
                        message.claimInfo = reader.readBytes();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): AssetClaim {
            return AssetClaim.deserialize(bytes);
        }
    }
    export class AssetLockHTLC extends pb_1.Message {
        constructor(data?: any[] | {
            hashBase64?: Uint8Array;
            expiryTimeSecs?: number;
            timeSpec?: AssetLockHTLC.TimeSpec;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], []);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("hashBase64" in data && data.hashBase64 != undefined) {
                    this.hashBase64 = data.hashBase64;
                }
                if ("expiryTimeSecs" in data && data.expiryTimeSecs != undefined) {
                    this.expiryTimeSecs = data.expiryTimeSecs;
                }
                if ("timeSpec" in data && data.timeSpec != undefined) {
                    this.timeSpec = data.timeSpec;
                }
            }
        }
        get hashBase64() {
            return pb_1.Message.getField(this, 1) as Uint8Array;
        }
        set hashBase64(value: Uint8Array) {
            pb_1.Message.setField(this, 1, value);
        }
        get expiryTimeSecs() {
            return pb_1.Message.getField(this, 2) as number;
        }
        set expiryTimeSecs(value: number) {
            pb_1.Message.setField(this, 2, value);
        }
        get timeSpec() {
            return pb_1.Message.getField(this, 3) as AssetLockHTLC.TimeSpec;
        }
        set timeSpec(value: AssetLockHTLC.TimeSpec) {
            pb_1.Message.setField(this, 3, value);
        }
        static fromObject(data: {
            hashBase64?: Uint8Array;
            expiryTimeSecs?: number;
            timeSpec?: AssetLockHTLC.TimeSpec;
        }) {
            const message = new AssetLockHTLC({});
            if (data.hashBase64 != null) {
                message.hashBase64 = data.hashBase64;
            }
            if (data.expiryTimeSecs != null) {
                message.expiryTimeSecs = data.expiryTimeSecs;
            }
            if (data.timeSpec != null) {
                message.timeSpec = data.timeSpec;
            }
            return message;
        }
        toObject() {
            const data: {
                hashBase64?: Uint8Array;
                expiryTimeSecs?: number;
                timeSpec?: AssetLockHTLC.TimeSpec;
            } = {};
            if (this.hashBase64 != null) {
                data.hashBase64 = this.hashBase64;
            }
            if (this.expiryTimeSecs != null) {
                data.expiryTimeSecs = this.expiryTimeSecs;
            }
            if (this.timeSpec != null) {
                data.timeSpec = this.timeSpec;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.hashBase64 !== undefined)
                writer.writeBytes(1, this.hashBase64);
            if (this.expiryTimeSecs !== undefined)
                writer.writeUint64(2, this.expiryTimeSecs);
            if (this.timeSpec !== undefined)
                writer.writeEnum(3, this.timeSpec);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): AssetLockHTLC {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new AssetLockHTLC();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.hashBase64 = reader.readBytes();
                        break;
                    case 2:
                        message.expiryTimeSecs = reader.readUint64();
                        break;
                    case 3:
                        message.timeSpec = reader.readEnum();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): AssetLockHTLC {
            return AssetLockHTLC.deserialize(bytes);
        }
    }
    export namespace AssetLockHTLC {
        export enum TimeSpec {
            EPOCH = 0,
            DURATION = 1
        }
    }
    export class AssetClaimHTLC extends pb_1.Message {
        constructor(data?: any[] | {
            hashPreimageBase64?: Uint8Array;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], []);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("hashPreimageBase64" in data && data.hashPreimageBase64 != undefined) {
                    this.hashPreimageBase64 = data.hashPreimageBase64;
                }
            }
        }
        get hashPreimageBase64() {
            return pb_1.Message.getField(this, 1) as Uint8Array;
        }
        set hashPreimageBase64(value: Uint8Array) {
            pb_1.Message.setField(this, 1, value);
        }
        static fromObject(data: {
            hashPreimageBase64?: Uint8Array;
        }) {
            const message = new AssetClaimHTLC({});
            if (data.hashPreimageBase64 != null) {
                message.hashPreimageBase64 = data.hashPreimageBase64;
            }
            return message;
        }
        toObject() {
            const data: {
                hashPreimageBase64?: Uint8Array;
            } = {};
            if (this.hashPreimageBase64 != null) {
                data.hashPreimageBase64 = this.hashPreimageBase64;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.hashPreimageBase64 !== undefined)
                writer.writeBytes(1, this.hashPreimageBase64);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): AssetClaimHTLC {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new AssetClaimHTLC();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.hashPreimageBase64 = reader.readBytes();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): AssetClaimHTLC {
            return AssetClaimHTLC.deserialize(bytes);
        }
    }
    export class AssetExchangeAgreement extends pb_1.Message {
        constructor(data?: any[] | {
            type?: string;
            id?: string;
            locker?: string;
            recipient?: string;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], []);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("type" in data && data.type != undefined) {
                    this.type = data.type;
                }
                if ("id" in data && data.id != undefined) {
                    this.id = data.id;
                }
                if ("locker" in data && data.locker != undefined) {
                    this.locker = data.locker;
                }
                if ("recipient" in data && data.recipient != undefined) {
                    this.recipient = data.recipient;
                }
            }
        }
        get type() {
            return pb_1.Message.getField(this, 1) as string;
        }
        set type(value: string) {
            pb_1.Message.setField(this, 1, value);
        }
        get id() {
            return pb_1.Message.getField(this, 2) as string;
        }
        set id(value: string) {
            pb_1.Message.setField(this, 2, value);
        }
        get locker() {
            return pb_1.Message.getField(this, 3) as string;
        }
        set locker(value: string) {
            pb_1.Message.setField(this, 3, value);
        }
        get recipient() {
            return pb_1.Message.getField(this, 4) as string;
        }
        set recipient(value: string) {
            pb_1.Message.setField(this, 4, value);
        }
        static fromObject(data: {
            type?: string;
            id?: string;
            locker?: string;
            recipient?: string;
        }) {
            const message = new AssetExchangeAgreement({});
            if (data.type != null) {
                message.type = data.type;
            }
            if (data.id != null) {
                message.id = data.id;
            }
            if (data.locker != null) {
                message.locker = data.locker;
            }
            if (data.recipient != null) {
                message.recipient = data.recipient;
            }
            return message;
        }
        toObject() {
            const data: {
                type?: string;
                id?: string;
                locker?: string;
                recipient?: string;
            } = {};
            if (this.type != null) {
                data.type = this.type;
            }
            if (this.id != null) {
                data.id = this.id;
            }
            if (this.locker != null) {
                data.locker = this.locker;
            }
            if (this.recipient != null) {
                data.recipient = this.recipient;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (typeof this.type === "string" && this.type.length)
                writer.writeString(1, this.type);
            if (typeof this.id === "string" && this.id.length)
                writer.writeString(2, this.id);
            if (typeof this.locker === "string" && this.locker.length)
                writer.writeString(3, this.locker);
            if (typeof this.recipient === "string" && this.recipient.length)
                writer.writeString(4, this.recipient);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): AssetExchangeAgreement {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new AssetExchangeAgreement();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.type = reader.readString();
                        break;
                    case 2:
                        message.id = reader.readString();
                        break;
                    case 3:
                        message.locker = reader.readString();
                        break;
                    case 4:
                        message.recipient = reader.readString();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): AssetExchangeAgreement {
            return AssetExchangeAgreement.deserialize(bytes);
        }
    }
    export class FungibleAssetExchangeAgreement extends pb_1.Message {
        constructor(data?: any[] | {
            type?: string;
            numUnits?: number;
            locker?: string;
            recipient?: string;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], []);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("type" in data && data.type != undefined) {
                    this.type = data.type;
                }
                if ("numUnits" in data && data.numUnits != undefined) {
                    this.numUnits = data.numUnits;
                }
                if ("locker" in data && data.locker != undefined) {
                    this.locker = data.locker;
                }
                if ("recipient" in data && data.recipient != undefined) {
                    this.recipient = data.recipient;
                }
            }
        }
        get type() {
            return pb_1.Message.getField(this, 1) as string;
        }
        set type(value: string) {
            pb_1.Message.setField(this, 1, value);
        }
        get numUnits() {
            return pb_1.Message.getField(this, 2) as number;
        }
        set numUnits(value: number) {
            pb_1.Message.setField(this, 2, value);
        }
        get locker() {
            return pb_1.Message.getField(this, 3) as string;
        }
        set locker(value: string) {
            pb_1.Message.setField(this, 3, value);
        }
        get recipient() {
            return pb_1.Message.getField(this, 4) as string;
        }
        set recipient(value: string) {
            pb_1.Message.setField(this, 4, value);
        }
        static fromObject(data: {
            type?: string;
            numUnits?: number;
            locker?: string;
            recipient?: string;
        }) {
            const message = new FungibleAssetExchangeAgreement({});
            if (data.type != null) {
                message.type = data.type;
            }
            if (data.numUnits != null) {
                message.numUnits = data.numUnits;
            }
            if (data.locker != null) {
                message.locker = data.locker;
            }
            if (data.recipient != null) {
                message.recipient = data.recipient;
            }
            return message;
        }
        toObject() {
            const data: {
                type?: string;
                numUnits?: number;
                locker?: string;
                recipient?: string;
            } = {};
            if (this.type != null) {
                data.type = this.type;
            }
            if (this.numUnits != null) {
                data.numUnits = this.numUnits;
            }
            if (this.locker != null) {
                data.locker = this.locker;
            }
            if (this.recipient != null) {
                data.recipient = this.recipient;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (typeof this.type === "string" && this.type.length)
                writer.writeString(1, this.type);
            if (this.numUnits !== undefined)
                writer.writeUint64(2, this.numUnits);
            if (typeof this.locker === "string" && this.locker.length)
                writer.writeString(3, this.locker);
            if (typeof this.recipient === "string" && this.recipient.length)
                writer.writeString(4, this.recipient);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): FungibleAssetExchangeAgreement {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new FungibleAssetExchangeAgreement();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.type = reader.readString();
                        break;
                    case 2:
                        message.numUnits = reader.readUint64();
                        break;
                    case 3:
                        message.locker = reader.readString();
                        break;
                    case 4:
                        message.recipient = reader.readString();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): FungibleAssetExchangeAgreement {
            return FungibleAssetExchangeAgreement.deserialize(bytes);
        }
    }
    export class AssetContractHTLC extends pb_1.Message {
        constructor(data?: any[] | {
            contractId?: string;
            agreement?: AssetExchangeAgreement;
            lock?: AssetLockHTLC;
            claim?: AssetClaimHTLC;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], []);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("contractId" in data && data.contractId != undefined) {
                    this.contractId = data.contractId;
                }
                if ("agreement" in data && data.agreement != undefined) {
                    this.agreement = data.agreement;
                }
                if ("lock" in data && data.lock != undefined) {
                    this.lock = data.lock;
                }
                if ("claim" in data && data.claim != undefined) {
                    this.claim = data.claim;
                }
            }
        }
        get contractId() {
            return pb_1.Message.getField(this, 1) as string;
        }
        set contractId(value: string) {
            pb_1.Message.setField(this, 1, value);
        }
        get agreement() {
            return pb_1.Message.getWrapperField(this, AssetExchangeAgreement, 2) as AssetExchangeAgreement;
        }
        set agreement(value: AssetExchangeAgreement) {
            pb_1.Message.setWrapperField(this, 2, value);
        }
        get lock() {
            return pb_1.Message.getWrapperField(this, AssetLockHTLC, 3) as AssetLockHTLC;
        }
        set lock(value: AssetLockHTLC) {
            pb_1.Message.setWrapperField(this, 3, value);
        }
        get claim() {
            return pb_1.Message.getWrapperField(this, AssetClaimHTLC, 4) as AssetClaimHTLC;
        }
        set claim(value: AssetClaimHTLC) {
            pb_1.Message.setWrapperField(this, 4, value);
        }
        static fromObject(data: {
            contractId?: string;
            agreement?: ReturnType<typeof AssetExchangeAgreement.prototype.toObject>;
            lock?: ReturnType<typeof AssetLockHTLC.prototype.toObject>;
            claim?: ReturnType<typeof AssetClaimHTLC.prototype.toObject>;
        }) {
            const message = new AssetContractHTLC({});
            if (data.contractId != null) {
                message.contractId = data.contractId;
            }
            if (data.agreement != null) {
                message.agreement = AssetExchangeAgreement.fromObject(data.agreement);
            }
            if (data.lock != null) {
                message.lock = AssetLockHTLC.fromObject(data.lock);
            }
            if (data.claim != null) {
                message.claim = AssetClaimHTLC.fromObject(data.claim);
            }
            return message;
        }
        toObject() {
            const data: {
                contractId?: string;
                agreement?: ReturnType<typeof AssetExchangeAgreement.prototype.toObject>;
                lock?: ReturnType<typeof AssetLockHTLC.prototype.toObject>;
                claim?: ReturnType<typeof AssetClaimHTLC.prototype.toObject>;
            } = {};
            if (this.contractId != null) {
                data.contractId = this.contractId;
            }
            if (this.agreement != null) {
                data.agreement = this.agreement.toObject();
            }
            if (this.lock != null) {
                data.lock = this.lock.toObject();
            }
            if (this.claim != null) {
                data.claim = this.claim.toObject();
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (typeof this.contractId === "string" && this.contractId.length)
                writer.writeString(1, this.contractId);
            if (this.agreement !== undefined)
                writer.writeMessage(2, this.agreement, () => this.agreement.serialize(writer));
            if (this.lock !== undefined)
                writer.writeMessage(3, this.lock, () => this.lock.serialize(writer));
            if (this.claim !== undefined)
                writer.writeMessage(4, this.claim, () => this.claim.serialize(writer));
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): AssetContractHTLC {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new AssetContractHTLC();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.contractId = reader.readString();
                        break;
                    case 2:
                        reader.readMessage(message.agreement, () => message.agreement = AssetExchangeAgreement.deserialize(reader));
                        break;
                    case 3:
                        reader.readMessage(message.lock, () => message.lock = AssetLockHTLC.deserialize(reader));
                        break;
                    case 4:
                        reader.readMessage(message.claim, () => message.claim = AssetClaimHTLC.deserialize(reader));
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): AssetContractHTLC {
            return AssetContractHTLC.deserialize(bytes);
        }
    }
    export class FungibleAssetContractHTLC extends pb_1.Message {
        constructor(data?: any[] | {
            contractId?: string;
            agreement?: FungibleAssetExchangeAgreement;
            lock?: AssetLockHTLC;
            claim?: AssetClaimHTLC;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], []);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("contractId" in data && data.contractId != undefined) {
                    this.contractId = data.contractId;
                }
                if ("agreement" in data && data.agreement != undefined) {
                    this.agreement = data.agreement;
                }
                if ("lock" in data && data.lock != undefined) {
                    this.lock = data.lock;
                }
                if ("claim" in data && data.claim != undefined) {
                    this.claim = data.claim;
                }
            }
        }
        get contractId() {
            return pb_1.Message.getField(this, 1) as string;
        }
        set contractId(value: string) {
            pb_1.Message.setField(this, 1, value);
        }
        get agreement() {
            return pb_1.Message.getWrapperField(this, FungibleAssetExchangeAgreement, 2) as FungibleAssetExchangeAgreement;
        }
        set agreement(value: FungibleAssetExchangeAgreement) {
            pb_1.Message.setWrapperField(this, 2, value);
        }
        get lock() {
            return pb_1.Message.getWrapperField(this, AssetLockHTLC, 3) as AssetLockHTLC;
        }
        set lock(value: AssetLockHTLC) {
            pb_1.Message.setWrapperField(this, 3, value);
        }
        get claim() {
            return pb_1.Message.getWrapperField(this, AssetClaimHTLC, 4) as AssetClaimHTLC;
        }
        set claim(value: AssetClaimHTLC) {
            pb_1.Message.setWrapperField(this, 4, value);
        }
        static fromObject(data: {
            contractId?: string;
            agreement?: ReturnType<typeof FungibleAssetExchangeAgreement.prototype.toObject>;
            lock?: ReturnType<typeof AssetLockHTLC.prototype.toObject>;
            claim?: ReturnType<typeof AssetClaimHTLC.prototype.toObject>;
        }) {
            const message = new FungibleAssetContractHTLC({});
            if (data.contractId != null) {
                message.contractId = data.contractId;
            }
            if (data.agreement != null) {
                message.agreement = FungibleAssetExchangeAgreement.fromObject(data.agreement);
            }
            if (data.lock != null) {
                message.lock = AssetLockHTLC.fromObject(data.lock);
            }
            if (data.claim != null) {
                message.claim = AssetClaimHTLC.fromObject(data.claim);
            }
            return message;
        }
        toObject() {
            const data: {
                contractId?: string;
                agreement?: ReturnType<typeof FungibleAssetExchangeAgreement.prototype.toObject>;
                lock?: ReturnType<typeof AssetLockHTLC.prototype.toObject>;
                claim?: ReturnType<typeof AssetClaimHTLC.prototype.toObject>;
            } = {};
            if (this.contractId != null) {
                data.contractId = this.contractId;
            }
            if (this.agreement != null) {
                data.agreement = this.agreement.toObject();
            }
            if (this.lock != null) {
                data.lock = this.lock.toObject();
            }
            if (this.claim != null) {
                data.claim = this.claim.toObject();
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (typeof this.contractId === "string" && this.contractId.length)
                writer.writeString(1, this.contractId);
            if (this.agreement !== undefined)
                writer.writeMessage(2, this.agreement, () => this.agreement.serialize(writer));
            if (this.lock !== undefined)
                writer.writeMessage(3, this.lock, () => this.lock.serialize(writer));
            if (this.claim !== undefined)
                writer.writeMessage(4, this.claim, () => this.claim.serialize(writer));
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): FungibleAssetContractHTLC {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new FungibleAssetContractHTLC();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.contractId = reader.readString();
                        break;
                    case 2:
                        reader.readMessage(message.agreement, () => message.agreement = FungibleAssetExchangeAgreement.deserialize(reader));
                        break;
                    case 3:
                        reader.readMessage(message.lock, () => message.lock = AssetLockHTLC.deserialize(reader));
                        break;
                    case 4:
                        reader.readMessage(message.claim, () => message.claim = AssetClaimHTLC.deserialize(reader));
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): FungibleAssetContractHTLC {
            return FungibleAssetContractHTLC.deserialize(bytes);
        }
    }
}
