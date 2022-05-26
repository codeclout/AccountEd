import { Elm } from "./src/Main.elm";

const root = document.querySelector("#app");
const app = Elm.Main.init({
  flags: { storedToken: localStorage.getItem("__AccountEd__") },
  node: root,
});

app.ports.storeToken.subscribe((t) => {
  localStorage.setItem("__AccountEd__", t);
});
