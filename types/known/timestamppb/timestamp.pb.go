// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/protobuf/timestamp.proto

package timestamppb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	sync "sync"
)

const (
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 0)
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(0 - protoimpl.MinVersion)
)

// A Timestamp represents a point in time independent of any time zone or local
// calendar, encoded as a count of seconds and fractions of seconds at
// nanosecond resolution. The count is relative to an epoch at UTC midnight on
// January 1, 1970, in the proleptic Gregorian calendar which extends the
// Gregorian calendar backwards to year one.
//
// All minutes are 60 seconds long. Leap seconds are "smeared" so that no leap
// second table is needed for interpretation, using a [24-hour linear
// smear](https://developers.google.com/time/smear).
//
// The range is from 0001-01-01T00:00:00Z to 9999-12-31T23:59:59.999999999Z. By
// restricting to that range, we ensure that we can convert to and from [RFC
// 3339](https://www.ietf.org/rfc/rfc3339.txt) date strings.
//
// # Examples
//
// Example 1: Compute Timestamp from POSIX `time()`.
//
//     Timestamp timestamp;
//     timestamp.set_seconds(time(NULL));
//     timestamp.set_nanos(0);
//
// Example 2: Compute Timestamp from POSIX `gettimeofday()`.
//
//     struct timeval tv;
//     gettimeofday(&tv, NULL);
//
//     Timestamp timestamp;
//     timestamp.set_seconds(tv.tv_sec);
//     timestamp.set_nanos(tv.tv_usec * 1000);
//
// Example 3: Compute Timestamp from Win32 `GetSystemTimeAsFileTime()`.
//
//     FILETIME ft;
//     GetSystemTimeAsFileTime(&ft);
//     UINT64 ticks = (((UINT64)ft.dwHighDateTime) << 32) | ft.dwLowDateTime;
//
//     // A Windows tick is 100 nanoseconds. Windows epoch 1601-01-01T00:00:00Z
//     // is 11644473600 seconds before Unix epoch 1970-01-01T00:00:00Z.
//     Timestamp timestamp;
//     timestamp.set_seconds((INT64) ((ticks / 10000000) - 11644473600LL));
//     timestamp.set_nanos((INT32) ((ticks % 10000000) * 100));
//
// Example 4: Compute Timestamp from Java `System.currentTimeMillis()`.
//
//     long millis = System.currentTimeMillis();
//
//     Timestamp timestamp = Timestamp.newBuilder().setSeconds(millis / 1000)
//         .setNanos((int) ((millis % 1000) * 1000000)).build();
//
//
// Example 5: Compute Timestamp from current time in Python.
//
//     timestamp = Timestamp()
//     timestamp.GetCurrentTime()
//
// # JSON Mapping
//
// In JSON format, the Timestamp type is encoded as a string in the
// [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt) format. That is, the
// format is "{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z"
// where {year} is always expressed using four digits while {month}, {day},
// {hour}, {min}, and {sec} are zero-padded to two digits each. The fractional
// seconds, which can go up to 9 digits (i.e. up to 1 nanosecond resolution),
// are optional. The "Z" suffix indicates the timezone ("UTC"); the timezone
// is required. A proto3 JSON serializer should always use UTC (as indicated by
// "Z") when printing the Timestamp type and a proto3 JSON parser should be
// able to accept both UTC and other timezones (as indicated by an offset).
//
// For example, "2017-01-15T01:30:15.01Z" encodes 15.01 seconds past
// 01:30 UTC on January 15, 2017.
//
// In JavaScript, one can convert a Date object to this format using the
// standard [toISOString()](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Date/toISOString)
// method. In Python, a standard `datetime.datetime` object can be converted
// to this format using [`strftime`](https://docs.python.org/2/library/time.html#time.strftime)
// with the time format spec '%Y-%m-%dT%H:%M:%S.%fZ'. Likewise, in Java, one
// can use the Joda Time's [`ISODateTimeFormat.dateTime()`](
// http://www.joda.org/joda-time/apidocs/org/joda/time/format/ISODateTimeFormat.html#dateTime%2D%2D
// ) to obtain a formatter capable of generating timestamps in this format.
//
//
type Timestamp struct {
	state protoimpl.MessageState
	// Represents seconds of UTC time since Unix epoch
	// 1970-01-01T00:00:00Z. Must be from 0001-01-01T00:00:00Z to
	// 9999-12-31T23:59:59Z inclusive.
	Seconds int64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// Non-negative fractions of a second at nanosecond resolution. Negative
	// second values with fractions must still have non-negative nanos values
	// that count forward in time. Must be from 0 to 999,999,999
	// inclusive.
	Nanos         int32 `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Timestamp) Reset() {
	*x = Timestamp{}
}

func (x *Timestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Timestamp) ProtoMessage() {}

func (x *Timestamp) ProtoReflect() protoreflect.Message {
	mi := &file_google_protobuf_timestamp_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

func (x *Timestamp) XXX_Methods() *protoiface.Methods {
	return file_google_protobuf_timestamp_proto_msgTypes[0].Methods()
}

// Deprecated: Use Timestamp.ProtoReflect.Type instead.
func (*Timestamp) Descriptor() ([]byte, []int) {
	return file_google_protobuf_timestamp_proto_rawDescGZIP(), []int{0}
}

func (x *Timestamp) GetSeconds() int64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *Timestamp) GetNanos() int32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

var File_google_protobuf_timestamp_proto protoreflect.FileDescriptor

var file_google_protobuf_timestamp_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x22, 0x3b, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x61, 0x6e,
	0x6f, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x61, 0x6e, 0x6f, 0x73, 0x42,
	0x85, 0x01, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x42, 0x0e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x70, 0x62, 0xf8, 0x01, 0x01,
	0xa2, 0x02, 0x03, 0x47, 0x50, 0x42, 0xaa, 0x02, 0x1e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x57, 0x65, 0x6c, 0x6c, 0x4b, 0x6e, 0x6f,
	0x77, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_protobuf_timestamp_proto_rawDescOnce sync.Once
	file_google_protobuf_timestamp_proto_rawDescData = file_google_protobuf_timestamp_proto_rawDesc
)

func file_google_protobuf_timestamp_proto_rawDescGZIP() []byte {
	file_google_protobuf_timestamp_proto_rawDescOnce.Do(func() {
		file_google_protobuf_timestamp_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_protobuf_timestamp_proto_rawDescData)
	})
	return file_google_protobuf_timestamp_proto_rawDescData
}

var file_google_protobuf_timestamp_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_google_protobuf_timestamp_proto_goTypes = []interface{}{
	(*Timestamp)(nil), // 0: google.protobuf.Timestamp
}
var file_google_protobuf_timestamp_proto_depIdxs = []int32{
	0, // starting offset of method output_type sub-list
	0, // starting offset of method input_type sub-list
	0, // starting offset of extension type_name sub-list
	0, // starting offset of extension extendee sub-list
	0, // starting offset of field type_name sub-list
}

func init() { file_google_protobuf_timestamp_proto_init() }
func file_google_protobuf_timestamp_proto_init() {
	if File_google_protobuf_timestamp_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_google_protobuf_timestamp_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Timestamp); i {
			case 0:
				return &v.state
			case 3:
				return &v.sizeCache
			case 4:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			RawDescriptor: file_google_protobuf_timestamp_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_protobuf_timestamp_proto_goTypes,
		DependencyIndexes: file_google_protobuf_timestamp_proto_depIdxs,
		MessageInfos:      file_google_protobuf_timestamp_proto_msgTypes,
	}.Build()
	File_google_protobuf_timestamp_proto = out.File
	file_google_protobuf_timestamp_proto_rawDesc = nil
	file_google_protobuf_timestamp_proto_goTypes = nil
	file_google_protobuf_timestamp_proto_depIdxs = nil
}
