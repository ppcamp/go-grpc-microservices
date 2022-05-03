<script lang="ts">
  export let onClick: Function = null;
  export let name: string;
  export let success: boolean = false;

  const click = () => onClick();
</script>

<button on:click={click} class:success>{name}</button>

<style>
  button {
    border-radius: 0.5em;
    padding: 7px 20px;
    border: none;
    background-color: rgb(28, 83, 202);
    color: white;
  }
  button:active {
    background-color: rgb(52, 83, 151);
  }
  .success {
    background-color: rgb(10, 92, 10);
  }
  .success:active {
    background-color: rgb(56, 117, 56);
  }
</style>
