# grpc校验规则

`protoc-gen-validate `(https://github.com/envoyproxy/protoc-gen-validate)

# Constraint Rules

## Numerics

> All numeric types (`float`, `double`, `int32`, `int64`, `uint32`, `uint64`, `sint32`, `sint64`, `fixed32`, `fixed64`, `sfixed32`, `sfixed64`) share the same rules.

- **const**: the field must be _exactly_ the specified value.

  ```protobuf
  // x must equal 1.23 exactly
  float x = 1 [(validate.rules).float.const = 1.23];
  ```

- **lt/lte/gt/gte**: these inequalities (`<`, `<=`, `>`, `>=`, respectively) allow for deriving ranges in which the field must reside.

  ```protobuf
  // x must be less than 10
  int32 x = 1 [(validate.rules).int32.lt = 10];

  // x must be greater than or equal to 20
  uint64 x = 1 [(validate.rules).uint64.gte = 20];

  // x must be in the range [30, 40)
  fixed32 x = 1 [(validate.rules).fixed32 = {gte:30, lt: 40}];
  ```

  Inverting the values of `lt(e)` and `gt(e)` is valid and creates an exclusive range.

  ```protobuf
  // x must be outside the range [30, 40)
  double x = 1 [(validate.rules).double = {lt:30, gte:40}];
  ```

- **in/not_in**: these two rules permit specifying allow/denylists for the values of a field.

  ```protobuf
  // x must be either 1, 2, or 3
  uint32 x = 1 [(validate.rules).uint32 = {in: [1,2,3]}];

  // x cannot be 0 nor 0.99
  float x = 1 [(validate.rules).float = {not_in: [0, 0.99]}];
  ```

- **ignore_empty**: this rule specifies that if field is empty or set to the default value, to ignore any validation rules. These are typically useful where being able to unset a field in an update request, or to skip validation for optional fields where switching to WKTs is not feasible.

  ```protobuf
  unint32 x = 1 [(validate.rules).uint32 = {ignore_empty: true, gte: 200}];
  ```

## Bools

- **const**: the field must be _exactly_ the specified value.

  ```protobuf
  // x must be set to true
  bool x = 1 [(validate.rules).bool.const = true];
  
  // x cannot be set to true
  bool x = 1 [(validate.rules).bool.const = false];
  ```

## Strings

- **const**: the field must be _exactly_ the specified value.

  ```protobuf
  // x must be set to "foo"
  string x = 1 [(validate.rules).string.const = "foo"];
  ```

- **len/min_len/max_len**: these rules constrain the number of characters (Unicode code points) in the field. Note that the number of characters may differ from the number of bytes in the string. The string is considered as-is, and does not normalize.

  ```protobuf
  // x must be exactly 5 characters long
  string x = 1 [(validate.rules).string.len = 5];

  // x must be at least 3 characters long
  string x = 1 [(validate.rules).string.min_len = 3];

  // x must be between 5 and 10 characters, inclusive
  string x = 1 [(validate.rules).string = {min_len: 5, max_len: 10}];
  ```

- **min_bytes/max_bytes**: these rules constrain the number of bytes in the field.

  ```protobuf
  // x must be at most 15 bytes long
  string x = 1 [(validate.rules).string.max_bytes = 15];

  // x must be between 128 and 1024 bytes long
  string x = 1 [(validate.rules).string = {min_bytes: 128, max_bytes: 1024}];
  ```

- **pattern**: the field must match the specified [RE2-compliant][re2] regular expression. The included expression should elide any delimiters (ie, `/\d+/` should just be `\d+`).

  ```protobuf
  // x must be a non-empty, case-insensitive hexadecimal string
  string x = 1 [(validate.rules).string.pattern = "(?i)^[0-9a-f]+$"];
  ```

- **prefix/suffix/contains/not_contains**: the field must contain the specified substring in an optionally explicit location, or not contain the specified substring.

  ```protobuf
  // x must begin with "foo"
  string x = 1 [(validate.rules).string.prefix = "foo"];

  // x must end with "bar"
  string x = 1 [(validate.rules).string.suffix = "bar"];

  // x must contain "baz" anywhere inside it
  string x = 1 [(validate.rules).string.contains = "baz"];
  
  // x cannot contain "baz" anywhere inside it
  string x = 1 [(validate.rules).string.not_contains = "baz"];

  // x must begin with "fizz" and end with "buzz"
  string x = 1 [(validate.rules).string = {prefix: "fizz", suffix: "buzz"}];

  // x must end with ".proto" and be less than 64 characters
  string x = 1 [(validate.rules).string = {suffix: ".proto", max_len:64}];
  ```

- **in/not_in**: these two rules permit specifying allow/denylists for the values of a field.

  ```protobuf
  // x must be either "foo", "bar", or "baz"
  string x = 1 [(validate.rules).string = {in: ["foo", "bar", "baz"]}];

  // x cannot be "fizz" nor "buzz"
  string x = 1 [(validate.rules).string = {not_in: ["fizz", "buzz"]}];
  ```

- **ignore_empty**: this rule specifies that if field is empty or set to the default value, to ignore any validation rules. These are typically useful where being able to unset a field in an update request, or to skip validation for optional fields where switching to WKTs is not feasible.

  ```protobuf
  string CountryCode = 1 [(validate.rules).string = {ignore_empty: true, len: 2}];
  ```

- **well-known formats**: these rules provide advanced constraints for common string patterns. These constraints will typically be more permissive and performant than equivalent regular expression patterns, while providing more explanatory failure descriptions.

  ```protobuf
  // x must be a valid email address (via RFC 1034)
  string x = 1 [(validate.rules).string.email = true];
  
  // x must be a valid address (IP or Hostname).
  string x = 1 [(validate.rules).string.address = true];
  
  // x must be a valid hostname (via RFC 1034)
  string x = 1 [(validate.rules).string.hostname = true];
  
  // x must be a valid IP address (either v4 or v6)
  string x = 1 [(validate.rules).string.ip = true];
  
  // x must be a valid IPv4 address
  // eg: "192.168.0.1"
  string x = 1 [(validate.rules).string.ipv4 = true];
  
  // x must be a valid IPv6 address
  // eg: "fe80::3"
  string x = 1 [(validate.rules).string.ipv6 = true];
  
  // x must be a valid absolute URI (via RFC 3986)
  string x = 1 [(validate.rules).string.uri = true];
  
  // x must be a valid URI reference (either absolute or relative)
  string x = 1 [(validate.rules).string.uri_ref = true];
  
  // x must be a valid UUID (via RFC 4122)
  string x = 1 [(validate.rules).string.uuid = true];
  
  // x must conform to a well known regex for HTTP header names (via RFC 7230)
  string x = 1 [(validate.rules).string.well_known_regex = HTTP_HEADER_NAME]
  
  // x must conform to a well known regex for HTTP header values (via RFC 7230) 
  string x = 1 [(validate.rules).string.well_known_regex = HTTP_HEADER_VALUE];
  
  // x must conform to a well known regex for headers, disallowing \r\n\0 characters.
  string x = 1 [(validate.rules).string {well_known_regex: HTTP_HEADER_VALUE, strict: false}];
  ```

## Bytes

> Literal values should be expressed with strings, using escaping where necessary.

- **const**: the field must be _exactly_ the specified value.

  ```protobuf
  // x must be set to "foo" ("\x66\x6f\x6f")
  bytes x = 1 [(validate.rules).bytes.const = "foo"];

  // x must be set to "\xf0\x90\x28\xbc"
  bytes x = 1 [(validate.rules).bytes.const = "\xf0\x90\x28\xbc"];
  ```

- **len/min_len/max_len**: these rules constrain the number of bytes in the field.

  ```protobuf
  // x must be exactly 3 bytes
  bytes x = 1 [(validate.rules).bytes.len = 3];

  // x must be at least 3 bytes long
  bytes x = 1 [(validate.rules).bytes.min_len = 3];

  // x must be between 5 and 10 bytes, inclusive
  bytes x = 1 [(validate.rules).bytes = {min_len: 5, max_len: 10}];
  ```

- **pattern**: the field must match the specified [RE2-compliant][re2] regular expression. The included expression should elide any delimiters (ie, `/\d+/` should just be `\d+`).

  ```protobuf
  // x must be a non-empty, ASCII byte sequence
  bytes x = 1 [(validate.rules).bytes.pattern = "^[\x00-\x7F]+$"];
  ```

- **prefix/suffix/contains**: the field must contain the specified byte sequence in an optionally explicit location.

  ```protobuf
  // x must begin with "\x99"
  bytes x = 1 [(validate.rules).bytes.prefix = "\x99"];

  // x must end with "buz\x7a"
  bytes x = 1 [(validate.rules).bytes.suffix = "buz\x7a"];

  // x must contain "baz" anywhere inside it
  bytes x = 1 [(validate.rules).bytes.contains = "baz"];
  ```

- **in/not_in**: these two rules permit specifying allow/denylists for the values of a field.

  ```protobuf
  // x must be either "foo", "bar", or "baz"
  bytes x = 1 [(validate.rules).bytes = {in: ["foo", "bar", "baz"]}];

  // x cannot be "fizz" nor "buzz"
  bytes x = 1 [(validate.rules).bytes = {not_in: ["fizz", "buzz"]}];
  ```

- **ignore_empty**: this rule specifies that if field is empty or set to the default value, to ignore any validation rules. These are typically useful where being able to unset a field in an update request, or to skip validation for optional fields where switching to WKTs is not feasible.

  ```protobuf
  bytes x = 1 [(validate.rules).bytes = {ignore_empty: true, in: ["foo", "bar", "baz"]}];
  ```

- **well-known formats**: these rules provide advanced constraints for common patterns. These constraints will typically be more permissive and performant than equivalent regular expression patterns, while providing more explanatory failure descriptions.

  ```protobuf
  // x must be a valid IP address (either v4 or v6) in byte format
  bytes x = 1 [(validate.rules).bytes.ip = true];
  
  // x must be a valid IPv4 address in byte format
  // eg: "\xC0\xA8\x00\x01"
  bytes x = 1 [(validate.rules).bytes.ipv4 = true];
  
  // x must be a valid IPv6 address in byte format
  // eg: "\x20\x01\x0D\xB8\x85\xA3\x00\x00\x00\x00\x8A\x2E\x03\x70\x73\x34"
  bytes x = 1 [(validate.rules).bytes.ipv6 = true];
  ```

## Enums

> All literal values should use the numeric (int32) value as defined in the enum descriptor.

The following examples use this `State` enum

```protobuf
enum State {
  INACTIVE = 0;
  PENDING  = 1;
  ACTIVE   = 2;
}
```

- **const**: the field must be _exactly_ the specified value.

  ```protobuf
  // x must be set to ACTIVE (2)
  State x = 1 [(validate.rules).enum.const = 2];
  ```

- **defined_only**: the field must be one of the specified values in the enum descriptor.

  ```protobuf
  // x can only be INACTIVE, PENDING, or ACTIVE
  State x = 1 [(validate.rules).enum.defined_only = true];
  ```

- **in/not_in**: these two rules permit specifying allow/denylists for the values of a field.

  ```protobuf
  // x must be either INACTIVE (0) or ACTIVE (2)
  State x = 1 [(validate.rules).enum = {in: [0,2]}];
  
  // x cannot be PENDING (1)
  State x = 1 [(validate.rules).enum = {not_in: [1]}];
  ```

## Messages

> If a field contains a message and the message has been generated with PGV, validation will be performed recursively. Message's not generated with PGV are skipped.

```protobuf
// if Person was generated with PGV and x is set,
// x's fields will be validated.
Person x = 1;
```

- **skip**: this rule specifies that the validation rules of this field should not be evaluated.

  ```protobuf
  // The fields on Person x will not be validated.
  Person x = 1 [(validate.rules).message.skip = true];
  ```

- **required**: this rule specifies that the field cannot be unset.

  ```protobuf
  // x cannot be unset
  Person x = 1 [(validate.rules).message.required = true];
  
  // x cannot be unset, but the validations on x will not be performed
  Person x = 1 [(validate.rules).message = {required: true, skip: true}];
  ```

## Repeated

- **min_items/max_items**: these rules control how many elements are contained in the field

  ```protobuf
  // x must contain at least 3 elements
  repeated int32 x = 1 [(validate.rules).repeated.min_items = 3];

  // x must contain between 5 and 10 Persons, inclusive
  repeated Person x = 1 [(validate.rules).repeated = {min_items: 5, max_items: 10}];

  // x must contain exactly 7 elements
  repeated double x = 1 [(validate.rules).repeated = {min_items: 7, max_items: 7}];
  ```

- **unique**: this rule requires that all elements in the field must be unique. This rule does not support repeated messages.

  ```protobuf
  // x must contain unique int64 values
  repeated int64 x = 1 [(validate.rules).repeated.unique = true];
  ```

- **items**: this rule specifies constraints that should be applied to each element in the field. Repeated message fields also have their validation rules applied unless `skip` is specified on this constraint.

  ```protobuf
  // x must contain positive float values
  repeated float x = 1 [(validate.rules).repeated.items.float.gt = 0];

  // x must contain Persons but don't validate them
  repeated Person x = 1 [(validate.rules).repeated.items.message.skip = true];
  ```

- **ignore_empty**: this rule specifies that if field is empty or set to the default value, to ignore any validation rules. These are typically useful where being able to unset a field in an update request, or to skip validation for optional fields where switching to WKTs is not feasible.

  ```protobuf
  repeated int64 x = 1 [(validate.rules).repeated = {ignore_empty: true, items: {int64: {gt: 200}}}];
  ```

## Maps

- **min_pairs/max_pairs**: these rules control how many KV pairs are contained in this field

  ```protobuf
  // x must contain at most 3 KV pairs
  map<string, uint64> x = 1 [(validate.rules).map.min_pairs = 3];

  // x must contain between 5 and 10 KV pairs
  map<string, string> x = 1 [(validate.rules)].map = {min_pairs: 5, max_pairs: 10}];

  // x must contain exactly 7 KV pairs
  map<string, Person> x = 1 [(validate.rules)].map = {min_pairs: 7, max_pairs: 7}];
  ```

- **no_sparse**: for map fields with message values, setting this rule to true disallows keys with unset values.

  ```protobuf
  // all values in x must be set
  map<uint64, Person> x = 1 [(validate.rules).map.no_sparse = true];
  ```

- **keys**: this rule specifies constraints that are applied to the keys in the field.

  ```protobuf
  // x's keys must all be negative
  <sint32, string> x = [(validate.rules).map.keys.sint32.lt = 0];
  ```

- **values**: this rule specifies constraints that are be applied to each value in the field. Repeated message fields also have their validation rules applied unless `skip` is specified on this constraint.

  ```protobuf
  // x must contain strings of at least 3 characters
  map<string, string> x = 1 [(validate.rules).map.values.string.min_len = 3];

  // x must contain Persons but doesn't validate them
  map<string, Person> x = 1 [(validate.rules).map.values.message.skip = true];
  ```

- **ignore_empty**: this rule specifies that if field is empty or set to the default value, to ignore any validation rules. These are typically useful where being able to unset a field in an update request, or to skip validation for optional fields where switching to WKTs is not feasible.

  ```protobuf
  map<string, string> x = 1 [(validate.rules).map = {ignore_empty: true, values: {string: {min_len: 3}}}];
  ```

## Well-Known Types (WKTs)

A set of [WKTs][wkts] are packaged with protoc and common message patterns useful in many domains.

### Scalar Value Wrappers

In the `proto3` syntax, there is no way of distinguishing between unset and the zero value of a scalar field. The value WKTs permit this differentiation by wrapping them in a message. PGV permits using the same scalar rules that the wrapper encapsulates.

```protobuf
// if it is set, x must be greater than 3
google.protobuf.Int32Value x = 1 [(validate.rules).int32.gt = 3];
```

Message Rules can also be used with scalar Well-Known Types (WKTs):

```protobuf
// Ensures that if a value is not set for age, it would not pass the validation despite its zero value being 0.
message X { google.protobuf.Int32Value age = 1 [(validate.rules).int32.gt = -1, (validate.rules).message.required = true]; }
```

### Anys

- **required**: this rule specifies that the field must be set

  ```protobuf
  // x cannot be unset
  google.protobuf.Any x = 1 [(validate.rules).any.required = true];
  ```

- **in/not_in**: these two rules permit specifying allow/denylists for the `type_url` value in this field. Consider using a `oneof` union instead of `in` if possible.

  ```protobuf
  // x must not be the Duration or Timestamp WKT
  google.protobuf.Any x = 1 [(validate.rules).any = {not_in: [
      "type.googleapis.com/google.protobuf.Duration",
      "type.googleapis.com/google.protobuf.Timestamp"
    ]}];
  ```

### Durations

- **required**: this rule specifies that the field must be set

  ```protobuf
  // x cannot be unset
  google.protobuf.Duration x = 1 [(validate.rules).duration.required = true];
  ```

- **const**: the field must be _exactly_ the specified value.

  ```protobuf
  // x must equal 1.5s exactly
  google.protobuf.Duration x = 1 [(validate.rules).duration.const = {
      seconds: 1,
      nanos:   500000000
    }];
  ```

- **lt/lte/gt/gte**: these inequalities (`<`, `<=`, `>`, `>=`, respectively) allow for deriving ranges in which the field must reside.

  ```protobuf
  // x must be less than 10s
  google.protobuf.Duration x = 1 [(validate.rules).duration.lt.seconds = 10];

  // x must be greater than or equal to 20ns
  google.protobuf.Duration x = 1 [(validate.rules).duration.gte.nanos = 20];

  // x must be in the range [0s, 1s)
  google.protobuf.Duration x = 1 [(validate.rules).duration = {
      gte: {},
      lt:  {seconds: 1}
    }];
  ```

  Inverting the values of `lt(e)` and `gt(e)` is valid and creates an exclusive range.

  ```protobuf
  // x must be outside the range [0s, 1s)
  google.protobuf.Duration x = 1 [(validate.rules).duration = {
      lt:  {},
      gte: {seconds: 1}
    }];
  ```

- **in/not_in**: these two rules permit specifying allow/denylists for the values of a field.

  ```protobuf
  // x must be either 0s or 1s
  google.protobuf.Duration x = 1 [(validate.rules).duration = {in: [
      {},
      {seconds: 1}
    ]}];
  
  // x cannot be 20s nor 500ns
  google.protobuf.Duration x = 1 [(validate.rules).duration = {not_in: [
      {seconds: 20},
      {nanos: 500}
    ]}];
  ```

### Timestamps

- **required**: this rule specifies that the field must be set

  ```protobuf
  // x cannot be unset
  google.protobuf.Timestamp x = 1 [(validate.rules).timestamp.required = true];
  ```

- **const**: the field must be _exactly_ the specified value.

  ```protobuf
  // x must equal 2009/11/10T23:00:00.500Z exactly
  google.protobuf.Timestamp x = 1 [(validate.rules).timestamp.const = {
      seconds: 63393490800,
      nanos:   500000000
    }];
  ```

- **lt/lte/gt/gte**: these inequalities (`<`, `<=`, `>`, `>=`, respectively) allow for deriving ranges in which the field must reside.

  ```protobuf
  // x must be less than the Unix Epoch
  google.protobuf.Timestamp x = 1 [(validate.rules).timestamp.lt.seconds = 0];

  // x must be greater than or equal to 2009/11/10T23:00:00Z
  google.protobuf.Timestamp x = 1 [(validate.rules).timestamp.gte.seconds = 63393490800];

  // x must be in the range [epoch, 2009/11/10T23:00:00Z)
  google.protobuf.Timestamp x = 1 [(validate.rules).timestamp = {
      gte: {},
      lt:  {seconds: 63393490800}
    }];
  ```

  Inverting the values of `lt(e)` and `gt(e)` is valid and creates an exclusive range.

  ```protobuf
  // x must be outside the range [epoch, 2009/11/10T23:00:00Z)
  google.protobuf.Timestamp x = 1 [(validate.rules).timestamp = {
      lt:  {},
      gte: {seconds: 63393490800}
    }];
  ```

- **lt_now/gt_now**: these inequalities allow for ranges relative to the current time. These rules cannot be used with the absolute rules above.

  ```protobuf
  // x must be less than the current timestamp
  google.protobuf.Timestamp x = 1 [(validate.rules).timestamp.lt_now = true];
  ```
- **within**: this rule specifies that the field's value should be within a duration of the current time. This rule can be used in conjunction with `lt_now` and `gt_now` to control those ranges.

  ```protobuf
  // x must be within ±1s of the current time
  google.protobuf.Timestamp x = 1 [(validate.rules).timestamp.within.seconds = 1];
  
  // x must be within the range (now, now+1h)
  google.protobuf.Timestamp x = 1 [(validate.rules).timestamp = {
      gt_now: true,
      within: {seconds: 3600}
    }];
  ```

## Message-Global

- **disabled**: All validation rules for the fields on a message can be nullified, including any message fields that support validation themselves.

  ```protobuf
  message Person {
    option (validate.disabled) = true;

    // x will not be required to be greater than 123
    uint64 x = 1 [(validate.rules).uint64.gt = 123];

    // y's fields will not be validated
    Person y = 2;
  }
  ```

- **ignored**: Don't generate a validate method or any related validation code for this message.

  ```protobuf
  message Person {
    option (validate.ignored) = true;
  
    // x will not be required to be greater than 123
    uint64 x = 1 [(validate.rules).uint64.gt = 123];
  
    // y's fields will not be validated
    Person y = 2;
  }
  ```

## OneOfs

- **required**: require that one of the fields in a `oneof` must be set. By default, none or one of the unioned fields can be set. Enabling this rules disallows having all of them unset.

  ```protobuf
  oneof id {
    // either x, y, or z must be set.
    option (validate.required) = true;
  
    string x = 1;
    int32  y = 2;
    Person z = 3;
  }
  ```