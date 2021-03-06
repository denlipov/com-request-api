# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ozonmp/com_request_api/v1/com_request_api.proto
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
  name='ozonmp/com_request_api/v1/com_request_api.proto',
  package='denlipov.com_request_api.v1',
  syntax='proto3',
  serialized_options=b'ZGgithub.com/denlipov/com-request-api/pkg/com-request-api;com_request_api',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n/ozonmp/com_request_api/v1/com_request_api.proto\x12\x1b\x64\x65nlipov.com_request_api.v1\x1a\x17validate/validate.proto\x1a\x1cgoogle/api/annotations.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"v\n\x07Request\x12\x0e\n\x02id\x18\x01 \x01(\x04R\x02id\x12!\n\x07service\x18\x02 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x02R\x07service\x12\x1b\n\x04user\x18\x03 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x02R\x04user\x12\x1b\n\x04text\x18\x04 \x01(\tB\x07\xfa\x42\x04r\x02\x10\x02R\x04text\"B\n\x18\x44\x65scribeRequestV1Request\x12&\n\nrequest_id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\trequestId\"W\n\x19\x44\x65scribeRequestV1Response\x12:\n\x05value\x18\x01 \x01(\x0b\x32$.denlipov.com_request_api.v1.RequestR\x05value\"X\n\x16\x43reateRequestV1Request\x12>\n\x07request\x18\x01 \x01(\x0b\x32$.denlipov.com_request_api.v1.RequestR\x07request\"A\n\x17\x43reateRequestV1Response\x12&\n\nrequest_id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\trequestId\"\x16\n\x14ListRequestV1Request\"W\n\x15ListRequestV1Response\x12>\n\x07request\x18\x01 \x03(\x0b\x32$.denlipov.com_request_api.v1.RequestR\x07request\"@\n\x16RemoveRequestV1Request\x12&\n\nrequest_id\x18\x01 \x01(\x04\x42\x07\xfa\x42\x04\x32\x02 \x00R\trequestId\"1\n\x17RemoveRequestV1Response\x12\x16\n\x06status\x18\x01 \x01(\x08R\x06status2\x9a\x05\n\x14\x43omRequestApiService\x12\xa5\x01\n\x11\x44\x65scribeRequestV1\x12\x35.denlipov.com_request_api.v1.DescribeRequestV1Request\x1a\x36.denlipov.com_request_api.v1.DescribeRequestV1Response\"!\x82\xd3\xe4\x93\x02\x1b\x12\x19/v1/requests/{request_id}\x12\x9c\x01\n\x0f\x43reateRequestV1\x12\x33.denlipov.com_request_api.v1.CreateRequestV1Request\x1a\x34.denlipov.com_request_api.v1.CreateRequestV1Response\"\x1e\x82\xd3\xe4\x93\x02\x18\"\x13/v1/requests/create:\x01*\x12\x91\x01\n\rListRequestV1\x12\x31.denlipov.com_request_api.v1.ListRequestV1Request\x1a\x32.denlipov.com_request_api.v1.ListRequestV1Response\"\x19\x82\xd3\xe4\x93\x02\x13\x12\x11/v1/requests/list\x12\xa6\x01\n\x0fRemoveRequestV1\x12\x33.denlipov.com_request_api.v1.RemoveRequestV1Request\x1a\x34.denlipov.com_request_api.v1.RemoveRequestV1Response\"(\x82\xd3\xe4\x93\x02\"* /v1/requests/remove/{request_id}BIZGgithub.com/denlipov/com-request-api/pkg/com-request-api;com_request_apib\x06proto3'
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
      name='service', full_name='denlipov.com_request_api.v1.Request.service', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\004r\002\020\002', json_name='service', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='user', full_name='denlipov.com_request_api.v1.Request.user', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\004r\002\020\002', json_name='user', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='text', full_name='denlipov.com_request_api.v1.Request.text', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=b'\372B\004r\002\020\002', json_name='text', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=168,
  serialized_end=286,
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
  serialized_start=288,
  serialized_end=354,
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
  serialized_start=356,
  serialized_end=443,
)


_CREATEREQUESTV1REQUEST = _descriptor.Descriptor(
  name='CreateRequestV1Request',
  full_name='denlipov.com_request_api.v1.CreateRequestV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='request', full_name='denlipov.com_request_api.v1.CreateRequestV1Request.request', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='request', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=445,
  serialized_end=533,
)


_CREATEREQUESTV1RESPONSE = _descriptor.Descriptor(
  name='CreateRequestV1Response',
  full_name='denlipov.com_request_api.v1.CreateRequestV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='request_id', full_name='denlipov.com_request_api.v1.CreateRequestV1Response.request_id', index=0,
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
  serialized_start=535,
  serialized_end=600,
)


