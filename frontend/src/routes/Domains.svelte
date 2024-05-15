<script lang="ts">
    import { onMount } from "svelte";
    import Window from "../layout/Window.svelte";
    import { Domains, Connected } from "../../wailsjs/go/app/App";
    import { domain } from "../../wailsjs/go/models";
    import { Button, P } from "flowbite-svelte";
    import { ArrowRightOutline } from "flowbite-svelte-icons";
    import Domain from "../components/Domain.svelte";
    import { getNotificationsContext } from "svelte-notifications";

    const { addNotification } = getNotificationsContext();

    let connected: boolean;
    let domains: domain.Domain[] = [];
    async function reloadDomains() {
        try {
            domains = await Domains();
            console.log(domains);
            addNotification({
                text: "Reloaded domains",
                position: "bottom-right",
                type: "info",
                removeAfter: 3000,
            });
        } catch (error) {
            addNotification({
                text: "Failed retrieving domain data",
                position: "bottom-right",
                type: "error",
                removeAfter: 5000,
            });
        }
    }

    async function create() {
        addNotification({
            text: "Not yet implemented",
            position: "bottom-right",
            type: "warn",
            removeAfter: 5000,
        });
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
                <Domain
                    name={d.name}
                    status={d.status}
                    on:domain-changed={reloadDomains}
                />
            {/each}
        </div>
        <div class="justify-start">
            <Button class="mx-1" on:click={create}>Create</Button>
            <Button class="mx-1" on:click={reloadDomains}>Refresh</Button>
        </div>
    {:else}
        <P class="mb-6 text-center text-white" size="xl"
            >You are not connected!</P
        >
        <Button href="#/settings">
            Settings
            <ArrowRightOutline class="w-6 h-6 ms-2" />
        </Button>
    {/if}
</Window>
