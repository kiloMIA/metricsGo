from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class BusRequest(_message.Message):
    __slots__ = ["BusNumber"]
    BUSNUMBER_FIELD_NUMBER: _ClassVar[int]
    BusNumber: int
    def __init__(self, BusNumber: _Optional[int] = ...) -> None: ...

class BusResponse(_message.Message):
    __slots__ = ["latitude", "longitude"]
    LATITUDE_FIELD_NUMBER: _ClassVar[int]
    LONGITUDE_FIELD_NUMBER: _ClassVar[int]
    latitude: float
    longitude: float
    def __init__(self, longitude: _Optional[float] = ..., latitude: _Optional[float] = ...) -> None: ...
