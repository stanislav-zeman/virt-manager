<script lang="ts">
    import Window from "../layout/Window.svelte";
    import { Domains } from "../../wailsjs/go/app/App";
    import { Button } from "flowbite-svelte";
    import Domain from "../components/Domain.svelte";
</script>

<Window heading="Domains">
    <div>
        {#await Domains()}
            Loading domains...
        {:then dss}
            {#each dss as d}
                <Domain name={d.name} status={d.status} />
            {/each}
        {:catch error}
            An error occured: {error.message}.
        {/await}
    </div>
    <div class="justify-start">
        <Button on:click{domains()}>Refresh</Button>
    </div>
</Window>
