<script>
    import Window from "../layout/Window.svelte";
    import { Label, Input, InputAddon, ButtonGroup } from "flowbite-svelte";
    import { ServerOutline } from "flowbite-svelte-icons";
</script>

<Window heading="Settings">
    <div class="flex-col">
        <Label for="hypervisor-uri" class="block text-white mb-2"
            >Hypervisor URI</Label
        >
        <ButtonGroup class="w-2/5">
            <InputAddon>
                <ServerOutline
                    class="w-4 h-4 text-gray-500 dark:text-gray-400"
                />
            </InputAddon>
            <Input id="hypervisor-uri" placeholder="quemu:///system" />
        </ButtonGroup>
    </div>
</Window>
