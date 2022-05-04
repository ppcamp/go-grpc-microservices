<script lang="ts">
  import { Router, Route } from "svelte-navigator";

  import PrivateRoute from "./utils/PrivateRoute.svelte";
  import Dashboard from "./views/Dashboard.svelte";
  import Login from "./views/Login.svelte";
</script>

<Router>
  <Route path="login"><Login /></Route>
  <PrivateRoute path="/"><Dashboard /></PrivateRoute>
  <PrivateRoute path="*" />
</Router>
