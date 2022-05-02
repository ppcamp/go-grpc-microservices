<script lang="ts">
</script>

<div class="container">
  <header>Hader</header>
  <main>Main</main>
  <aside>Aside</aside>
  <footer>Footer</footer>
</div>

<style>
  :global(body) {
    margin: 0;
    padding: 0;
    box-sizing: border-box;
    background: #fafafa;
  }

  .container {
    display: grid;
    grid-template-columns: 3fr 1fr 1fr;
    grid-template-rows: 20vh 70vh 10vh;
    grid-template-areas:
      "header header header"
      "main main aside"
      "footer footer footer";
    gap: 1em 1em;
  }

  header {
    background-color: rgb(9, 255, 0);
    grid-area: header;
    /* grid-column: 1/4; */
  }

  main {
    background-color: aqua;
    grid-area: main;
    /* grid-column: 1/2; */
  }

  footer {
    background-color: red;
    grid-area: footer;
    /* grid-column: 1/4; */
  }

  aside {
    background-color: yellow;
    grid-area: aside;
    /* grid-column: 2/4; */
  }
</style>
