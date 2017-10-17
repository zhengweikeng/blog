/*eslint-disable block-scoped-var, no-redeclare, no-control-regex, no-prototype-builtins*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

$root.Address = (function() {

    /**
     * Properties of an Address.
     * @exports IAddress
     * @interface IAddress
     * @property {string} [addr] Address addr
     * @property {number} [code] Address code
     */

    /**
     * Constructs a new Address.
     * @exports Address
     * @classdesc Represents an Address.
     * @constructor
     * @param {IAddress=} [properties] Properties to set
     */
    function Address(properties) {
        if (properties)
            for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                if (properties[keys[i]] != null)
                    this[keys[i]] = properties[keys[i]];
    }

    /**
     * Address addr.
     * @member {string}addr
     * @memberof Address
     * @instance
     */
    Address.prototype.addr = "";

    /**
     * Address code.
     * @member {number}code
     * @memberof Address
     * @instance
     */
    Address.prototype.code = 0;

    /**
     * Creates a new Address instance using the specified properties.
     * @function create
     * @memberof Address
     * @static
     * @param {IAddress=} [properties] Properties to set
     * @returns {Address} Address instance
     */
    Address.create = function create(properties) {
        return new Address(properties);
    };

    /**
     * Encodes the specified Address message. Does not implicitly {@link Address.verify|verify} messages.
     * @function encode
     * @memberof Address
     * @static
     * @param {IAddress} message Address message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    Address.encode = function encode(message, writer) {
        if (!writer)
            writer = $Writer.create();
        if (message.addr != null && message.hasOwnProperty("addr"))
            writer.uint32(/* id 1, wireType 2 =*/10).string(message.addr);
        if (message.code != null && message.hasOwnProperty("code"))
            writer.uint32(/* id 2, wireType 0 =*/16).int32(message.code);
        return writer;
    };

    /**
     * Encodes the specified Address message, length delimited. Does not implicitly {@link Address.verify|verify} messages.
     * @function encodeDelimited
     * @memberof Address
     * @static
     * @param {IAddress} message Address message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    Address.encodeDelimited = function encodeDelimited(message, writer) {
        return this.encode(message, writer).ldelim();
    };

    /**
     * Decodes an Address message from the specified reader or buffer.
     * @function decode
     * @memberof Address
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @param {number} [length] Message length if known beforehand
     * @returns {Address} Address
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    Address.decode = function decode(reader, length) {
        if (!(reader instanceof $Reader))
            reader = $Reader.create(reader);
        var end = length === undefined ? reader.len : reader.pos + length, message = new $root.Address();
        while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
            case 1:
                message.addr = reader.string();
                break;
            case 2:
                message.code = reader.int32();
                break;
            default:
                reader.skipType(tag & 7);
                break;
            }
        }
        return message;
    };

    /**
     * Decodes an Address message from the specified reader or buffer, length delimited.
     * @function decodeDelimited
     * @memberof Address
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @returns {Address} Address
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    Address.decodeDelimited = function decodeDelimited(reader) {
        if (!(reader instanceof $Reader))
            reader = new $Reader(reader);
        return this.decode(reader, reader.uint32());
    };

    /**
     * Verifies an Address message.
     * @function verify
     * @memberof Address
     * @static
     * @param {Object.<string,*>} message Plain object to verify
     * @returns {string|null} `null` if valid, otherwise the reason why it is not
     */
    Address.verify = function verify(message) {
        if (typeof message !== "object" || message === null)
            return "object expected";
        if (message.addr != null && message.hasOwnProperty("addr"))
            if (!$util.isString(message.addr))
                return "addr: string expected";
        if (message.code != null && message.hasOwnProperty("code"))
            if (!$util.isInteger(message.code))
                return "code: integer expected";
        return null;
    };

    /**
     * Creates an Address message from a plain object. Also converts values to their respective internal types.
     * @function fromObject
     * @memberof Address
     * @static
     * @param {Object.<string,*>} object Plain object
     * @returns {Address} Address
     */
    Address.fromObject = function fromObject(object) {
        if (object instanceof $root.Address)
            return object;
        var message = new $root.Address();
        if (object.addr != null)
            message.addr = String(object.addr);
        if (object.code != null)
            message.code = object.code | 0;
        return message;
    };

    /**
     * Creates a plain object from an Address message. Also converts values to other types if specified.
     * @function toObject
     * @memberof Address
     * @static
     * @param {Address} message Address
     * @param {$protobuf.IConversionOptions} [options] Conversion options
     * @returns {Object.<string,*>} Plain object
     */
    Address.toObject = function toObject(message, options) {
        if (!options)
            options = {};
        var object = {};
        if (options.defaults) {
            object.addr = "";
            object.code = 0;
        }
        if (message.addr != null && message.hasOwnProperty("addr"))
            object.addr = message.addr;
        if (message.code != null && message.hasOwnProperty("code"))
            object.code = message.code;
        return object;
    };

    /**
     * Converts this Address to JSON.
     * @function toJSON
     * @memberof Address
     * @instance
     * @returns {Object.<string,*>} JSON object
     */
    Address.prototype.toJSON = function toJSON() {
        return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
    };

    return Address;
})();

