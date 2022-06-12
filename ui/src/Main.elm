port module Main exposing (main)

import Browser exposing (Document, UrlRequest)
import Browser.Navigation as Navigation
import Dict exposing (update)
import Errors.Types exposing (ClientLogin(..))
import Html exposing (Html, div, h1, text)
import Html.Attributes exposing (..)
import Json.Decode exposing (bool, decodeString, int, maybe, string)
import Routes
import Url exposing (Url)


type Page
    = Admin
    | Auth
    | Login
    | Register
    | NotFound



--    | ServerError


type Token
    = Token String


type alias Flags =
    { storedToken : Maybe String }


type alias Model =
    { page : Page
    , navKey : Navigation.Key
    , token : Maybe Token
    }


type Msg
    = NewRoute (Maybe Routes.Route)
    | Req UrlRequest


initialModel : Maybe Token -> Navigation.Key -> Model
initialModel t navigationKey =
    { page = NotFound
    , navKey = navigationKey
    , token = t
    }


init : Flags -> Url -> Navigation.Key -> ( Model, Cmd Msg )
init flags url navigationKey =
    let
        t =
            Maybe.map Token flags.storedToken
    in
    setNewPage (Routes.match url) (initialModel t navigationKey)


main =
    Browser.application
        { init = init
        , view = view
        , update = update
        , subscriptions = subscriptions
        , onUrlRequest = Req
        , onUrlChange = Routes.match >> NewRoute
        }


subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.none


setNewPage : Maybe Routes.Route -> Model -> ( Model, Cmd Msg )
setNewPage maybeRoute model =
    case maybeRoute of
        Just Routes.Admin ->
            ( { model | page = Admin }, Cmd.none )

        Just Routes.Auth ->
            ( { model | page = Auth }, Cmd.none )

        Just (Routes.Login authToken) ->
            ( { model | page = Login, token = Just (Token authToken) }, storeToken authToken )

        Just Routes.Register ->
            ( { model | page = Register }, Cmd.none )

        Nothing ->
            ( { model | page = NotFound }, Cmd.none )


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        NewRoute maybeRoute ->
            setNewPage maybeRoute model

        _ ->
            ( model, Cmd.none )


viewContent : Page -> ( String, Html Msg )
viewContent page =
    case page of
        Admin ->
            ( "Administrator"
            , h1 [] [ text "Admin" ]
            )

        Auth ->
            ( "Auth"
            , h1 [] [ text "Auth" ]
            )

        Login ->
            ( "Login Form"
            , h1 [] [ text "Login" ]
            )

        Register ->
            ( "Onboarding"
            , h1 [] [ text "Register New Account" ]
            )

        NotFound ->
            ( "Not Found"
            , div [ class "not-found" ]
                [ h1 [] [ text "Page Not Found" ] ]
            )


view : Model -> Document Msg
view model =
    let
        ( title, content ) =
            viewContent model.page
    in
    { title = title
    , body = [ content ]
    }


port storeToken : String -> Cmd msg