_LISTREQUESTV1REQUEST = _descriptor.Descriptor(
  name='ListRequestV1Request',
  full_name='denlipov.com_request_api.v1.ListRequestV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
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
  serialized_start=602,
  serialized_end=624,
)


_LISTREQUESTV1RESPONSE = _descriptor.Descriptor(
  name='ListRequestV1Response',
  full_name='denlipov.com_request_api.v1.ListRequestV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='request', full_name='denlipov.com_request_api.v1.ListRequestV1Response.request', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='request', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=626,
  serialized_end=713,
)


_REMOVEREQUESTV1REQUEST = _descriptor.Descriptor(
  name='RemoveRequestV1Request',
  full_name='denlipov.com_request_api.v1.RemoveRequestV1Request',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='request_id', full_name='denlipov.com_request_api.v1.RemoveRequestV1Request.request_id', index=0,
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
  serialized_start=715,
  serialized_end=779,
)


_REMOVEREQUESTV1RESPONSE = _descriptor.Descriptor(
  name='RemoveRequestV1Response',
  full_name='denlipov.com_request_api.v1.RemoveRequestV1Response',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='denlipov.com_request_api.v1.RemoveRequestV1Response.status', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='status', file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=781,
  serialized_end=830,
)

_DESCRIBEREQUESTV1RESPONSE.fields_by_name['value'].message_type = _REQUEST
_CREATEREQUESTV1REQUEST.fields_by_name['request'].message_type = _REQUEST
_LISTREQUESTV1RESPONSE.fields_by_name['request'].message_type = _REQUEST
DESCRIPTOR.message_types_by_name['Request'] = _REQUEST
DESCRIPTOR.message_types_by_name['DescribeRequestV1Request'] = _DESCRIBEREQUESTV1REQUEST
DESCRIPTOR.message_types_by_name['DescribeRequestV1Response'] = _DESCRIBEREQUESTV1RESPONSE
DESCRIPTOR.message_types_by_name['CreateRequestV1Request'] = _CREATEREQUESTV1REQUEST
DESCRIPTOR.message_types_by_name['CreateRequestV1Response'] = _CREATEREQUESTV1RESPONSE
DESCRIPTOR.message_types_by_name['ListRequestV1Request'] = _LISTREQUESTV1REQUEST
DESCRIPTOR.message_types_by_name['ListRequestV1Response'] = _LISTREQUESTV1RESPONSE
DESCRIPTOR.message_types_by_name['RemoveRequestV1Request'] = _REMOVEREQUESTV1REQUEST
DESCRIPTOR.message_types_by_name['RemoveRequestV1Response'] = _REMOVEREQUESTV1RESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Request = _reflection.GeneratedProtocolMessageType('Request', (_message.Message,), {
  'DESCRIPTOR' : _REQUEST,
  '__module__' : 'ozonmp.com_request_api.v1.com_request_api_pb2'
  # @@protoc_insertion_point(class_scope:denlipov.com_request_api.v1.Request)
  })
_sym_db.RegisterMessage(Request)

DescribeRequestV1Request = _reflection.GeneratedProtocolMessageType('DescribeRequestV1Request', (_message.Message,), {
  'DESCRIPTOR' : _DESCRIBEREQUESTV1REQUEST,
  '__module__' : 'ozonmp.com_request_api.v1.com_request_api_pb2'
  # @@protoc_insertion_point(class_scope:denlipov.com_request_api.v1.DescribeRequestV1Request)
  })
_sym_db.RegisterMessage(DescribeRequestV1Request)

DescribeRequestV1Response = _reflection.GeneratedProtocolMessageType('DescribeRequestV1Response', (_message.Message,), {
  'DESCRIPTOR' : _DESCRIBEREQUESTV1RESPONSE,
  '__module__' : 'ozonmp.com_request_api.v1.com_request_api_pb2'
  # @@protoc_insertion_point(class_scope:denlipov.com_request_api.v1.DescribeRequestV1Response)
  })
_sym_db.RegisterMessage(DescribeRequestV1Response)

CreateRequestV1Request = _reflection.GeneratedProtocolMessageType('CreateRequestV1Request', (_message.Message,), {
  'DESCRIPTOR' : _CREATEREQUESTV1REQUEST,
  '__module__' : 'ozonmp.com_request_api.v1.com_request_api_pb2'
  # @@protoc_insertion_point(class_scope:denlipov.com_request_api.v1.CreateRequestV1Request)
  })