$root.Phone = (function() {

    /**
     * Properties of a Phone.
     * @exports IPhone
     * @interface IPhone
     * @property {number|Long} [phoneNum] Phone phoneNum
     */

    /**
     * Constructs a new Phone.
     * @exports Phone
     * @classdesc Represents a Phone.
     * @constructor
     * @param {IPhone=} [properties] Properties to set
     */
    function Phone(properties) {
        if (properties)
            for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                if (properties[keys[i]] != null)
                    this[keys[i]] = properties[keys[i]];
    }

    /**
     * Phone phoneNum.
     * @member {number|Long}phoneNum
     * @memberof Phone
     * @instance
     */
    Phone.prototype.phoneNum = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

    /**
     * Creates a new Phone instance using the specified properties.
     * @function create
     * @memberof Phone
     * @static
     * @param {IPhone=} [properties] Properties to set
     * @returns {Phone} Phone instance
     */
    Phone.create = function create(properties) {
        return new Phone(properties);
    };

    /**
     * Encodes the specified Phone message. Does not implicitly {@link Phone.verify|verify} messages.
     * @function encode
     * @memberof Phone
     * @static
     * @param {IPhone} message Phone message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    Phone.encode = function encode(message, writer) {
        if (!writer)
            writer = $Writer.create();
        if (message.phoneNum != null && message.hasOwnProperty("phoneNum"))
            writer.uint32(/* id 1, wireType 0 =*/8).int64(message.phoneNum);
        return writer;
    };

    /**
     * Encodes the specified Phone message, length delimited. Does not implicitly {@link Phone.verify|verify} messages.
     * @function encodeDelimited
     * @memberof Phone
     * @static
     * @param {IPhone} message Phone message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    Phone.encodeDelimited = function encodeDelimited(message, writer) {
        return this.encode(message, writer).ldelim();
    };

    /**
     * Decodes a Phone message from the specified reader or buffer.
     * @function decode
     * @memberof Phone
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @param {number} [length] Message length if known beforehand
     * @returns {Phone} Phone
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    Phone.decode = function decode(reader, length) {
        if (!(reader instanceof $Reader))
            reader = $Reader.create(reader);
        var end = length === undefined ? reader.len : reader.pos + length, message = new $root.Phone();
        while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
            case 1:
                message.phoneNum = reader.int64();
                break;
            default:
                reader.skipType(tag & 7);
                break;
            }
        }
        return message;
    };

    /**
     * Decodes a Phone message from the specified reader or buffer, length delimited.
     * @function decodeDelimited
     * @memberof Phone
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @returns {Phone} Phone
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    Phone.decodeDelimited = function decodeDelimited(reader) {
        if (!(reader instanceof $Reader))
            reader = new $Reader(reader);
        return this.decode(reader, reader.uint32());
    };

    /**
     * Verifies a Phone message.
     * @function verify
     * @memberof Phone
     * @static
     * @param {Object.<string,*>} message Plain object to verify
     * @returns {string|null} `null` if valid, otherwise the reason why it is not
     */
    Phone.verify = function verify(message) {
        if (typeof message !== "object" || message === null)
            return "object expected";
        if (message.phoneNum != null && message.hasOwnProperty("phoneNum"))
            if (!$util.isInteger(message.phoneNum) && !(message.phoneNum && $util.isInteger(message.phoneNum.low) && $util.isInteger(message.phoneNum.high)))
                return "phoneNum: integer|Long expected";
        return null;
    };

    /**
     * Creates a Phone message from a plain object. Also converts values to their respective internal types.
     * @function fromObject
     * @memberof Phone
     * @static
     * @param {Object.<string,*>} object Plain object
     * @returns {Phone} Phone
     */
    Phone.fromObject = function fromObject(object) {
        if (object instanceof $root.Phone)
            return object;
        var message = new $root.Phone();
        if (object.phoneNum != null)
            if ($util.Long)
                (message.phoneNum = $util.Long.fromValue(object.phoneNum)).unsigned = false;
            else if (typeof object.phoneNum === "string")
                message.phoneNum = parseInt(object.phoneNum, 10);
            else if (typeof object.phoneNum === "number")
                message.phoneNum = object.phoneNum;
            else if (typeof object.phoneNum === "object")
                message.phoneNum = new $util.LongBits(object.phoneNum.low >>> 0, object.phoneNum.high >>> 0).toNumber();
        return message;
    };

    /**
     * Creates a plain object from a Phone message. Also converts values to other types if specified.
     * @function toObject
     * @memberof Phone
     * @static
     * @param {Phone} message Phone
     * @param {$protobuf.IConversionOptions} [options] Conversion options
     * @returns {Object.<string,*>} Plain object
     */
    Phone.toObject = function toObject(message, options) {
        if (!options)
            options = {};
        var object = {};
        if (options.defaults)
            if ($util.Long) {
                var long = new $util.Long(0, 0, false);
                object.phoneNum = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
            } else
                object.phoneNum = options.longs === String ? "0" : 0;
        if (message.phoneNum != null && message.hasOwnProperty("phoneNum"))
            if (typeof message.phoneNum === "number")
                object.phoneNum = options.longs === String ? String(message.phoneNum) : message.phoneNum;
            else
                object.phoneNum = options.longs === String ? $util.Long.prototype.toString.call(message.phoneNum) : options.longs === Number ? new $util.LongBits(message.phoneNum.low >>> 0, message.phoneNum.high >>> 0).toNumber() : message.phoneNum;
        return object;
    };

    /**
     * Converts this Phone to JSON.
     * @function toJSON
     * @memberof Phone
     * @instance
     * @returns {Object.<string,*>} JSON object
     */
    Phone.prototype.toJSON = function toJSON() {
        return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
    };

    return Phone;
})();

