<script lang="ts">
  import { Router, Route } from "svelte-navigator";

  import PrivateRoute from "./views/PrivateRoute.svelte";
  import CreateUser from "./views/CreateUser.svelte";
  import Dashboard from "./views/Dashboard.svelte";
  import Login from "./views/Login.svelte";
</script>

<Router>
  <!-- <Route path="login"><Login /></Route> -->
  <Route path="/"><CreateUser /></Route>
  <!-- <PrivateRoute path="/"><Dashboard /></PrivateRoute> -->
</Router>
