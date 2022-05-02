<script lang="ts">
  export let type: string = "text";
  export let label: string = "Label";
  export let name: string = "";
  export let required: boolean = false;
</script>

<div class="item">
  <label for={name} class="description">{label}</label>
  <input {type} {name} id={name} {required} class="text" />
</div>

<style>
  .item {
    margin: 1.5em 0em;
  }
  .text {
    border: none;
    width: 100%;
    border-bottom: 0.1em solid var(--main-color);
  }
  .text:focus {
    outline: none;
  }
  .description {
    padding-bottom: 0.5em;
  }
  /* Change the white to any color */
  input:-webkit-autofill,
  input:-webkit-autofill:hover,
  input:-webkit-autofill:focus,
  input:-webkit-autofill:active {
    -webkit-box-shadow: 0 0 0 30px white inset !important;
  }
</style>