/**
 * Pets enum.
 * @exports Pets
 * @enum {string}
 * @property {number} DOG=0 DOG value
 * @property {number} CAT=1 CAT value
 */
$root.Pets = (function() {
    var valuesById = {}, values = Object.create(valuesById);
    values[valuesById[0] = "DOG"] = 0;
    values[valuesById[1] = "CAT"] = 1;
    return values;
})();

$root.Person = (function() {

    /**
     * Properties of a Person.
     * @exports IPerson
     * @interface IPerson
     * @property {string} [name] Person name
     * @property {number} [age] Person age
     * @property {string} [email] Person email
     * @property {Uint8Array} [foo] Person foo
     * @property {IAddress} [address] Person address
     * @property {Array.<string>} [favorite] Person favorite
     * @property {Object.<string,IPhone>} [phone] Person phone
     * @property {boolean} [sex] Person sex
     * @property {string} [imageUrl] Person imageUrl
     * @property {Uint8Array} [imageData] Person imageData
     * @property {Pets} [pet] Person pet
     */

    /**
     * Constructs a new Person.
     * @exports Person
     * @classdesc Represents a Person.
     * @constructor
     * @param {IPerson=} [properties] Properties to set
     */
    function Person(properties) {
        this.favorite = [];
        this.phone = {};
        if (properties)
            for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                if (properties[keys[i]] != null)
                    this[keys[i]] = properties[keys[i]];
    }

    /**
     * Person name.
     * @member {string}name
     * @memberof Person
     * @instance
     */
    Person.prototype.name = "";

    /**
     * Person age.
     * @member {number}age
     * @memberof Person
     * @instance
     */
    Person.prototype.age = 0;

    /**
     * Person email.
     * @member {string}email
     * @memberof Person
     * @instance
     */
    Person.prototype.email = "";

    /**
     * Person foo.
     * @member {Uint8Array}foo
     * @memberof Person
     * @instance
     */
    Person.prototype.foo = $util.newBuffer([]);

    /**
     * Person address.
     * @member {(IAddress|null|undefined)}address
     * @memberof Person
     * @instance
     */
    Person.prototype.address = null;

    /**
     * Person favorite.
     * @member {Array.<string>}favorite
     * @memberof Person
     * @instance
     */
    Person.prototype.favorite = $util.emptyArray;

    /**
     * Person phone.
     * @member {Object.<string,IPhone>}phone
     * @memberof Person
     * @instance
     */
    Person.prototype.phone = $util.emptyObject;

    /**
     * Person sex.
     * @member {boolean}sex
     * @memberof Person
     * @instance
     */
    Person.prototype.sex = false;

    /**
     * Person imageUrl.
     * @member {string}imageUrl
     * @memberof Person
     * @instance
     */
    Person.prototype.imageUrl = "";

    /**
     * Person imageData.
     * @member {Uint8Array}imageData
     * @memberof Person
     * @instance
     */
    Person.prototype.imageData = $util.newBuffer([]);

    /**
     * Person pet.
     * @member {Pets}pet
     * @memberof Person
     * @instance
     */
    Person.prototype.pet = 0;

    // OneOf field names bound to virtual getters and setters
    var $oneOfFields;

    /**
     * Person avatar.
     * @member {string|undefined} avatar
     * @memberof Person
     * @instance
     */
    Object.defineProperty(Person.prototype, "avatar", {
        get: $util.oneOfGetter($oneOfFields = ["imageUrl", "imageData"]),
        set: $util.oneOfSetter($oneOfFields)
    });

    /**
     * Creates a new Person instance using the specified properties.
     * @function create
     * @memberof Person
     * @static
     * @param {IPerson=} [properties] Properties to set
     * @returns {Person} Person instance
     */
    Person.create = function create(properties) {
        return new Person(properties);
    };

    /**
     * Encodes the specified Person message. Does not implicitly {@link Person.verify|verify} messages.
     * @function encode
     * @memberof Person
     * @static
     * @param {IPerson} message Person message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    Person.encode = function encode(message, writer) {
        if (!writer)
            writer = $Writer.create();
        if (message.name != null && message.hasOwnProperty("name"))
            writer.uint32(/* id 1, wireType 2 =*/10).string(message.name);
        if (message.age != null && message.hasOwnProperty("age"))
            writer.uint32(/* id 2, wireType 0 =*/16).int32(message.age);
        if (message.email != null && message.hasOwnProperty("email"))
            writer.uint32(/* id 3, wireType 2 =*/26).string(message.email);
        if (message.foo != null && message.hasOwnProperty("foo"))
            writer.uint32(/* id 4, wireType 2 =*/34).bytes(message.foo);
        if (message.address != null && message.hasOwnProperty("address"))
            $root.Address.encode(message.address, writer.uint32(/* id 5, wireType 2 =*/42).fork()).ldelim();
        if (message.favorite != null && message.favorite.length)
            for (var i = 0; i < message.favorite.length; ++i)
                writer.uint32(/* id 6, wireType 2 =*/50).string(message.favorite[i]);
        if (message.phone != null && message.hasOwnProperty("phone"))
            for (var keys = Object.keys(message.phone), i = 0; i < keys.length; ++i) {
                writer.uint32(/* id 7, wireType 2 =*/58).fork().uint32(/* id 1, wireType 2 =*/10).string(keys[i]);
                $root.Phone.encode(message.phone[keys[i]], writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim().ldelim();
            }
        if (message.sex != null && message.hasOwnProperty("sex"))
            writer.uint32(/* id 8, wireType 0 =*/64).bool(message.sex);
        if (message.imageUrl != null && message.hasOwnProperty("imageUrl"))
            writer.uint32(/* id 9, wireType 2 =*/74).string(message.imageUrl);
        if (message.imageData != null && message.hasOwnProperty("imageData"))
            writer.uint32(/* id 10, wireType 2 =*/82).bytes(message.imageData);
        if (message.pet != null && message.hasOwnProperty("pet"))
            writer.uint32(/* id 11, wireType 0 =*/88).int32(message.pet);
        return writer;
    };

    /**
     * Encodes the specified Person message, length delimited. Does not implicitly {@link Person.verify|verify} messages.
     * @function encodeDelimited
     * @memberof Person
     * @static
     * @param {IPerson} message Person message or plain object to encode
     * @param {$protobuf.Writer} [writer] Writer to encode to
     * @returns {$protobuf.Writer} Writer
     */
    Person.encodeDelimited = function encodeDelimited(message, writer) {
        return this.encode(message, writer).ldelim();
    };

    /**
     * Decodes a Person message from the specified reader or buffer.
     * @function decode
     * @memberof Person
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @param {number} [length] Message length if known beforehand
     * @returns {Person} Person
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    Person.decode = function decode(reader, length) {
        if (!(reader instanceof $Reader))
            reader = $Reader.create(reader);
        var end = length === undefined ? reader.len : reader.pos + length, message = new $root.Person(), key;
        while (reader.pos < end) {
            var tag = reader.uint32();
            switch (tag >>> 3) {
            case 1:
                message.name = reader.string();
                break;
            case 2:
                message.age = reader.int32();
                break;
            case 3:
                message.email = reader.string();
                break;
            case 4:
                message.foo = reader.bytes();
                break;
            case 5:
                message.address = $root.Address.decode(reader, reader.uint32());
                break;
            case 6:
                if (!(message.favorite && message.favorite.length))
                    message.favorite = [];
                message.favorite.push(reader.string());
                break;
            case 7:
                reader.skip().pos++;
                if (message.phone === $util.emptyObject)
                    message.phone = {};
                key = reader.string();
                reader.pos++;
                message.phone[key] = $root.Phone.decode(reader, reader.uint32());
                break;
            case 8:
                message.sex = reader.bool();
                break;
            case 9:
                message.imageUrl = reader.string();
                break;
            case 10:
                message.imageData = reader.bytes();
                break;
            case 11:
                message.pet = reader.int32();
                break;
            default:
                reader.skipType(tag & 7);
                break;
            }
        }
        return message;
    };

    /**
     * Decodes a Person message from the specified reader or buffer, length delimited.
     * @function decodeDelimited
     * @memberof Person
     * @static
     * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
     * @returns {Person} Person
     * @throws {Error} If the payload is not a reader or valid buffer
     * @throws {$protobuf.util.ProtocolError} If required fields are missing
     */
    Person.decodeDelimited = function decodeDelimited(reader) {
        if (!(reader instanceof $Reader))
            reader = new $Reader(reader);
        return this.decode(reader, reader.uint32());
    };

    /**
     * Verifies a Person message.
     * @function verify
     * @memberof Person
     * @static
     * @param {Object.<string,*>} message Plain object to verify
     * @returns {string|null} `null` if valid, otherwise the reason why it is not
     */
    Person.verify = function verify(message) {
        if (typeof message !== "object" || message === null)
            return "object expected";
        var properties = {};
        if (message.name != null && message.hasOwnProperty("name"))
            if (!$util.isString(message.name))
                return "name: string expected";
        if (message.age != null && message.hasOwnProperty("age"))
            if (!$util.isInteger(message.age))
                return "age: integer expected";
        if (message.email != null && message.hasOwnProperty("email"))
            if (!$util.isString(message.email))
                return "email: string expected";
        if (message.foo != null && message.hasOwnProperty("foo"))
            if (!(message.foo && typeof message.foo.length === "number" || $util.isString(message.foo)))
                return "foo: buffer expected";
        if (message.address != null && message.hasOwnProperty("address")) {
            var error = $root.Address.verify(message.address);
            if (error)
                return "address." + error;
        }
        if (message.favorite != null && message.hasOwnProperty("favorite")) {
            if (!Array.isArray(message.favorite))
                return "favorite: array expected";
            for (var i = 0; i < message.favorite.length; ++i)
                if (!$util.isString(message.favorite[i]))
                    return "favorite: string[] expected";
        }
        if (message.phone != null && message.hasOwnProperty("phone")) {
            if (!$util.isObject(message.phone))
                return "phone: object expected";
            var key = Object.keys(message.phone);
            for (var i = 0; i < key.length; ++i) {
                error = $root.Phone.verify(message.phone[key[i]]);
                if (error)
                    return "phone." + error;
            }
        }
        if (message.sex != null && message.hasOwnProperty("sex"))
            if (typeof message.sex !== "boolean")
                return "sex: boolean expected";
        if (message.imageUrl != null && message.hasOwnProperty("imageUrl")) {
            properties.avatar = 1;
            if (!$util.isString(message.imageUrl))
                return "imageUrl: string expected";
        }
        if (message.imageData != null && message.hasOwnProperty("imageData")) {
            if (properties.avatar === 1)
                return "avatar: multiple values";
            properties.avatar = 1;
            if (!(message.imageData && typeof message.imageData.length === "number" || $util.isString(message.imageData)))
                return "imageData: buffer expected";
        }
        if (message.pet != null && message.hasOwnProperty("pet"))
            switch (message.pet) {
            default:
                return "pet: enum value expected";
            case 0:
            case 1:
                break;
            }
        return null;
    };

    /**
     * Creates a Person message from a plain object. Also converts values to their respective internal types.
     * @function fromObject
     * @memberof Person
     * @static
     * @param {Object.<string,*>} object Plain object
     * @returns {Person} Person
     */
    Person.fromObject = function fromObject(object) {
        if (object instanceof $root.Person)
            return object;
        var message = new $root.Person();
        if (object.name != null)
            message.name = String(object.name);
        if (object.age != null)
            message.age = object.age | 0;
        if (object.email != null)
            message.email = String(object.email);
        if (object.foo != null)
            if (typeof object.foo === "string")
                $util.base64.decode(object.foo, message.foo = $util.newBuffer($util.base64.length(object.foo)), 0);
            else if (object.foo.length)
                message.foo = object.foo;
        if (object.address != null) {
            if (typeof object.address !== "object")
                throw TypeError(".Person.address: object expected");
            message.address = $root.Address.fromObject(object.address);
        }
        if (object.favorite) {
            if (!Array.isArray(object.favorite))
                throw TypeError(".Person.favorite: array expected");
            message.favorite = [];
            for (var i = 0; i < object.favorite.length; ++i)
                message.favorite[i] = String(object.favorite[i]);
        }
        if (object.phone) {
            if (typeof object.phone !== "object")
                throw TypeError(".Person.phone: object expected");
            message.phone = {};
            for (var keys = Object.keys(object.phone), i = 0; i < keys.length; ++i) {
                if (typeof object.phone[keys[i]] !== "object")
                    throw TypeError(".Person.phone: object expected");
                message.phone[keys[i]] = $root.Phone.fromObject(object.phone[keys[i]]);
            }
        }
        if (object.sex != null)
            message.sex = Boolean(object.sex);
        if (object.imageUrl != null)
            message.imageUrl = String(object.imageUrl);
        if (object.imageData != null)
            if (typeof object.imageData === "string")
                $util.base64.decode(object.imageData, message.imageData = $util.newBuffer($util.base64.length(object.imageData)), 0);
            else if (object.imageData.length)
                message.imageData = object.imageData;
        switch (object.pet) {
        case "DOG":
        case 0:
            message.pet = 0;
            break;
        case "CAT":
        case 1:
            message.pet = 1;
            break;
        }
        return message;
    };

    /**
     * Creates a plain object from a Person message. Also converts values to other types if specified.
     * @function toObject
     * @memberof Person
     * @static
     * @param {Person} message Person
     * @param {$protobuf.IConversionOptions} [options] Conversion options
     * @returns {Object.<string,*>} Plain object
     */
    Person.toObject = function toObject(message, options) {
        if (!options)
            options = {};
        var object = {};
        if (options.arrays || options.defaults)
            object.favorite = [];
        if (options.objects || options.defaults)
            object.phone = {};
        if (options.defaults) {
            object.name = "";
            object.age = 0;
            object.email = "";
            object.foo = options.bytes === String ? "" : [];
            object.address = null;
            object.sex = false;
            object.pet = options.enums === String ? "DOG" : 0;
        }
        if (message.name != null && message.hasOwnProperty("name"))
            object.name = message.name;
        if (message.age != null && message.hasOwnProperty("age"))
            object.age = message.age;
        if (message.email != null && message.hasOwnProperty("email"))
            object.email = message.email;
        if (message.foo != null && message.hasOwnProperty("foo"))
            object.foo = options.bytes === String ? $util.base64.encode(message.foo, 0, message.foo.length) : options.bytes === Array ? Array.prototype.slice.call(message.foo) : message.foo;
        if (message.address != null && message.hasOwnProperty("address"))
            object.address = $root.Address.toObject(message.address, options);
        if (message.favorite && message.favorite.length) {
            object.favorite = [];
            for (var j = 0; j < message.favorite.length; ++j)
                object.favorite[j] = message.favorite[j];
        }
        var keys2;
        if (message.phone && (keys2 = Object.keys(message.phone)).length) {
            object.phone = {};
            for (var j = 0; j < keys2.length; ++j)
                object.phone[keys2[j]] = $root.Phone.toObject(message.phone[keys2[j]], options);
        }
        if (message.sex != null && message.hasOwnProperty("sex"))
            object.sex = message.sex;
        if (message.imageUrl != null && message.hasOwnProperty("imageUrl")) {
            object.imageUrl = message.imageUrl;
            if (options.oneofs)
                object.avatar = "imageUrl";
        }
        if (message.imageData != null && message.hasOwnProperty("imageData")) {
            object.imageData = options.bytes === String ? $util.base64.encode(message.imageData, 0, message.imageData.length) : options.bytes === Array ? Array.prototype.slice.call(message.imageData) : message.imageData;
            if (options.oneofs)
                object.avatar = "imageData";
        }
        if (message.pet != null && message.hasOwnProperty("pet"))
            object.pet = options.enums === String ? $root.Pets[message.pet] : message.pet;
        return object;
    };

    /**
     * Converts this Person to JSON.
     * @function toJSON
     * @memberof Person
     * @instance
     * @returns {Object.<string,*>} JSON object
     */
    Person.prototype.toJSON = function toJSON() {
        return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
    };

    return Person;
})();

module.exports = $root;
