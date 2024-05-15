<script lang="ts">
    import { Button } from "flowbite-svelte";
    import {
        Start,
        Resume,
        Suspend,
        Shutdown,
        Destroy,
    } from "../../wailsjs/go/app/App";
    import { createEventDispatcher } from "svelte";
    import { getNotificationsContext } from "svelte-notifications";

    const { addNotification } = getNotificationsContext();
    const dispatch = createEventDispatcher();

    export let name: string;
    export let status: string;
    async function start() {
        try {
            await Start(name);
            dispatch("domain-changed");
        } catch {
            addNotification({
                text: "Could not start the domain",
                position: "bottom-right",
                type: "error",
                removeAfter: 5000,
            });
        }
    }
    async function resume() {
        try {
            await Resume(name);
            dispatch("domain-changed");
        } catch {
            addNotification({
                text: "Could not resume the domain",
                position: "bottom-right",
                type: "error",
                removeAfter: 5000,
            });
        }
    }
    async function suspend() {
        try {
            await Suspend(name);
            dispatch("domain-changed");
        } catch {
            addNotification({
                text: "Could not suspend the domain",
                position: "bottom-right",
                type: "error",
                removeAfter: 5000,
            });
        }
    }
    async function shutdown() {
        try {
            await Shutdown(name);
            dispatch("domain-changed");
        } catch {
            addNotification({
                text: "Could not shutdown the domain",
                position: "bottom-right",
                type: "error",
                removeAfter: 5000,
            });
        }
    }
    async function destroy() {
        try {
            await Destroy(name);
            dispatch("domain-changed");
        } catch {
            addNotification({
                text: "Could not destroy the domain",
                position: "bottom-right",
                type: "error",
                removeAfter: 5000,
            });
        }
    }
</script>

<div
    class="mx-10 my-5 flex flex-row border-2 rounded-md h-16 items-center content-center justify-evenly bg-slate-700"
>
    <div class="flex flex-row justify-evenly w-1/2">
        <div>
            Name: <b>{name}</b>
        </div>
        <div>
            Status: <b>{status}</b>
        </div>
    </div>
    <div>
        {#if status === "Shutoff"}
            <Button color="green" on:click={start}>Start</Button>
        {/if}
        {#if status === "Paused"}
            <Button color="green" on:click={resume}>Resume</Button>
        {/if}
        {#if status === "Running"}
            <Button color="yellow" on:click={suspend}>Suspend</Button>
            <Button on:click={shutdown}>Shutdown</Button>
            <Button color="red" on:click={destroy}>Destroy</Button>
        {/if}
    </div>
</div>
