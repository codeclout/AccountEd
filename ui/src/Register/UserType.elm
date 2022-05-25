module Register.UserType exposing (Model, Msg, init, update, view)

import Html exposing (Html, a, div, fieldset, h1, input, label, span, text)
import Html.Attributes exposing (..)
import Html.Events exposing (onClick)



-- import Debug exposing (log)
-- MODEL


type alias Model =
    { orgType : Organization
    , homeSchooler : String
    , group : String
    , school : String
    }


initialModel : Model
initialModel =
    { group = "Small Group"
    , homeSchooler = "Person"
    , orgType = HomeSchooler
    , school = "School"
    }


init : Model
init =
    initialModel


type Organization
    = HomeSchooler
    | SmallGroup
    | School



-- UPDATE


type Msg
    = SwitchTo Organization


update : Msg -> Model -> Model
update msg model =
    case msg of
        SwitchTo newOrgType ->
            { model | orgType = newOrgType }



-- VIEW


view : Model -> Html Msg
view model =
    div []
        [ div []
            [ h1 [] [ text "Tell us what's your situation" ]
            , fieldset []
                [ div []
                    [ radio (SwitchTo HomeSchooler) model.homeSchooler
                    , radio (SwitchTo SmallGroup) model.group
                    , radio (SwitchTo School) model.school
                    ]
                ]
            , div []
                [ a [ href "onboarding-step-2" ]
                    [ text "Next" ]
                ]
            ]
        ]


radio : msg -> String -> Html msg
radio msg tname =
    label []
        [ input [ type_ "radio", name "organization-onboarding-radio", onClick msg ] []
        , div []
            [ span []
                [ text tname ]
            ]
        ]
