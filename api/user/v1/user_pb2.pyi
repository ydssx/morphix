"""
@generated by mypy-protobuf.  Do not edit manually!
isort:skip_file
"""
import builtins
import collections.abc
import google.protobuf.descriptor
import google.protobuf.internal.containers
import google.protobuf.internal.enum_type_wrapper
import google.protobuf.message
import sys
import typing

if sys.version_info >= (3, 10):
    import typing as typing_extensions
else:
    import typing_extensions

DESCRIPTOR: google.protobuf.descriptor.FileDescriptor

@typing_extensions.final
class User(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    ID_FIELD_NUMBER: builtins.int
    USERNAME_FIELD_NUMBER: builtins.int
    PASSWORD_FIELD_NUMBER: builtins.int
    EMAIL_FIELD_NUMBER: builtins.int
    PHONE_FIELD_NUMBER: builtins.int
    AVATAR_FIELD_NUMBER: builtins.int
    NICKNAME_FIELD_NUMBER: builtins.int
    id: builtins.int
    username: builtins.str
    password: builtins.str
    email: builtins.str
    phone: builtins.str
    avatar: builtins.str
    nickname: builtins.str
    def __init__(
        self,
        *,
        id: builtins.int = ...,
        username: builtins.str = ...,
        password: builtins.str = ...,
        email: builtins.str = ...,
        phone: builtins.str = ...,
        avatar: builtins.str = ...,
        nickname: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["avatar", b"avatar", "email", b"email", "id", b"id", "nickname", b"nickname", "password", b"password", "phone", b"phone", "username", b"username"]) -> None: ...

global___User = User

@typing_extensions.final
class RegistrationRequest(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    class _RegisterType:
        ValueType = typing.NewType("ValueType", builtins.int)
        V: typing_extensions.TypeAlias = ValueType

    class _RegisterTypeEnumTypeWrapper(google.protobuf.internal.enum_type_wrapper._EnumTypeWrapper[RegistrationRequest._RegisterType.ValueType], builtins.type):
        DESCRIPTOR: google.protobuf.descriptor.EnumDescriptor
        SMS: RegistrationRequest._RegisterType.ValueType  # 0
        """通过短信验证码注册"""
        PASSWORD: RegistrationRequest._RegisterType.ValueType  # 1
        """通过用户名密码注册"""

    class RegisterType(_RegisterType, metaclass=_RegisterTypeEnumTypeWrapper):
        """注册类型"""

    SMS: RegistrationRequest.RegisterType.ValueType  # 0
    """通过短信验证码注册"""
    PASSWORD: RegistrationRequest.RegisterType.ValueType  # 1
    """通过用户名密码注册"""

    USERNAME_FIELD_NUMBER: builtins.int
    PASSWORD_FIELD_NUMBER: builtins.int
    EMAIL_FIELD_NUMBER: builtins.int
    PHONE_FIELD_NUMBER: builtins.int
    SMS_CODE_FIELD_NUMBER: builtins.int
    REGISTER_TYPE_FIELD_NUMBER: builtins.int
    username: builtins.str
    """用户名"""
    password: builtins.str
    """密码"""
    email: builtins.str
    """邮箱"""
    phone: builtins.str
    """手机号"""
    sms_code: builtins.str
    """短信验证码"""
    register_type: global___RegistrationRequest.RegisterType.ValueType
    def __init__(
        self,
        *,
        username: builtins.str = ...,
        password: builtins.str = ...,
        email: builtins.str = ...,
        phone: builtins.str = ...,
        sms_code: builtins.str = ...,
        register_type: global___RegistrationRequest.RegisterType.ValueType = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["email", b"email", "password", b"password", "phone", b"phone", "register_type", b"register_type", "sms_code", b"sms_code", "username", b"username"]) -> None: ...

global___RegistrationRequest = RegistrationRequest

@typing_extensions.final
class LoginRequest(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    USERNAME_FIELD_NUMBER: builtins.int
    PASSWORD_FIELD_NUMBER: builtins.int
    PHONE_NUMBER_FIELD_NUMBER: builtins.int
    username: builtins.str
    """用户名"""
    password: builtins.str
    """密码"""
    phone_number: builtins.str
    """手机号"""
    def __init__(
        self,
        *,
        username: builtins.str = ...,
        password: builtins.str = ...,
        phone_number: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["password", b"password", "phone_number", b"phone_number", "username", b"username"]) -> None: ...

global___LoginRequest = LoginRequest

@typing_extensions.final
class LogoutRequest(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    USER_ID_FIELD_NUMBER: builtins.int
    user_id: builtins.int
    def __init__(
        self,
        *,
        user_id: builtins.int = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["user_id", b"user_id"]) -> None: ...

global___LogoutRequest = LogoutRequest

@typing_extensions.final
class UpdateProfileRequest(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    EMAIL_FIELD_NUMBER: builtins.int
    PHONE_FIELD_NUMBER: builtins.int
    USERNAME_FIELD_NUMBER: builtins.int
    email: builtins.str
    phone: builtins.str
    username: builtins.str
    def __init__(
        self,
        *,
        email: builtins.str = ...,
        phone: builtins.str = ...,
        username: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["email", b"email", "phone", b"phone", "username", b"username"]) -> None: ...

global___UpdateProfileRequest = UpdateProfileRequest

@typing_extensions.final
class ResetPasswordRequest(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    USERNAME_FIELD_NUMBER: builtins.int
    VERIFICATION_CODE_FIELD_NUMBER: builtins.int
    NEW_PASSWORD_FIELD_NUMBER: builtins.int
    username: builtins.str
    """用户名"""
    verification_code: builtins.str
    """验证码"""
    new_password: builtins.str
    """新密码"""
    def __init__(
        self,
        *,
        username: builtins.str = ...,
        verification_code: builtins.str = ...,
        new_password: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["new_password", b"new_password", "username", b"username", "verification_code", b"verification_code"]) -> None: ...

global___ResetPasswordRequest = ResetPasswordRequest

@typing_extensions.final
class AuthenticationResponse(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    USER_ID_FIELD_NUMBER: builtins.int
    TOKEN_FIELD_NUMBER: builtins.int
    user_id: builtins.str
    token: builtins.str
    """认证令牌"""
    def __init__(
        self,
        *,
        user_id: builtins.str = ...,
        token: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["token", b"token", "user_id", b"user_id"]) -> None: ...

global___AuthenticationResponse = AuthenticationResponse

@typing_extensions.final
class AuthorizationRequest(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    USER_ID_FIELD_NUMBER: builtins.int
    RESOURCE_FIELD_NUMBER: builtins.int
    ACTIONS_FIELD_NUMBER: builtins.int
    user_id: builtins.str
    resource: builtins.str
    @property
    def actions(self) -> google.protobuf.internal.containers.RepeatedScalarFieldContainer[builtins.str]: ...
    def __init__(
        self,
        *,
        user_id: builtins.str = ...,
        resource: builtins.str = ...,
        actions: collections.abc.Iterable[builtins.str] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["actions", b"actions", "resource", b"resource", "user_id", b"user_id"]) -> None: ...

global___AuthorizationRequest = AuthorizationRequest

@typing_extensions.final
class UserListResponse(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    USERS_FIELD_NUMBER: builtins.int
    @property
    def users(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___User]: ...
    def __init__(
        self,
        *,
        users: collections.abc.Iterable[global___User] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["users", b"users"]) -> None: ...

global___UserListResponse = UserListResponse

@typing_extensions.final
class UserListRequest(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    PAGE_FIELD_NUMBER: builtins.int
    LIMIT_FIELD_NUMBER: builtins.int
    page: builtins.int
    limit: builtins.int
    def __init__(
        self,
        *,
        page: builtins.int = ...,
        limit: builtins.int = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["limit", b"limit", "page", b"page"]) -> None: ...

global___UserListRequest = UserListRequest

@typing_extensions.final
class ManageUserPermissionRequest(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    class _Mode:
        ValueType = typing.NewType("ValueType", builtins.int)
        V: typing_extensions.TypeAlias = ValueType

    class _ModeEnumTypeWrapper(google.protobuf.internal.enum_type_wrapper._EnumTypeWrapper[ManageUserPermissionRequest._Mode.ValueType], builtins.type):
        DESCRIPTOR: google.protobuf.descriptor.EnumDescriptor
        USER_PERMISSION_ADD: ManageUserPermissionRequest._Mode.ValueType  # 0
        """增加用户权限"""
        USER_PERMISSION_DELETE: ManageUserPermissionRequest._Mode.ValueType  # 1
        """删除用户权限"""
        USER_PERMISSION_UPDATE: ManageUserPermissionRequest._Mode.ValueType  # 2
        """更新用户权限"""
        USER_PERMISSION_REPLACE: ManageUserPermissionRequest._Mode.ValueType  # 3
        """替换用户权限"""
        ROLE_PERMISSION_ADD: ManageUserPermissionRequest._Mode.ValueType  # 4
        """增加角色权限"""
        ROLE_PERMISSION_DELETE: ManageUserPermissionRequest._Mode.ValueType  # 5
        """删除角色权限"""
        ROLE_PERMISSION_UPDATE: ManageUserPermissionRequest._Mode.ValueType  # 6
        """更新角色权限"""
        ROLE_PERMISSION_REPLACE: ManageUserPermissionRequest._Mode.ValueType  # 7
        """替换角色权限"""

    class Mode(_Mode, metaclass=_ModeEnumTypeWrapper): ...
    USER_PERMISSION_ADD: ManageUserPermissionRequest.Mode.ValueType  # 0
    """增加用户权限"""
    USER_PERMISSION_DELETE: ManageUserPermissionRequest.Mode.ValueType  # 1
    """删除用户权限"""
    USER_PERMISSION_UPDATE: ManageUserPermissionRequest.Mode.ValueType  # 2
    """更新用户权限"""
    USER_PERMISSION_REPLACE: ManageUserPermissionRequest.Mode.ValueType  # 3
    """替换用户权限"""
    ROLE_PERMISSION_ADD: ManageUserPermissionRequest.Mode.ValueType  # 4
    """增加角色权限"""
    ROLE_PERMISSION_DELETE: ManageUserPermissionRequest.Mode.ValueType  # 5
    """删除角色权限"""
    ROLE_PERMISSION_UPDATE: ManageUserPermissionRequest.Mode.ValueType  # 6
    """更新角色权限"""
    ROLE_PERMISSION_REPLACE: ManageUserPermissionRequest.Mode.ValueType  # 7
    """替换角色权限"""

    USER_ID_FIELD_NUMBER: builtins.int
    ROLE_IDS_FIELD_NUMBER: builtins.int
    PERMISSION_IDS_FIELD_NUMBER: builtins.int
    MODE_FIELD_NUMBER: builtins.int
    user_id: builtins.int
    @property
    def role_ids(self) -> google.protobuf.internal.containers.RepeatedScalarFieldContainer[builtins.int]: ...
    @property
    def permission_ids(self) -> google.protobuf.internal.containers.RepeatedScalarFieldContainer[builtins.int]: ...
    mode: global___ManageUserPermissionRequest.Mode.ValueType
    """操作模式"""
    def __init__(
        self,
        *,
        user_id: builtins.int = ...,
        role_ids: collections.abc.Iterable[builtins.int] | None = ...,
        permission_ids: collections.abc.Iterable[builtins.int] | None = ...,
        mode: global___ManageUserPermissionRequest.Mode.ValueType = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["mode", b"mode", "permission_ids", b"permission_ids", "role_ids", b"role_ids", "user_id", b"user_id"]) -> None: ...

global___ManageUserPermissionRequest = ManageUserPermissionRequest

@typing_extensions.final
class LogEntry(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    USER_ID_FIELD_NUMBER: builtins.int
    ACTION_FIELD_NUMBER: builtins.int
    TIMESTAMP_FIELD_NUMBER: builtins.int
    RESOURCE_FIELD_NUMBER: builtins.int
    MESSAGE_FIELD_NUMBER: builtins.int
    user_id: builtins.int
    action: builtins.str
    """操作 例如: login, logout"""
    timestamp: builtins.str
    """时间戳"""
    resource: builtins.str
    """资源 例如: /api/users"""
    message: builtins.str
    """消息 例如: 登录成功"""
    def __init__(
        self,
        *,
        user_id: builtins.int = ...,
        action: builtins.str = ...,
        timestamp: builtins.str = ...,
        resource: builtins.str = ...,
        message: builtins.str = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["action", b"action", "message", b"message", "resource", b"resource", "timestamp", b"timestamp", "user_id", b"user_id"]) -> None: ...

global___LogEntry = LogEntry

@typing_extensions.final
class GetUserRequest(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    USER_ID_FIELD_NUMBER: builtins.int
    user_id: builtins.int
    def __init__(
        self,
        *,
        user_id: builtins.int = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["user_id", b"user_id"]) -> None: ...

global___GetUserRequest = GetUserRequest

@typing_extensions.final
class GetUserPermissionRequest(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    USER_ID_FIELD_NUMBER: builtins.int
    user_id: builtins.int
    def __init__(
        self,
        *,
        user_id: builtins.int = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["user_id", b"user_id"]) -> None: ...

global___GetUserPermissionRequest = GetUserPermissionRequest

@typing_extensions.final
class UserPermissionListResponse(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    PERMISSION_FIELD_NUMBER: builtins.int
    @property
    def permission(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___UserPermission]: ...
    def __init__(
        self,
        *,
        permission: collections.abc.Iterable[global___UserPermission] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["permission", b"permission"]) -> None: ...

global___UserPermissionListResponse = UserPermissionListResponse

@typing_extensions.final
class UserPermission(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    RESOURCE_FIELD_NUMBER: builtins.int
    ACTIONS_FIELD_NUMBER: builtins.int
    ROLES_FIELD_NUMBER: builtins.int
    resource: builtins.str
    """资源 例如: /api/users"""
    @property
    def actions(self) -> google.protobuf.internal.containers.RepeatedScalarFieldContainer[builtins.str]:
        """权限 例如: GET, POST, PUT, DELETE"""
    @property
    def roles(self) -> google.protobuf.internal.containers.RepeatedScalarFieldContainer[builtins.str]:
        """角色 例如: admin, user"""
    def __init__(
        self,
        *,
        resource: builtins.str = ...,
        actions: collections.abc.Iterable[builtins.str] | None = ...,
        roles: collections.abc.Iterable[builtins.str] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["actions", b"actions", "resource", b"resource", "roles", b"roles"]) -> None: ...

global___UserPermission = UserPermission

@typing_extensions.final
class UserActivityListResponse(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    ACTIVITY_FIELD_NUMBER: builtins.int
    @property
    def activity(self) -> google.protobuf.internal.containers.RepeatedCompositeFieldContainer[global___UserActivity]: ...
    def __init__(
        self,
        *,
        activity: collections.abc.Iterable[global___UserActivity] | None = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["activity", b"activity"]) -> None: ...

global___UserActivityListResponse = UserActivityListResponse

@typing_extensions.final
class GetUserActivityRequest(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    USER_ID_FIELD_NUMBER: builtins.int
    PAGE_FIELD_NUMBER: builtins.int
    LIMIT_FIELD_NUMBER: builtins.int
    user_id: builtins.int
    page: builtins.int
    limit: builtins.int
    def __init__(
        self,
        *,
        user_id: builtins.int = ...,
        page: builtins.int = ...,
        limit: builtins.int = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["limit", b"limit", "page", b"page", "user_id", b"user_id"]) -> None: ...

global___GetUserActivityRequest = GetUserActivityRequest

@typing_extensions.final
class UserActivity(google.protobuf.message.Message):
    DESCRIPTOR: google.protobuf.descriptor.Descriptor

    TIMESTAMP_FIELD_NUMBER: builtins.int
    ACTION_FIELD_NUMBER: builtins.int
    RESOURCE_FIELD_NUMBER: builtins.int
    MESSAGE_FIELD_NUMBER: builtins.int
    ID_FIELD_NUMBER: builtins.int
    timestamp: builtins.str
    action: builtins.str
    """操作 例如: login, logout"""
    resource: builtins.str
    """资源 例如: /api/users"""
    message: builtins.str
    id: builtins.int
    def __init__(
        self,
        *,
        timestamp: builtins.str = ...,
        action: builtins.str = ...,
        resource: builtins.str = ...,
        message: builtins.str = ...,
        id: builtins.int = ...,
    ) -> None: ...
    def ClearField(self, field_name: typing_extensions.Literal["action", b"action", "id", b"id", "message", b"message", "resource", b"resource", "timestamp", b"timestamp"]) -> None: ...

global___UserActivity = UserActivity
