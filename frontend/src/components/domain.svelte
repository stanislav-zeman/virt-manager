<script lang="ts">
    import { Label, Input, Button, Modal } from "flowbite-svelte";
    import {
        Start,
        Resume,
        Suspend,
        Shutdown,
        Destroy,
        AttachDevice,
        DetachDevice,
    } from "../../wailsjs/go/app/App";
    import { createEventDispatcher } from "svelte";
    import { getNotificationsContext } from "svelte-notifications";
    import { FloppyDiskSolid } from "flowbite-svelte-icons";

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

    let diskModal = false;
    let diskType: string;
    let diskDevice: string;
    let path: string;
    let target: string;
    async function attach() {
        try {
            await AttachDevice(name, diskType, diskDevice, path, target);
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
            await DetachDevice(name, diskType, diskDevice, path, target);
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
    <div class="flex items-center justify-evenly">
        {#if status === "Shutoff"}
            <Button class="mx-1" color="green" on:click={start}>Start</Button>
        {/if}
        {#if status === "Paused"}
            <Button class="mx-1" color="green" on:click={resume}>Resume</Button>
        {/if}
        {#if status === "Running"}
            <Button class="mx-1" color="yellow" on:click={suspend}
                >Suspend</Button
            >
            <Button class="mx-1" on:click={shutdown}>Shutdown</Button>
            <Button class="mx-1" color="red" on:click={destroy}>Destroy</Button>
        {/if}
        {#if status === "Running"}
            <Button
                class="!p-2 mx-1"
                color="light"
                on:click={() => (diskModal = true)}
                ><FloppyDiskSolid class="w-6 h-6" /></Button
            >
            <Modal
                title="Disk Management"
                size="lg"
                bind:open={diskModal}
                autoclose
            >
                <form class="flex-col my-6 w-full">
                    <div class="flex my-3 justify-evenly w-full">
                        <div class="w-1/3">
                            <Label for="disk-type" class="block text-base mb-2"
                                >Disk Type</Label
                            >
                            <Input
                                id="disk-type"
                                bind:value={diskType}
                                placeholder="file"
                            />
                        </div>
                        <div class="w-1/3">
                            <Label
                                for="disk-device"
                                class="block text-base mb-2">Disk Device</Label
                            >
                            <Input
                                id="disk-device"
                                bind:value={diskDevice}
                                placeholder="disk"
                            />
                        </div>
                    </div>
                    <div class="flex my-3 justify-evenly w-full">
                        <div class="w-1/3">
                            <Label for="path" class="block text-base mb-2"
                                >Path</Label
                            >
                            <Input
                                id="path"
                                bind:value={path}
                                placeholder="/var/lib/libvirt/random.img"
                            />
                        </div>
                        <div class="w-1/3">
                            <Label for="target" class="block text-base mb-2"
                                >Target</Label
                            >
                            <Input
                                id="target"
                                bind:value={target}
                                placeholder="vdb"
                            />
                        </div>
                    </div>
                </form>
                <svelte:fragment slot="footer">
                    <Button color="green" on:click={attach}>Attach</Button>
                    <Button color="red" on:click={detach}>Detach</Button>
                </svelte:fragment>
            </Modal>
        {/if}
    </div>
</div>
