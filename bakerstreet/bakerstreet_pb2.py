# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: bakerstreet.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='bakerstreet.proto',
  package='com.appknox.bakerstreet',
  syntax='proto3',
  serialized_pb=_b('\n\x11\x62\x61kerstreet.proto\x12\x17\x63om.appknox.bakerstreet\"-\n\x07Message\x12\r\n\x05Title\x18\x01 \x01(\t\x12\x13\n\x0b\x44\x65scription\x18\x02 \x01(\t\"\x19\n\x08Packages\x12\r\n\x05Names\x18\x01 \x03(\t\"S\n\x06\x44\x65vice\x12\x0c\n\x04Uuid\x18\x01 \x01(\t\x12\x10\n\x08IsTablet\x18\x02 \x01(\x08\x12\x10\n\x08Platform\x18\x03 \x01(\x05\x12\x17\n\x0fPlatformVersion\x18\x04 \x01(\t\"-\n\x07\x46inding\x12\r\n\x05Title\x18\x01 \x01(\t\x12\x13\n\x0b\x44\x65scription\x18\x02 \x01(\t2\xc5\x04\n\x08Moriarty\x12I\n\x04Info\x12 .com.appknox.bakerstreet.Message\x1a\x1f.com.appknox.bakerstreet.Device\x12J\n\x04\x45\x63ho\x12 .com.appknox.bakerstreet.Message\x1a .com.appknox.bakerstreet.Message\x12O\n\tLaunchApp\x12 .com.appknox.bakerstreet.Message\x1a .com.appknox.bakerstreet.Message\x12Q\n\x0bHealthCheck\x12 .com.appknox.bakerstreet.Message\x1a .com.appknox.bakerstreet.Message\x12S\n\rRemovePackage\x12 .com.appknox.bakerstreet.Message\x1a .com.appknox.bakerstreet.Message\x12T\n\x0eInstallPackage\x12 .com.appknox.bakerstreet.Message\x1a .com.appknox.bakerstreet.Message\x12S\n\x0cListPackages\x12 .com.appknox.bakerstreet.Message\x1a!.com.appknox.bakerstreet.Packages2\xa7\x01\n\x07Mycroft\x12I\n\x04Poll\x12\x1f.com.appknox.bakerstreet.Device\x1a .com.appknox.bakerstreet.Message\x12Q\n\x0bPushFinding\x12 .com.appknox.bakerstreet.Finding\x1a .com.appknox.bakerstreet.Messageb\x06proto3')
)




_MESSAGE = _descriptor.Descriptor(
  name='Message',
  full_name='com.appknox.bakerstreet.Message',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Title', full_name='com.appknox.bakerstreet.Message.Title', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Description', full_name='com.appknox.bakerstreet.Message.Description', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=46,
  serialized_end=91,
)


_PACKAGES = _descriptor.Descriptor(
  name='Packages',
  full_name='com.appknox.bakerstreet.Packages',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Names', full_name='com.appknox.bakerstreet.Packages.Names', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=93,
  serialized_end=118,
)


_DEVICE = _descriptor.Descriptor(
  name='Device',
  full_name='com.appknox.bakerstreet.Device',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Uuid', full_name='com.appknox.bakerstreet.Device.Uuid', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='IsTablet', full_name='com.appknox.bakerstreet.Device.IsTablet', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Platform', full_name='com.appknox.bakerstreet.Device.Platform', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='PlatformVersion', full_name='com.appknox.bakerstreet.Device.PlatformVersion', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=120,
  serialized_end=203,
)


_FINDING = _descriptor.Descriptor(
  name='Finding',
  full_name='com.appknox.bakerstreet.Finding',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Title', full_name='com.appknox.bakerstreet.Finding.Title', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='Description', full_name='com.appknox.bakerstreet.Finding.Description', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=205,
  serialized_end=250,
)

