module Errors.Types exposing (ClientLogin(..))


type DataError
    = DataError String


type IdpError
    = IdpError String


type InvalidIdType
    = InvalidIdType String


type LoginError
    = LoginError String


type NetworkTypeError
    = NetworkTypeError String


type ResponseError
    = ResponseError String


type Status4xxError
    = Status4xxError String


type ClientLogin
    = GroupNotFoundError Status4xxError
    | HttpResponseError ResponseError
    | InvalidIdError InvalidIdType
    | JsonDataError DataError
    | LoginResponseFailed LoginError
    | LoginRequestFailed LoginError
    | LoginRequired LoginError
    | NetworkError NetworkTypeError
    | OauthException IdpError
    | OrganizationNotFoundError Status4xxError
    | UserNotFoundError Status4xxError