_sym_db.RegisterMessage(CreateRequestV1Request)

CreateRequestV1Response = _reflection.GeneratedProtocolMessageType('CreateRequestV1Response', (_message.Message,), {
  'DESCRIPTOR' : _CREATEREQUESTV1RESPONSE,
  '__module__' : 'ozonmp.com_request_api.v1.com_request_api_pb2'
  # @@protoc_insertion_point(class_scope:denlipov.com_request_api.v1.CreateRequestV1Response)
  })
_sym_db.RegisterMessage(CreateRequestV1Response)

ListRequestV1Request = _reflection.GeneratedProtocolMessageType('ListRequestV1Request', (_message.Message,), {
  'DESCRIPTOR' : _LISTREQUESTV1REQUEST,
  '__module__' : 'ozonmp.com_request_api.v1.com_request_api_pb2'
  # @@protoc_insertion_point(class_scope:denlipov.com_request_api.v1.ListRequestV1Request)
  })
_sym_db.RegisterMessage(ListRequestV1Request)

ListRequestV1Response = _reflection.GeneratedProtocolMessageType('ListRequestV1Response', (_message.Message,), {
  'DESCRIPTOR' : _LISTREQUESTV1RESPONSE,
  '__module__' : 'ozonmp.com_request_api.v1.com_request_api_pb2'
  # @@protoc_insertion_point(class_scope:denlipov.com_request_api.v1.ListRequestV1Response)
  })
_sym_db.RegisterMessage(ListRequestV1Response)

RemoveRequestV1Request = _reflection.GeneratedProtocolMessageType('RemoveRequestV1Request', (_message.Message,), {
  'DESCRIPTOR' : _REMOVEREQUESTV1REQUEST,
  '__module__' : 'ozonmp.com_request_api.v1.com_request_api_pb2'
  # @@protoc_insertion_point(class_scope:denlipov.com_request_api.v1.RemoveRequestV1Request)
  })
_sym_db.RegisterMessage(RemoveRequestV1Request)

RemoveRequestV1Response = _reflection.GeneratedProtocolMessageType('RemoveRequestV1Response', (_message.Message,), {
  'DESCRIPTOR' : _REMOVEREQUESTV1RESPONSE,
  '__module__' : 'ozonmp.com_request_api.v1.com_request_api_pb2'
  # @@protoc_insertion_point(class_scope:denlipov.com_request_api.v1.RemoveRequestV1Response)
  })
_sym_db.RegisterMessage(RemoveRequestV1Response)


DESCRIPTOR._options = None
_REQUEST.fields_by_name['service']._options = None
_REQUEST.fields_by_name['user']._options = None
_REQUEST.fields_by_name['text']._options = None
_DESCRIBEREQUESTV1REQUEST.fields_by_name['request_id']._options = None
_CREATEREQUESTV1RESPONSE.fields_by_name['request_id']._options = None
_REMOVEREQUESTV1REQUEST.fields_by_name['request_id']._options = None

_COMREQUESTAPISERVICE = _descriptor.ServiceDescriptor(
  name='ComRequestApiService',
  full_name='denlipov.com_request_api.v1.ComRequestApiService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=833,
  serialized_end=1499,
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
  _descriptor.MethodDescriptor(
    name='CreateRequestV1',
    full_name='denlipov.com_request_api.v1.ComRequestApiService.CreateRequestV1',
    index=1,
    containing_service=None,
    input_type=_CREATEREQUESTV1REQUEST,
    output_type=_CREATEREQUESTV1RESPONSE,
    serialized_options=b'\202\323\344\223\002\030\"\023/v1/requests/create:\001*',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='ListRequestV1',
    full_name='denlipov.com_request_api.v1.ComRequestApiService.ListRequestV1',
    index=2,
    containing_service=None,
    input_type=_LISTREQUESTV1REQUEST,
    output_type=_LISTREQUESTV1RESPONSE,
    serialized_options=b'\202\323\344\223\002\023\022\021/v1/requests/list',
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='RemoveRequestV1',
    full_name='denlipov.com_request_api.v1.ComRequestApiService.RemoveRequestV1',
    index=3,
    containing_service=None,
    input_type=_REMOVEREQUESTV1REQUEST,
    output_type=_REMOVEREQUESTV1RESPONSE,
    serialized_options=b'\202\323\344\223\002\"* /v1/requests/remove/{request_id}',
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_COMREQUESTAPISERVICE)

DESCRIPTOR.services_by_name['ComRequestApiService'] = _COMREQUESTAPISERVICE

# @@protoc_insertion_point(module_scope)
