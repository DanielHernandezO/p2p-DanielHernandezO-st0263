# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: search_file.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x11search_file.proto\"U\n\rFileRequested\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x13\n\x0bip_response\x18\x03 \x01(\t\x12\x15\n\rport_response\x18\x04 \x01(\t\"\x1b\n\x08Response\x12\x0f\n\x07message\x18\x01 \x01(\t\"\'\n\x0bJoinRequest\x12\n\n\x02ip\x18\x01 \x01(\t\x12\x0c\n\x04port\x18\x02 \x01(\t\"\x87\x01\n\x10JoinPeersRequest\x12\n\n\x02ip\x18\x01 \x01(\t\x12\x0c\n\x04port\x18\x02 \x01(\t\x12+\n\x05peers\x18\x03 \x03(\x0b\x32\x1c.JoinPeersRequest.PeersEntry\x1a,\n\nPeersEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"W\n\x05Peers\x12 \n\x05peers\x18\x01 \x03(\x0b\x32\x11.Peers.PeersEntry\x1a,\n\nPeersEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\r\n\x0bVoidRequest21\n\nSearchFile\x12#\n\x06Search\x12\x0e.FileRequested\x1a\t.Response2T\n\x04Join\x12+\n\x0b\x45xecuteJoin\x12\x11.JoinPeersRequest\x1a\t.Response\x12\x1f\n\x04Join\x12\x0c.JoinRequest\x1a\t.Response2L\n\x05\x46\x65tch\x12$\n\x0c\x45xecuteFetch\x12\x0c.VoidRequest\x1a\x06.Peers\x12\x1d\n\x05\x46\x65tch\x12\x0c.VoidRequest\x1a\x06.Peersb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'search_file_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  _globals['_JOINPEERSREQUEST_PEERSENTRY']._options = None
  _globals['_JOINPEERSREQUEST_PEERSENTRY']._serialized_options = b'8\001'
  _globals['_PEERS_PEERSENTRY']._options = None
  _globals['_PEERS_PEERSENTRY']._serialized_options = b'8\001'
  _globals['_FILEREQUESTED']._serialized_start=21
  _globals['_FILEREQUESTED']._serialized_end=106
  _globals['_RESPONSE']._serialized_start=108
  _globals['_RESPONSE']._serialized_end=135
  _globals['_JOINREQUEST']._serialized_start=137
  _globals['_JOINREQUEST']._serialized_end=176
  _globals['_JOINPEERSREQUEST']._serialized_start=179
  _globals['_JOINPEERSREQUEST']._serialized_end=314
  _globals['_JOINPEERSREQUEST_PEERSENTRY']._serialized_start=270
  _globals['_JOINPEERSREQUEST_PEERSENTRY']._serialized_end=314
  _globals['_PEERS']._serialized_start=316
  _globals['_PEERS']._serialized_end=403
  _globals['_PEERS_PEERSENTRY']._serialized_start=270
  _globals['_PEERS_PEERSENTRY']._serialized_end=314
  _globals['_VOIDREQUEST']._serialized_start=405
  _globals['_VOIDREQUEST']._serialized_end=418
  _globals['_SEARCHFILE']._serialized_start=420
  _globals['_SEARCHFILE']._serialized_end=469
  _globals['_JOIN']._serialized_start=471
  _globals['_JOIN']._serialized_end=555
  _globals['_FETCH']._serialized_start=557
  _globals['_FETCH']._serialized_end=633
# @@protoc_insertion_point(module_scope)