module Routes exposing (Route(..), match)

import Html exposing (a)
import Url exposing (Url)
import Url.Parser as Parser exposing (Parser)


type Route
    = Admin
    | Login
    | Register



-- create routes as type Route


routes : Parser (Route -> r) r
routes =
    Parser.oneOf
        [ Parser.map Admin Parser.top
        , Parser.map Login (Parser.s "login")
        , Parser.map Register (Parser.s "register")
        ]


match : Url -> Maybe Route
match url =
    Parser.parse routes url