DESCRIPTOR.message_types_by_name['Message'] = _MESSAGE
DESCRIPTOR.message_types_by_name['Packages'] = _PACKAGES
DESCRIPTOR.message_types_by_name['Device'] = _DEVICE
DESCRIPTOR.message_types_by_name['Finding'] = _FINDING
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Message = _reflection.GeneratedProtocolMessageType('Message', (_message.Message,), dict(
  DESCRIPTOR = _MESSAGE,
  __module__ = 'bakerstreet_pb2'
  # @@protoc_insertion_point(class_scope:com.appknox.bakerstreet.Message)
  ))
_sym_db.RegisterMessage(Message)

Packages = _reflection.GeneratedProtocolMessageType('Packages', (_message.Message,), dict(
  DESCRIPTOR = _PACKAGES,
  __module__ = 'bakerstreet_pb2'
  # @@protoc_insertion_point(class_scope:com.appknox.bakerstreet.Packages)
  ))
_sym_db.RegisterMessage(Packages)

Device = _reflection.GeneratedProtocolMessageType('Device', (_message.Message,), dict(
  DESCRIPTOR = _DEVICE,
  __module__ = 'bakerstreet_pb2'
  # @@protoc_insertion_point(class_scope:com.appknox.bakerstreet.Device)
  ))
_sym_db.RegisterMessage(Device)

Finding = _reflection.GeneratedProtocolMessageType('Finding', (_message.Message,), dict(
  DESCRIPTOR = _FINDING,
  __module__ = 'bakerstreet_pb2'
  # @@protoc_insertion_point(class_scope:com.appknox.bakerstreet.Finding)
  ))
_sym_db.RegisterMessage(Finding)


