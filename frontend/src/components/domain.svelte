<script lang="ts">
    import { Button } from "flowbite-svelte";
    export let name: string;
    export let status: string;
</script>

<div
    class="mx-10 my-5 flex flex-row border-2 rounded-md h-16 items-center content-center justify-evenly"
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
        <Button>Start</Button>
        <Button>Shutdown</Button>
        <Button>Suspend</Button>
        <Button>Resume</Button>
        <Button>Destroy</Button>
    </div>
</div>
