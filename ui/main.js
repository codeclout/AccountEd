import { Elm } from "./src/Main.elm";

const root = document.querySelector("#app");
const app = Elm.Main.init({ node: root });