try:
  # THESE ELEMENTS WILL BE DEPRECATED.
  # Please use the generated *_pb2_grpc.py files instead.
  import grpc
  from grpc.beta import implementations as beta_implementations
  from grpc.beta import interfaces as beta_interfaces
  from grpc.framework.common import cardinality
  from grpc.framework.interfaces.face import utilities as face_utilities


  class MoriartyStub(object):
    # missing associated documentation comment in .proto file
    pass

    def __init__(self, channel):
      """Constructor.

      Args:
        channel: A grpc.Channel.
      """
      self.Info = channel.unary_unary(
          '/com.appknox.bakerstreet.Moriarty/Info',
          request_serializer=Message.SerializeToString,
          response_deserializer=Device.FromString,
          )
      self.Echo = channel.unary_unary(
          '/com.appknox.bakerstreet.Moriarty/Echo',
          request_serializer=Message.SerializeToString,
          response_deserializer=Message.FromString,
          )
      self.LaunchApp = channel.unary_unary(
          '/com.appknox.bakerstreet.Moriarty/LaunchApp',
          request_serializer=Message.SerializeToString,
          response_deserializer=Message.FromString,
          )
      self.HealthCheck = channel.unary_unary(
          '/com.appknox.bakerstreet.Moriarty/HealthCheck',
          request_serializer=Message.SerializeToString,
          response_deserializer=Message.FromString,
          )
      self.RemovePackage = channel.unary_unary(
          '/com.appknox.bakerstreet.Moriarty/RemovePackage',
          request_serializer=Message.SerializeToString,
          response_deserializer=Message.FromString,
          )
      self.InstallPackage = channel.unary_unary(
          '/com.appknox.bakerstreet.Moriarty/InstallPackage',
          request_serializer=Message.SerializeToString,
          response_deserializer=Message.FromString,
          )
      self.ListPackages = channel.unary_unary(
          '/com.appknox.bakerstreet.Moriarty/ListPackages',
          request_serializer=Message.SerializeToString,
          response_deserializer=Packages.FromString,
          )


  class MoriartyServicer(object):
    # missing associated documentation comment in .proto file
    pass

    def Info(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def Echo(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def LaunchApp(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def HealthCheck(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def RemovePackage(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def InstallPackage(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def ListPackages(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')


  def add_MoriartyServicer_to_server(servicer, server):
    rpc_method_handlers = {
        'Info': grpc.unary_unary_rpc_method_handler(
            servicer.Info,
            request_deserializer=Message.FromString,
            response_serializer=Device.SerializeToString,
        ),
        'Echo': grpc.unary_unary_rpc_method_handler(
            servicer.Echo,
            request_deserializer=Message.FromString,
            response_serializer=Message.SerializeToString,
        ),
        'LaunchApp': grpc.unary_unary_rpc_method_handler(
            servicer.LaunchApp,
            request_deserializer=Message.FromString,
            response_serializer=Message.SerializeToString,
        ),
        'HealthCheck': grpc.unary_unary_rpc_method_handler(
            servicer.HealthCheck,
            request_deserializer=Message.FromString,
            response_serializer=Message.SerializeToString,
        ),
        'RemovePackage': grpc.unary_unary_rpc_method_handler(
            servicer.RemovePackage,
            request_deserializer=Message.FromString,
            response_serializer=Message.SerializeToString,
        ),
        'InstallPackage': grpc.unary_unary_rpc_method_handler(
            servicer.InstallPackage,
            request_deserializer=Message.FromString,
            response_serializer=Message.SerializeToString,
        ),
        'ListPackages': grpc.unary_unary_rpc_method_handler(
            servicer.ListPackages,
            request_deserializer=Message.FromString,
            response_serializer=Packages.SerializeToString,
        ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
        'com.appknox.bakerstreet.Moriarty', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


  class MycroftStub(object):
    # missing associated documentation comment in .proto file
    pass

    def __init__(self, channel):
      """Constructor.

      Args:
        channel: A grpc.Channel.
      """
      self.Poll = channel.unary_unary(
          '/com.appknox.bakerstreet.Mycroft/Poll',
          request_serializer=Device.SerializeToString,
          response_deserializer=Message.FromString,
          )
      self.PushFinding = channel.unary_unary(
          '/com.appknox.bakerstreet.Mycroft/PushFinding',
          request_serializer=Finding.SerializeToString,
          response_deserializer=Message.FromString,
          )


  class MycroftServicer(object):
    # missing associated documentation comment in .proto file
    pass

    def Poll(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def PushFinding(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')


  def add_MycroftServicer_to_server(servicer, server):
    rpc_method_handlers = {
        'Poll': grpc.unary_unary_rpc_method_handler(
            servicer.Poll,
            request_deserializer=Device.FromString,
            response_serializer=Message.SerializeToString,
        ),
        'PushFinding': grpc.unary_unary_rpc_method_handler(
            servicer.PushFinding,
            request_deserializer=Finding.FromString,
            response_serializer=Message.SerializeToString,
        ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
        'com.appknox.bakerstreet.Mycroft', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


  class BetaMoriartyServicer(object):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This class was generated
    only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0."""
    # missing associated documentation comment in .proto file
    pass
    def Info(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def Echo(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def LaunchApp(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def HealthCheck(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def RemovePackage(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def InstallPackage(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def ListPackages(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)


  class BetaMoriartyStub(object):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This class was generated
    only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0."""
    # missing associated documentation comment in .proto file
    pass
    def Info(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      # missing associated documentation comment in .proto file
      pass
      raise NotImplementedError()
    Info.future = None
    def Echo(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      # missing associated documentation comment in .proto file
      pass
      raise NotImplementedError()
    Echo.future = None
    def LaunchApp(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      # missing associated documentation comment in .proto file
      pass
      raise NotImplementedError()
    LaunchApp.future = None
    def HealthCheck(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      # missing associated documentation comment in .proto file
      pass
      raise NotImplementedError()
    HealthCheck.future = None
    def RemovePackage(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      # missing associated documentation comment in .proto file
      pass
      raise NotImplementedError()
    RemovePackage.future = None
    def InstallPackage(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      # missing associated documentation comment in .proto file
      pass
      raise NotImplementedError()
    InstallPackage.future = None
    def ListPackages(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      # missing associated documentation comment in .proto file
      pass
      raise NotImplementedError()
    ListPackages.future = None


  def beta_create_Moriarty_server(servicer, pool=None, pool_size=None, default_timeout=None, maximum_timeout=None):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This function was
    generated only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0"""
    request_deserializers = {
      ('com.appknox.bakerstreet.Moriarty', 'Echo'): Message.FromString,
      ('com.appknox.bakerstreet.Moriarty', 'HealthCheck'): Message.FromString,
      ('com.appknox.bakerstreet.Moriarty', 'Info'): Message.FromString,
      ('com.appknox.bakerstreet.Moriarty', 'InstallPackage'): Message.FromString,
      ('com.appknox.bakerstreet.Moriarty', 'LaunchApp'): Message.FromString,
      ('com.appknox.bakerstreet.Moriarty', 'ListPackages'): Message.FromString,
      ('com.appknox.bakerstreet.Moriarty', 'RemovePackage'): Message.FromString,
    }
    response_serializers = {
      ('com.appknox.bakerstreet.Moriarty', 'Echo'): Message.SerializeToString,
      ('com.appknox.bakerstreet.Moriarty', 'HealthCheck'): Message.SerializeToString,
      ('com.appknox.bakerstreet.Moriarty', 'Info'): Device.SerializeToString,
      ('com.appknox.bakerstreet.Moriarty', 'InstallPackage'): Message.SerializeToString,
      ('com.appknox.bakerstreet.Moriarty', 'LaunchApp'): Message.SerializeToString,
      ('com.appknox.bakerstreet.Moriarty', 'ListPackages'): Packages.SerializeToString,
      ('com.appknox.bakerstreet.Moriarty', 'RemovePackage'): Message.SerializeToString,
    }
    method_implementations = {
      ('com.appknox.bakerstreet.Moriarty', 'Echo'): face_utilities.unary_unary_inline(servicer.Echo),
      ('com.appknox.bakerstreet.Moriarty', 'HealthCheck'): face_utilities.unary_unary_inline(servicer.HealthCheck),
      ('com.appknox.bakerstreet.Moriarty', 'Info'): face_utilities.unary_unary_inline(servicer.Info),
      ('com.appknox.bakerstreet.Moriarty', 'InstallPackage'): face_utilities.unary_unary_inline(servicer.InstallPackage),
      ('com.appknox.bakerstreet.Moriarty', 'LaunchApp'): face_utilities.unary_unary_inline(servicer.LaunchApp),
      ('com.appknox.bakerstreet.Moriarty', 'ListPackages'): face_utilities.unary_unary_inline(servicer.ListPackages),
      ('com.appknox.bakerstreet.Moriarty', 'RemovePackage'): face_utilities.unary_unary_inline(servicer.RemovePackage),
    }
    server_options = beta_implementations.server_options(request_deserializers=request_deserializers, response_serializers=response_serializers, thread_pool=pool, thread_pool_size=pool_size, default_timeout=default_timeout, maximum_timeout=maximum_timeout)
    return beta_implementations.server(method_implementations, options=server_options)


  def beta_create_Moriarty_stub(channel, host=None, metadata_transformer=None, pool=None, pool_size=None):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This function was
    generated only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0"""
    request_serializers = {
      ('com.appknox.bakerstreet.Moriarty', 'Echo'): Message.SerializeToString,
      ('com.appknox.bakerstreet.Moriarty', 'HealthCheck'): Message.SerializeToString,
      ('com.appknox.bakerstreet.Moriarty', 'Info'): Message.SerializeToString,
      ('com.appknox.bakerstreet.Moriarty', 'InstallPackage'): Message.SerializeToString,
      ('com.appknox.bakerstreet.Moriarty', 'LaunchApp'): Message.SerializeToString,
      ('com.appknox.bakerstreet.Moriarty', 'ListPackages'): Message.SerializeToString,
      ('com.appknox.bakerstreet.Moriarty', 'RemovePackage'): Message.SerializeToString,
    }
    response_deserializers = {
      ('com.appknox.bakerstreet.Moriarty', 'Echo'): Message.FromString,
      ('com.appknox.bakerstreet.Moriarty', 'HealthCheck'): Message.FromString,
      ('com.appknox.bakerstreet.Moriarty', 'Info'): Device.FromString,
      ('com.appknox.bakerstreet.Moriarty', 'InstallPackage'): Message.FromString,
      ('com.appknox.bakerstreet.Moriarty', 'LaunchApp'): Message.FromString,
      ('com.appknox.bakerstreet.Moriarty', 'ListPackages'): Packages.FromString,
      ('com.appknox.bakerstreet.Moriarty', 'RemovePackage'): Message.FromString,
    }
    cardinalities = {
      'Echo': cardinality.Cardinality.UNARY_UNARY,
      'HealthCheck': cardinality.Cardinality.UNARY_UNARY,
      'Info': cardinality.Cardinality.UNARY_UNARY,
      'InstallPackage': cardinality.Cardinality.UNARY_UNARY,
      'LaunchApp': cardinality.Cardinality.UNARY_UNARY,
      'ListPackages': cardinality.Cardinality.UNARY_UNARY,
      'RemovePackage': cardinality.Cardinality.UNARY_UNARY,
    }
    stub_options = beta_implementations.stub_options(host=host, metadata_transformer=metadata_transformer, request_serializers=request_serializers, response_deserializers=response_deserializers, thread_pool=pool, thread_pool_size=pool_size)
    return beta_implementations.dynamic_stub(channel, 'com.appknox.bakerstreet.Moriarty', cardinalities, options=stub_options)


  class BetaMycroftServicer(object):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This class was generated
    only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0."""
    # missing associated documentation comment in .proto file
    pass
    def Poll(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def PushFinding(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)


  class BetaMycroftStub(object):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This class was generated
    only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0."""
    # missing associated documentation comment in .proto file
    pass
    def Poll(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      # missing associated documentation comment in .proto file
      pass
      raise NotImplementedError()
    Poll.future = None
    def PushFinding(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      # missing associated documentation comment in .proto file
      pass
      raise NotImplementedError()
    PushFinding.future = None


  def beta_create_Mycroft_server(servicer, pool=None, pool_size=None, default_timeout=None, maximum_timeout=None):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This function was
    generated only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0"""
    request_deserializers = {
      ('com.appknox.bakerstreet.Mycroft', 'Poll'): Device.FromString,
      ('com.appknox.bakerstreet.Mycroft', 'PushFinding'): Finding.FromString,
    }
    response_serializers = {
      ('com.appknox.bakerstreet.Mycroft', 'Poll'): Message.SerializeToString,
      ('com.appknox.bakerstreet.Mycroft', 'PushFinding'): Message.SerializeToString,
    }
    method_implementations = {
      ('com.appknox.bakerstreet.Mycroft', 'Poll'): face_utilities.unary_unary_inline(servicer.Poll),
      ('com.appknox.bakerstreet.Mycroft', 'PushFinding'): face_utilities.unary_unary_inline(servicer.PushFinding),
    }
    server_options = beta_implementations.server_options(request_deserializers=request_deserializers, response_serializers=response_serializers, thread_pool=pool, thread_pool_size=pool_size, default_timeout=default_timeout, maximum_timeout=maximum_timeout)
    return beta_implementations.server(method_implementations, options=server_options)


  def beta_create_Mycroft_stub(channel, host=None, metadata_transformer=None, pool=None, pool_size=None):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This function was
    generated only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0"""
    request_serializers = {
      ('com.appknox.bakerstreet.Mycroft', 'Poll'): Device.SerializeToString,
      ('com.appknox.bakerstreet.Mycroft', 'PushFinding'): Finding.SerializeToString,
    }
    response_deserializers = {
      ('com.appknox.bakerstreet.Mycroft', 'Poll'): Message.FromString,
      ('com.appknox.bakerstreet.Mycroft', 'PushFinding'): Message.FromString,
    }
    cardinalities = {
      'Poll': cardinality.Cardinality.UNARY_UNARY,
      'PushFinding': cardinality.Cardinality.UNARY_UNARY,
    }
    stub_options = beta_implementations.stub_options(host=host, metadata_transformer=metadata_transformer, request_serializers=request_serializers, response_deserializers=response_deserializers, thread_pool=pool, thread_pool_size=pool_size)
    return beta_implementations.dynamic_stub(channel, 'com.appknox.bakerstreet.Mycroft', cardinalities, options=stub_options)
except ImportError:
  pass
# @@protoc_insertion_point(module_scope)
