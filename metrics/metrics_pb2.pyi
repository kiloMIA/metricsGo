from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class PollutionRequest(_message.Message):
    __slots__ = ["city", "type"]
    CITY_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    city: int
    type: str
    def __init__(self, city: _Optional[int] = ..., type: _Optional[str] = ...) -> None: ...

class PollutionResponse(_message.Message):
    __slots__ = ["city", "co2", "district", "pm25"]
    CITY_FIELD_NUMBER: _ClassVar[int]
    CO2_FIELD_NUMBER: _ClassVar[int]
    DISTRICT_FIELD_NUMBER: _ClassVar[int]
    PM25_FIELD_NUMBER: _ClassVar[int]
    city: int
    co2: int
    district: str
    pm25: int
    def __init__(self, city: _Optional[int] = ..., district: _Optional[str] = ..., co2: _Optional[int] = ..., pm25: _Optional[int] = ...) -> None: ...

class TemperatureRequest(_message.Message):
    __slots__ = ["city", "type"]
    CITY_FIELD_NUMBER: _ClassVar[int]
    TYPE_FIELD_NUMBER: _ClassVar[int]
    city: int
    type: str
    def __init__(self, city: _Optional[int] = ..., type: _Optional[str] = ...) -> None: ...

class TemperatureResponse(_message.Message):
    __slots__ = ["city", "district", "humidity", "temperature"]
    CITY_FIELD_NUMBER: _ClassVar[int]
    DISTRICT_FIELD_NUMBER: _ClassVar[int]
    HUMIDITY_FIELD_NUMBER: _ClassVar[int]
    TEMPERATURE_FIELD_NUMBER: _ClassVar[int]
    city: int
    district: str
    humidity: int
    temperature: int
    def __init__(self, city: _Optional[int] = ..., district: _Optional[str] = ..., temperature: _Optional[int] = ..., humidity: _Optional[int] = ...) -> None: ...
