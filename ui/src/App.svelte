<script lang="ts">
  import { Router, Route } from "svelte-navigator";
  import Dashboard from "./views/Dashboard.svelte";
  import Login from "./views/Login.svelte";
</script>

<Router>
  <Route path="login"><Login /></Route>
  <Route path="dashboard"><Dashboard /></Route>
</Router>
