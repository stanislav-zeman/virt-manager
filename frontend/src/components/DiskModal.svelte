<script lang="ts">
    import { Button, Modal } from "flowbite-svelte";
    import { AttachDevice, DetachDevice } from "../../wailsjs/go/app/App";
    import { createEventDispatcher } from "svelte";
    import { getNotificationsContext } from "svelte-notifications";

    const { addNotification } = getNotificationsContext();
    const dispatch = createEventDispatcher();

    export let domain: string;
    export let diskModal: boolean;
    async function attach() {
        try {
            await AttachDevice(domain, "", "", "", "");
            dispatch("domain-changed");
        } catch {
            addNotification({
                text: "Could not attach the device",
                position: "bottom-right",
                type: "error",
                removeAfter: 5000,
            });
        }
    }
    async function detach() {
        try {
            await DetachDevice(domain, "", "", "", "");
            dispatch("domain-changed");
        } catch {
            addNotification({
                text: "Could not detach the device",
                position: "bottom-right",
                type: "error",
                removeAfter: 5000,
            });
        }
    }
</script>

<Modal title="Disk Management" bind:open={diskModal} autoclose>
    <p class="text-base leading-relaxed text-gray-500 dark:text-gray-400">
        With less than a month to go before the European Union enacts new
        consumer privacy laws for its citizens, companies around the world are
        updating their terms of service agreements to comply.
    </p>
    <p class="text-base leading-relaxed text-gray-500 dark:text-gray-400">
        The European Union’s General Data Protection Regulation (G.D.P.R.) goes
        into effect on May 25 and is meant to ensure a common set of data rights
        in the European Union. It requires organizations to notify users as soon
        as possible of high-risk data breaches that could personally affect
        them.
    </p>

    <svelte:fragment slot="footer">
        <Button color="green" on:click={attach}>Attach</Button>
        <Button color="red" on:click={detach}>Detach</Button>
    </svelte:fragment>
</Modal>
