module Main exposing (main)

import Html exposing (div, h1, text)
import Html.Attributes exposing (..)

view : a -> Html.Html msg
view model =
    div [ class "content" ]
        [ h1 [] [ text "AccountEd"]]

main : Html.Html msg
main =
    view "Not Implemented"