<script lang="ts">
    import Window from "../layout/Window.svelte";
    import { Domains, Connected } from "../../wailsjs/go/app/App";
    import { Button, P, Spinner } from "flowbite-svelte";
    import Domain from "../components/Domain.svelte";
    import { getNotificationsContext } from "svelte-notifications";

    const { addNotification } = getNotificationsContext();
</script>

<Window heading="Domains">
    {#await Connected()}
        <Spinner class="m-6" />
    {:then state}
        {#if state}
            <div>
                {#await Domains()}
                    <Spinner class="m-6" />
                {:then dss}
                    {#each dss as d}
                        <Domain name={d.name} status={d.status} />
                    {/each}
                {:catch error}
                    An error occured: {error}
                {/await}
            </div>
            <div class="justify-start">
                <Button on:click{domains()}>Refresh</Button>
            </div>
        {:else}
            <P class="text-center text-white" size="xl"
                >You are not connected!</P
            >
        {/if}
    {/await}
</Window>
