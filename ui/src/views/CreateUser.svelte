<script lang="ts">
  import { create_user } from "../api/user";

  import TextField from "../components/TextField.svelte";
  import { useFocus } from "svelte-navigator";

  const onClick = async () => {
    await create_user({
      password: "some",
      birthdate: new Date(),
      email: "p.aihistpca",
      nickname: "asda",
      first_name: "text",
      last_name: "afas",
      middle_name: "asca",
    });
  };

  const registerFocus = useFocus();
</script>

<main>
  <div class="grid">
    <form action="#">
      <TextField name="user" type="text" label="UserId" />
      <TextField name="password" type="password" label="Password" />
      <button on:click={onClick} use:registerFocus type="button">Create</button>
    </form>
  </div>
</main>
