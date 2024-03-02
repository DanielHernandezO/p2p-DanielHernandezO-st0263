from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class FileRequested(_message.Message):
    __slots__ = ("id", "name", "ip_response", "port_response")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    IP_RESPONSE_FIELD_NUMBER: _ClassVar[int]
    PORT_RESPONSE_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    ip_response: str
    port_response: str
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., ip_response: _Optional[str] = ..., port_response: _Optional[str] = ...) -> None: ...

class Response(_message.Message):
    __slots__ = ("message",)
    MESSAGE_FIELD_NUMBER: _ClassVar[int]
    message: str
    def __init__(self, message: _Optional[str] = ...) -> None: ...

class JoinRequest(_message.Message):
    __slots__ = ("ip", "port")
    IP_FIELD_NUMBER: _ClassVar[int]
    PORT_FIELD_NUMBER: _ClassVar[int]
    ip: str
    port: str
    def __init__(self, ip: _Optional[str] = ..., port: _Optional[str] = ...) -> None: ...

class JoinPeersRequest(_message.Message):
    __slots__ = ("ip", "port", "peers")
    class PeersEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    IP_FIELD_NUMBER: _ClassVar[int]
    PORT_FIELD_NUMBER: _ClassVar[int]
    PEERS_FIELD_NUMBER: _ClassVar[int]
    ip: str
    port: str
    peers: _containers.ScalarMap[str, str]
    def __init__(self, ip: _Optional[str] = ..., port: _Optional[str] = ..., peers: _Optional[_Mapping[str, str]] = ...) -> None: ...

class Peers(_message.Message):
    __slots__ = ("peers",)
    class PeersEntry(_message.Message):
        __slots__ = ("key", "value")
        KEY_FIELD_NUMBER: _ClassVar[int]
        VALUE_FIELD_NUMBER: _ClassVar[int]
        key: str
        value: str
        def __init__(self, key: _Optional[str] = ..., value: _Optional[str] = ...) -> None: ...
    PEERS_FIELD_NUMBER: _ClassVar[int]
    peers: _containers.ScalarMap[str, str]
    def __init__(self, peers: _Optional[_Mapping[str, str]] = ...) -> None: ...

class VoidRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
