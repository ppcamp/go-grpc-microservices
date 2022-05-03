<script lang="ts">
  import Button from "../components/Button.svelte";

  import Card from "../components/Card.svelte";
  import TextField from "../components/TextField.svelte";

  const devname = "Pedro Augusto C Santos";
</script>

<div class="container">
  <main>
    <div />
    <div class="form">
      <Card>
        <h3>Login</h3>
        <form action="#" method="post">
          <TextField type="email" label="Email" name="email" />
          <TextField type="password" label="Password" name="password" />
          <div class="flexend grid">
            <Button name="Submit" success />
          </div>
        </form>
      </Card>
    </div>
  </main>
  <footer>
    <div>
      Developed by. {devname}
    </div>
  </footer>
</div>

<style>
  :global(body) {
    margin: 0;
    padding: 0;
    background: #fff;
  }
  :root {
    --main-color: rgb(71, 52, 177);
  }
  :global(*) {
    box-sizing: border-box;
  }

  main {
    grid-area: main;
    display: flex;
    align-items: center;
    justify-content: center;
  }
  footer {
    grid-area: footer;
    display: flex;
    justify-content: center;
    align-items: center;
  }
  .container {
    display: grid;
    grid-template-columns: 1fr;
    grid-template-rows: 10vh 60vh 10vh;
    grid-template-areas: "." "main" "footer";
    gap: 10vh 0em;
  }

  .form {
    width: max(500px, 10vh);
    height: 1fr;
  }

  form {
    width: 100%;
    padding: 1em;
  }

  .grid {
    display: flex;
    height: max-content;
    min-height: 50px;
  }
  .flexend {
    justify-content: flex-end;
    align-items: flex-end;
  }
</style>
