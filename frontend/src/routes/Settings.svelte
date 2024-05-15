<script lang="ts">
    import Window from "../layout/Window.svelte";
    import {
        Label,
        Input,
        InputAddon,
        ButtonGroup,
        Button,
    } from "flowbite-svelte";
    import { ServerOutline } from "flowbite-svelte-icons";
    import { Connect } from "../../wailsjs/go/app/App";
    import { getNotificationsContext } from "svelte-notifications";

    const { addNotification } = getNotificationsContext();

    let hypervisorUri: string;
    async function connect() {
        try {
            await Connect(hypervisorUri);
            addNotification({
                text: "Connected",
                position: "bottom-right",
                type: "success",
                removeAfter: 3000,
            });
        } catch (error) {
            addNotification({
                text: "Could not connect to the hypervisor",
                position: "bottom-right",
                type: "error",
                removeAfter: 5000,
            });
        }
    }
</script>

<Window heading="Settings">
    <div class="flex-col">
        <form class="flex-col">
            <Label for="hypervisor-uri" class="block text-lg text-white mb-2"
                >Hypervisor URI</Label
            >
            <div class="flex content-center justify-center">
                <ButtonGroup class="w-2/5 mx-2">
                    <InputAddon>
                        <ServerOutline
                            class="w-4 h-4 text-gray-500 dark:text-gray-400"
                        />
                    </InputAddon>
                    <Input
                        id="hypervisor-uri"
                        bind:value={hypervisorUri}
                        placeholder="qemu:///system"
                    />
                </ButtonGroup>
                <Button class="mx-2" on:click={connect}>Connect</Button>
            </div>
        </form>
    </div>
</Window>
