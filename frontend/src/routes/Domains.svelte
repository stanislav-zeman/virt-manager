<script lang="ts">
    import { onMount } from "svelte";
    import Window from "../layout/Window.svelte";
    import { Domains, Connected } from "../../wailsjs/go/app/App";
    import { domain } from "../../wailsjs/go/models";
    import { Button, P } from "flowbite-svelte";
    import Domain from "../components/Domain.svelte";
    import { getNotificationsContext } from "svelte-notifications";

    const { addNotification } = getNotificationsContext();

    let connected: boolean;
    let domains: domain.Domain[] = [];
    async function reloadDomains() {
        try {
            domains = await Domains();
        } catch (error) {
            addNotification({
                text: "Failed retrieving domain data",
                position: "bottom-right",
                type: "error",
                removeAfter: 5000,
            });
        }
    }

    onMount(async () => {
        connected = await Connected();
        if (connected) {
            reloadDomains();
        }
    });
</script>

<Window heading="Domains">
    {#if connected}
        <div>
            {#each domains as d}
                <Domain name={d.name} status={d.status} />
            {/each}
        </div>
        <div class="justify-start">
            <Button on:click{reloadDomains}>Refresh</Button>
        </div>
    {:else}
        <P class="text-center text-white" size="xl">You are not connected!</P>
    {/if}
</Window>
