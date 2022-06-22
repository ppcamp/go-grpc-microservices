<script lang="ts">
  export let transition: boolean = false;
</script>

<div class="container" class:transition>
  <slot />
</div>

<style>
  .container {
    padding: 1em;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    display: flex;
    background-color: rgb(255, 255, 255);
    border: none;
    width: 100%;
    height: 100%;
    border-radius: 0.5em;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
  }
  .transition {
    transition: 0.3s;
  }
  .transition:hover {
    box-shadow: 0 12px 24px 0 rgba(0, 0, 0, 0.2);
  }
</style>
