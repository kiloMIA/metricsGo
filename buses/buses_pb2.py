# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: buses.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x0b\x62uses.proto\x12\x05\x62uses\"\x1f\n\nBusRequest\x12\x11\n\tBusNumber\x18\x01 \x01(\x03\"2\n\x0b\x42usResponse\x12\x11\n\tlongitude\x18\x01 \x01(\x01\x12\x10\n\x08latitude\x18\x02 \x01(\x01\x32\x43\n\nBusService\x12\x35\n\nRequestBus\x12\x11.buses.BusRequest\x1a\x12.buses.BusResponse\"\x00\x42\x12Z\x10./proto/buses/pbb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'buses_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\020./proto/buses/pb'
  _BUSREQUEST._serialized_start=22
  _BUSREQUEST._serialized_end=53
  _BUSRESPONSE._serialized_start=55
  _BUSRESPONSE._serialized_end=105
  _BUSSERVICE._serialized_start=107
  _BUSSERVICE._serialized_end=174
# @@protoc_insertion_point(module_scope)