<script lang="ts">
  import { Route, navigate } from "svelte-navigator";
  import auth from "../stores/auth";

  export let path: string;

  //   $: isAuthenticated = $auth.Token;
  //   $: if (!$isAuthenticated) {
  //     navigate("/login", { state: { from: $location.pathname }, replace: true });
  //   }
  let isAuthenticated: boolean;
  auth.Token.subscribe((token) => {
    isAuthenticated = token.length > 0;
  });

  $: if (!isAuthenticated) {
    navigate("/login", { state: { from: location.pathname }, replace: true });
  }
</script>

{#if isAuthenticated}
  <Route {path} let:params let:location let:navigate>
    <slot {params} {location} {navigate} />
  </Route>
{/if}
