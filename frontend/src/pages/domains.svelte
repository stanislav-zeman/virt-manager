<script lang="ts">
    import { Domains } from "../../wailsjs/go/app/App";
    import { Button } from "flowbite-svelte";
    import Domain from "../components/domain.svelte";
</script>

<div class="w-4/5">
    <div>Domains!</div>
    <!-- <div> -->
    <!--     {#await Domains() then dss} -->
    <!--         {#each dss as d} -->
    <!--             <Domain name={d.name} status={d.status} /> -->
    <!--         {/each} -->
    <!--     {/await} -->
    <!-- </div> -->
    <!-- <div class="justify-start"> -->
    <!--     <Button on:click{domains()}>Refresh</Button> -->
    <!-- </div> -->
</div>
