# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: denlipov/com_request_api/v1/com_request_api.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from validate import validate_pb2 as validate_dot_validate__pb2
from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='denlipov/com_request_api/v1/com_request_api.proto',
  package='denlipov.com_request_api.v1',
  syntax='proto3',
  serialized_options=b'ZGgithub.com/denlipov/com-request-api/pkg/com-request-api;com_request_api',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n1denlipov/com_request_api/v1/com_request_api.proto\x12\x1b\x64\x65nlipov.com_request_api.v1\x1a\x17validate/validate.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"a\n\x07Request\x12\x0e\n\x02id\x18\x01 \x01(\x04R\x02id\x12\x10\n\x03\x66oo\x18\x02 \x01(\x04R\x03\x66oo\x12\x34\n\x07\x63reated\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.TimestampR\x07\x63reated\"B\n\x18\x44\x65scribeRequestV1Request\x12&\n\nrequest_id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\trequestId\"W\n\x19\x44\x65scribeRequestV1Response\x12:\n\x05value\x18\x01 \x01(\x0b\x32$.denlipov.com_request_api.v1.RequestR\x05value2\xbe\x01\n\x14\x43omRequestApiService\x12\xa5\x01\n\x11\x44\x65scribeRequestV1\x12\x35.denlipov.com_request_api.v1.DescribeRequestV1Request\x1a\x36.denlipov.com_request_api.v1.DescribeRequestV1Response\"!\x82\xd3\xe4\x93\x02\x1b\x12\x19/v1/requests/{request_id}BIZGgithub.com/denlipov/com-request-api/pkg/com-request-api;com_request_apib\x06proto3'
  ,
  dependencies=[validate_dot_validate__pb2.DESCRIPTOR,google_dot_api_dot_annotations__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])




_REQUEST = _descriptor.Descriptor(
  name='Request',
  full_name='denlipov.com_request_api.v1.Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='denlipov.com_request_api.v1.Request.id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='id', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='foo', full_name='denlipov.com_request_api.v1.Request.foo', index=1,
      number=2, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='foo', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='created', full_name='denlipov.com_request_api.v1.Request.created', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='created', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=170,
  serialized_end=267,
)


_DESCRIBEREQUESTV1REQUEST = _descriptor.Descriptor(
  name='DescribeRequestV1Request',
  full_name='denlipov.com_request_api.v1.DescribeRequestV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='request_id', full_name='denlipov.com_request_api.v1.DescribeRequestV1Request.request_id', index=0,
      number=1, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\0042\002 \000', json_name='requestId', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=269,
  serialized_end=335,
)


_DESCRIBEREQUESTV1RESPONSE = _descriptor.Descriptor(
  name='DescribeRequestV1Response',
  full_name='denlipov.com_request_api.v1.DescribeRequestV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='denlipov.com_request_api.v1.DescribeRequestV1Response.value', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='value', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=337,
  serialized_end=424,
)

_REQUEST.fields_by_name['created'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_DESCRIBEREQUESTV1RESPONSE.fields_by_name['value'].message_type = _REQUEST
DESCRIPTOR.message_types_by_name['Request'] = _REQUEST
DESCRIPTOR.message_types_by_name['DescribeRequestV1Request'] = _DESCRIBEREQUESTV1REQUEST
DESCRIPTOR.message_types_by_name['DescribeRequestV1Response'] = _DESCRIBEREQUESTV1RESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Request = _reflection.GeneratedProtocolMessageType('Request', (_message.Message,), {
  'DESCRIPTOR' : _REQUEST,
  '__module__' : 'denlipov.com_request_api.v1.com_request_api_pb2'
  # @@protoc_insertion_point(class_scope:denlipov.com_request_api.v1.Request)
  })
_sym_db.RegisterMessage(Request)

DescribeRequestV1Request = _reflection.GeneratedProtocolMessageType('DescribeRequestV1Request', (_message.Message,), {
  'DESCRIPTOR' : _DESCRIBEREQUESTV1REQUEST,
  '__module__' : 'denlipov.com_request_api.v1.com_request_api_pb2'
  # @@protoc_insertion_point(class_scope:denlipov.com_request_api.v1.DescribeRequestV1Request)
  })
_sym_db.RegisterMessage(DescribeRequestV1Request)

DescribeRequestV1Response = _reflection.GeneratedProtocolMessageType('DescribeRequestV1Response', (_message.Message,), {
  'DESCRIPTOR' : _DESCRIBEREQUESTV1RESPONSE,
  '__module__' : 'denlipov.com_request_api.v1.com_request_api_pb2'
  # @@protoc_insertion_point(class_scope:denlipov.com_request_api.v1.DescribeRequestV1Response)
  })
_sym_db.RegisterMessage(DescribeRequestV1Response)


DESCRIPTOR._options = None
_DESCRIBEREQUESTV1REQUEST.fields_by_name['request_id']._options = None

_COMREQUESTAPISERVICE = _descriptor.ServiceDescriptor(
  name='ComRequestApiService',
  full_name='denlipov.com_request_api.v1.ComRequestApiService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=427,
  serialized_end=617,
  methods=[
  _descriptor.MethodDescriptor(
    name='DescribeRequestV1',
    full_name='denlipov.com_request_api.v1.ComRequestApiService.DescribeRequestV1',
    index=0,
    containing_service=None,
    input_type=_DESCRIBEREQUESTV1REQUEST,
    output_type=_DESCRIBEREQUESTV1RESPONSE,
    serialized_options=b'\202\323\344\223\002\033\022\031/v1/requests/{request_id}',
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_COMREQUESTAPISERVICE)

DESCRIPTOR.services_by_name['ComRequestApiService'] = _COMREQUESTAPISERVICE

# @@protoc_insertion_point(module_scope)
